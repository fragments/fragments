package main

import (
	"context"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/fragments/fragments/internal/server"
	"github.com/fragments/fragments/internal/state"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func newEnvironmentCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:     "environment",
		Aliases: []string{"env"},
		Short:   "Create or modify target deployment environments",
	}

	cmd.AddCommand(newEnvironmentCreateCommand())

	return cmd
}

func newEnvironmentCreateCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "create",
		Short: "Create a new environment",
	}

	flags := cmd.Flags()
	name := flags.StringP("name", "n", "", "Environment name")
	infraName := flags.StringP("infrastructure", "i", "", "Infrastructure provider")
	username := flags.StringP("username", "u", "", "Username for authenticating with infrastructure provider")
	password := flags.StringP("password", "p", "", "Password for authenticating with infrastructure provider")
	labels := flags.StringSliceP("label", "l", []string{}, "Label(s) to put on environment")

	cmd.PreRunE = func(cmd *cobra.Command, args []string) error {
		if *name == "" {
			return errors.New("name must be set")
		}
		if *infraName == "" {
			return errors.New("infrastructure must be set")
		}
		if *username == "" {
			return errors.New("username must be set")
		}
		if *password == "" {
			return errors.New("password must be set")
		}

		return nil
	}

	cmd.Run = func(cmd *cobra.Command, args []string) {
		infra, err := parseInfrastructure(*infraName)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			os.Exit(1)
		}

		l, err := extractLabels(*labels)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: format must be key=value\n", err)
			os.Exit(1)
		}

		kv, err := getKV(flags)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			os.Exit(1)
		}

		input := &server.EnvironmentInput{
			Name:           *name,
			Labels:         l,
			Infrastructure: infra,
			Username:       *username,
			Password:       *password,
		}

		s := server.New(kv, nil)
		ctx := context.TODO()
		if err := s.CreateEnvironment(ctx, input); err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			os.Exit(1)
		}
	}

	return cmd
}

func parseInfrastructure(name string) (state.InfrastructureType, error) {
	n := strings.ToLower(name)
	switch n {
	case string(state.InfrastructureTypeAWS):
		return state.InfrastructureTypeAWS, nil
	default:
		return "", errors.Errorf("unsupported infrastructure %q", name)
	}
}

var labelRegex = regexp.MustCompile(`^(\w+)\s*=\s*(\w+)$`)

func extractLabels(raw []string) (map[string]string, error) {
	out := make(map[string]string)
	for _, l := range raw {
		parts := labelRegex.FindStringSubmatch(l)
		if len(parts) < 3 {
			return nil, errors.Errorf("%q not a valid label", l)
		}
		k := parts[1]
		v := parts[2]
		out[k] = v
	}
	return out, nil
}
