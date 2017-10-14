package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"github.com/fragments/fragments/internal/backend"
	"github.com/fragments/fragments/internal/filestore"
	homedir "github.com/mitchellh/go-homedir"
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
	flags.String("vault", "http://0.0.0.0:8200", "Vault address for storing secrets")

	cmd.AddCommand(newApplyCommand())
	cmd.AddCommand(newEnvironmentCommand())

	_ = cmd.Execute()
}

func getETCD(flags *pflag.FlagSet) (*backend.ETCD, error) {
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

func getVault(flags *pflag.FlagSet) (*backend.Vault, error) {
	address, err := flags.GetString("vault")
	if err != nil {
		return nil, err
	}
	vault, err := backend.NewVaultClient(address)
	if err != nil {
		return nil, errors.Wrap(err, "could not get secret backend")
	}
	return vault, nil
}

func getFilestore() (*filestore.Local, error) {
	home, err := homedir.Dir()
	if err != nil {
		return nil, err
	}
	uploads := filepath.Join(home, ".fragments", "uploads")
	source := filepath.Join(home, ".fragments", "source")
	sourceStore, err := filestore.NewLocal(uploads, source)
	if err != nil {
		return nil, err
	}
	return sourceStore, nil
}

func checkErr(err error) {
	if err == nil {
		return
	}

	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}

func contextFromSignal() context.Context {
	ctx, cancel := context.WithCancel(context.Background())

	// Cancel context on signal
	go func() {
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
		<-sig
		cancel()
	}()

	return ctx
}
