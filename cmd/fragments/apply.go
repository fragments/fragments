package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"os/signal"
	"path/filepath"
	"sort"
	"strings"
	"syscall"

	"github.com/fragments/fragments/internal/client"
	"github.com/fragments/fragments/internal/filestore"
	"github.com/fragments/fragments/internal/server"
	"github.com/fragments/fragments/internal/state"
	"github.com/golang/sync/errgroup"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func newApplyCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "apply [dir]",
		Short: "Apply resource changes",
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
		walkOptions := &client.WalkOptions{
			Ignore: *ignore,
		}

		// Loop through all targets to resolve resources
		resourcePaths := []string{}
		for _, target := range args {
			paths, err := client.Walk(target, walkOptions)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%s: %s\n", target, err)
				os.Exit(1)
			}
			resourcePaths = append(resourcePaths, paths...)
		}

		resources := []client.Resource{}
		for _, path := range resourcePaths {
			res, err := client.Load(path)
			if err != nil {
				fmt.Fprintf(os.Stderr, "could not load resource: %s: %s\n", err, path)
				os.Exit(1)
			}
			// A nil resource is returned in case the resource was not found in file
			if res != nil {
				resources = append(resources, res...)
			}
		}

		if err := client.CheckDuplicates(resources); err != nil {
			if err != nil {
				fmt.Fprintf(os.Stderr, "%s\n", err)
				os.Exit(1)
			}
		}

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		// Cancel context on signal
		go func() {
			sig := make(chan os.Signal, 1)
			signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
			<-sig
			cancel()
		}()

		excludeSource := *ignore
		for _, r := range resources {
			if r.Type() == client.ResourceTypeFunction {
				excludeSource = append(excludeSource, filepath.Dir(r.File()))
			}
		}

		home, err := homedir.Dir()
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			os.Exit(1)
		}

		uploads := filepath.Join(home, ".fragments", "uploads")
		source := filepath.Join(home, ".fragments", "source")
		sourceStore, err := filestore.NewLocal(uploads, source)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			os.Exit(1)
		}

		kv, err := getKV(flags)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			os.Exit(1)
		}

		s := server.New(kv, nil, sourceStore)
		g, ctx := errgroup.WithContext(ctx)
		for _, r := range resources {
			r := r
			g.Go(func() error {
				meta := r.Meta()
				file := r.File()
				if function, ok := r.(client.Function); ok {
					spec := function.Function()
					if err := applyFunction(ctx, s, meta, file, spec, excludeSource); err != nil {
						return errors.Wrap(err, "could not apply function")
					}
					return nil
				}
				if deployment, ok := r.(client.Deployment); ok {
					if err := applyDeployment(ctx, s, meta, deployment.Deployment()); err != nil {
						return errors.Wrap(err, "could not apply deployment")
					}
					return nil
				}
				return errors.Errorf("unsupported resource %q: %s", r.Type(), file)
			})
		}

		if err := g.Wait(); err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			os.Exit(1)
		}

		if kvCloser, ok := kv.(io.Closer); ok {
			if err := kvCloser.Close(); err != nil {
				fmt.Fprintf(os.Stderr, "could not close etcd: %s\n", err)
			}
		}
	}

	return cmd
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
