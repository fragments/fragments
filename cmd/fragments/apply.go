package main

import (
	"context"
	"encoding/hex"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/fragments/fragments/internal/client"
	"github.com/fragments/fragments/internal/reconciler"
	"github.com/fragments/fragments/internal/server"
	"github.com/fragments/fragments/internal/state"
	"github.com/golang/sync/errgroup"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func newApplyCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "apply [dir]",
		Short: "Apply model changes",
	}

	flags := cmd.Flags()
	ignore := flags.StringSliceP("ignore", "i", []string{"node_modules", "vendor"}, "File/directory patterns to ignore")

	cmd.PreRunE = func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return errors.New("target directory must be set")
		}

		for _, target := range args {
			stat, err := os.Stat(target)
			if err != nil {
				return err
			}
			if !stat.IsDir() {
				return errors.Errorf("target must be a directory: %s", target)
			}
		}

		return nil
	}

	cmd.Run = func(cmd *cobra.Command, args []string) {
		models, err := getModels(args, *ignore)
		checkErr(err)

		excludeSource := functionDirs(models)
		excludeSource = append(excludeSource, *ignore...)

		fileStore, err := getFilestore()
		checkErr(errors.Wrap(err, "could not set up local filestore"))

		etcd, err := getETCD(flags)
		checkErr(errors.Wrap(err, "could not set up etcd"))

		vault, err := getVault(flags)
		checkErr(errors.Wrap(err, "could not set up vault"))

		ctx := contextFromSignal()

		s := server.New(etcd, nil, fileStore)
		err = apply(ctx, s, models, excludeSource)
		checkErr(err)

		reco := reconciler.New(etcd, vault, fileStore)
		err = reco.Run(ctx)
		checkErr(err)

		err = etcd.Close()
		checkErr(err)
	}

	return cmd
}

// functionDirs returns directories containing a function model.
func functionDirs(models []client.Model) []string {
	out := []string{}
	for _, r := range models {
		if r.Type() == client.ModelTypeFunction {
			out = append(out, filepath.Dir(r.File()))
		}
	}
	return out
}

// apply applies all models on the server.
func apply(ctx context.Context, server *server.Server, models []client.Model, excludeSource []string) error {
	g, ctx := errgroup.WithContext(ctx)
	for _, r := range models {
		r := r
		g.Go(func() error {
			meta := r.Meta()
			file := r.File()
			if function, ok := r.(client.Function); ok {
				spec := function.Function()
				if err := applyFunction(ctx, server, meta, file, spec, excludeSource); err != nil {
					return errors.Wrap(err, "could not apply function")
				}
				return nil
			}
			if deployment, ok := r.(client.Deployment); ok {
				if err := applyDeployment(ctx, server, meta, deployment.Deployment()); err != nil {
					return errors.Wrap(err, "could not apply deployment")
				}
				return nil
			}
			return errors.Errorf("unsupported model %q: %s", r.Type(), file)
		})
	}

	return g.Wait()
}

// getModels walks the target directories and returns discovered models.
func getModels(targets, ignore []string) ([]client.Model, error) {
	// Loop through all targets to resolve models
	modelPaths := []string{}
	for _, target := range targets {
		paths, err := client.Walk(target, ignore)
		checkErr(errors.Wrap(err, target))
		modelPaths = append(modelPaths, paths...)
	}

	models := []client.Model{}
	for _, path := range modelPaths {
		res, err := client.Load(path)
		checkErr(errors.Wrapf(err, "could not load model: %s", path))
		// A nil model is returned in case a model was not found in file
		if res != nil {
			models = append(models, res...)
		}
	}

	if err := client.CheckDuplicates(models); err != nil {
		return nil, err
	}

	return models, nil
}

func applyFunction(ctx context.Context, srv *server.Server, meta *client.Meta, file string, spec *client.FunctionSpec, ignore []string) error {
	// Collect function source files
	dir := filepath.Dir(file)
	source, err := client.CollectSource(dir, ignore)
	if err != nil {
		return errors.Wrap(err, "could not collect function source")
	}

	if len(source) == 0 {
		return errors.New("function contains no source")
	}

	// Ensure consistent order for hashing
	sort.Strings(source)

	// Calculate checksum
	// Exclude:
	// - Function definition (so source hash doesn't change on config change)
	// - Sub functions
	checksumExclude := []string{file}
	for _, pattern := range ignore {
		// Ensure we don't exclude a parent function's directory. If do the pattern
		// will match the sub function and its source won't get included in the
		// checksum (don't exclude /foo if we're calculating checksum for /foo/bar)
		if !strings.Contains(dir, pattern) {
			checksumExclude = append(checksumExclude, pattern)
		}
	}
	shasum, err := client.Checksum(source, checksumExclude)
	if err != nil {
		return errors.Wrap(err, "could not calculate source checksum")
	}

	// Construct request
	function := &state.Function{
		Checksum: hex.EncodeToString(shasum),
		Meta: state.Meta{
			Name:   meta.Name,
			Labels: meta.Labels,
		},
		Runtime: spec.Runtime,
	}
	if spec.AWS != nil {
		function.AWS = &state.FunctionAWS{
			Timeout: spec.AWS.Timeout,
			Memory:  spec.AWS.Memory,
		}
	}

	uploadReq, err := srv.PutFunction(ctx, function)
	if err != nil {
		return errors.Wrap(err, "could not put function")
	}

	if uploadReq != nil {
		if err := upload(source, uploadReq); err != nil {
			return errors.Wrap(err, "upload failed")
		}

		if err := srv.ConfirmUpload(ctx, uploadReq.Token); err != nil {
			return errors.Wrap(err, "could not confirm upload")
		}
	}

	return nil
}

func applyDeployment(ctx context.Context, srv *server.Server, meta *client.Meta, deployment *client.DeploymentSpec) error {
	deploy := &state.Deployment{
		Meta: state.Meta{
			Name:   meta.Name,
			Labels: meta.Labels,
		},
		EnvironmentLabels: deployment.EnvironmentLabels,
		FunctionLabels:    deployment.FunctionLabels,
	}
	if err := srv.PutDeployment(ctx, deploy); err != nil {
		return errors.Wrap(err, "PutDeployment failed")
	}
	return nil
}

func upload(source []string, uploadReq *server.UploadRequest) error {
	targz, err := client.Compress(source)
	if err != nil {
		return errors.Wrap(err, "could not archive source")
	}

	if err := client.Upload(targz, uploadReq.URL); err != nil {
		return errors.Wrap(err, "upload failed")
	}

	return nil
}
