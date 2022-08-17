package main

import (
	"os"

	cmd "github.com/afeldman/gofish/cmd/gofish"
	"github.com/spf13/cobra"
)

var rootCmd *cobra.Command

func main() {
	rootCmd = cmd.NewRootCmd()
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
