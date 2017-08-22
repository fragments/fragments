package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"sort"
	"strings"
	"syscall"
	"time"

	"github.com/fragments/fragments/internal/backend"
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
	etcdEndpoints := flags.StringSliceP("etcd", "e", []string{"0.0.0.0:2379"}, "ETCD endpoints to connect to for storing state")

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

		etcd, err := backend.NewETCDClient(*etcdEndpoints, 3*time.Second)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			os.Exit(1)
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

		s := server.New(etcd, sourceStore)
		g, ctx := errgroup.WithContext(ctx)
		for _, r := range resources {
			r := r
			g.Go(func() error {
				if function, ok := r.(client.Function); ok {
					meta := r.Meta()
					file := r.File()
					spec := function.Function()
					return applyFunction(ctx, s, meta, file, spec, excludeSource)
				}
				return nil
			})
		}

		if err := g.Wait(); err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			os.Exit(1)
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
	}

	return nil
}

func upload(source []string, uploadReq *server.UploadRequest) error {
	return errors.New("not implemented")
}
