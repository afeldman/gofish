package main

import (
	"os"

	cmd "github.com/afeldman/gofish/cmd/gofish"
)

func main() {
	rootCmd = cmd.NewRootCmd()
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
