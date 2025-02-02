package cmd

import (
	"github.com/afeldman/gofish/pkg/ahoi"
)

// ensureFood checks to see if the default fish food exists.
//
// If the pack does not exist, this function will create it.
// If it does, it will update to the latest.
func ensureFood() error {
	ahoi.Ahoiln("Installing default fish food...")

	addArgs := []string{
		"https://github.com/afeldman/fish-food",
	}

	rigCmd, _, err := rootCmd.Find([]string{"rig", "add"})
	if err != nil {
		return err
	}
	return rigCmd.RunE(rigCmd, addArgs)
}
