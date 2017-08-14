package main

import (
	"github.com/spf13/cobra"
)

func main() {
	var cmd = &cobra.Command{
		Use: "fragments",
	}

	cmd.AddCommand(newApplyCommand())

	_ = cmd.Execute()
}
