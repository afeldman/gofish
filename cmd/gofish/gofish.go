package cmd

import (
	"github.com/afeldman/gofish/pkg/logger"
	"github.com/spf13/cobra"
)

var (
	// logLevel is a value to indicate how verbose the user would like the logs to be.
	logLevel string
	rootCmd  *cobra.Command
)

var globalUsage = `The package manager.
`

func NewRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "gofish",
		Short:        globalUsage,
		Long:         globalUsage,
		SilenceUsage: true,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if err := logger.Init(logLevel, "./gofish.log", false); err != nil {
				return
			}

		},
	}
	p := cmd.PersistentFlags()
	p.StringVar(&logLevel, "log-level", "info", "log level")

	cmd.AddCommand(
		newCleanupCmd(),
		newCreateCmd(),
		newHomeCmd(),
		newInfoCmd(),
		newInitCmd(),
		newInstallCmd(),
		newLinkCmd(),
		newLintCmd(),
		newListCmd(),
		newPinCmd(),
		newRigCmd(),
		newRottenCmd(),
		newSearchCmd(),
		newSwitchCmd(),
		newTankCmd(),
		newUninstallCmd(),
		newUnlinkCmd(),
		newUnpinCmd(),
		newUpdateCmd(),
		newUpgradeCmd(),
		newVersionCmd(),
	)

	return cmd
}
