//go:build !windows
// +build !windows

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newTankCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "tank",
		Short: "display information about fish's environment",
		RunE: func(cmd *cobra.Command, args []string) error {
			t := tank{}
			t.fill()
			for k, v := range t {
				fmt.Printf("export %s=%q\n", k, v)
			}
			fmt.Print("# Run this command to configure your shell:\n# eval $(gofish tank)\n")
			return nil
		},
	}
	return cmd
}
