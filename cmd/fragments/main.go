package main

import (
	"time"

	"github.com/fragments/fragments/internal/backend"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func main() {
	var cmd = &cobra.Command{
		Use: "fragments",
	}

	flags := cmd.PersistentFlags()
	flags.StringSliceP("etcd", "e", []string{"0.0.0.0:2379"}, "ETCD endpoints to connect to for storing state")

	cmd.AddCommand(newApplyCommand())
	cmd.AddCommand(newEnvironmentCommand())

	_ = cmd.Execute()
}

func getKV(flags *pflag.FlagSet) (backend.KV, error) {
	endpoints, err := flags.GetStringSlice("etcd")
	if err != nil {
		return nil, err
	}
	etcd, err := backend.NewETCDClient(endpoints, 3*time.Second)
	if err != nil {
		return nil, errors.Wrap(err, "could not get backend")
	}
	return etcd, nil
}
