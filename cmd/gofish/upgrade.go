package cmd

import (
	"fmt"
	"sort"
	"time"

	"github.com/Masterminds/semver"
	"github.com/afeldman/gofish/pkg/ahoi"
	"github.com/spf13/cobra"
)

func newUpgradeCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "upgrade [food..]",
		Short: "upgrade all fish food. If arguments are provided, only the specified fish foods are upgraded.",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := updateRigs(); err != nil {
				return err
			}
			var foodNames []string
			if len(args) > 0 {
				foodNames = args
			} else {
				foodNames = findFood()
			}
			nothingUpgraded := true
			for _, name := range foodNames {
				installedVersions := findFoodVersions(name)
				if len(installedVersions) == 0 {
					ahoi.Ahoif("%s: no installed versions to upgrade\n", name)
					continue
				}
				vs := make(semver.Collection, len(installedVersions))
				for i, r := range installedVersions {
					v, err := semver.NewVersion(r)
					if err != nil {
						ahoi.Ahoif("Upgrading %s...\n", name)
						return fmt.Errorf("Error parsing version: %v", err)
					}
					vs[i] = v
				}
				sort.Sort(vs)
				// we can safely assume there's at least one release installed
				latestInstalledVersion := vs[len(vs)-1]
				food, err := getFood(name)
				if err != nil {
					return err
				}
				if latestInstalledVersion.String() == food.Version {
					continue
				}
				nothingUpgraded = false
				ahoi.Ahoif("Upgrading %s...\n", food.Name)
				start := time.Now()
				if err := food.Install(); err != nil {
					return err
				}
				t := time.Now()
				ahoi.Successf("%s %s: upgraded in %s\n", food.Name, food.Version, t.Sub(start).String())
			}
			if nothingUpgraded {
				ahoi.Successf("Everything up to date!\n")
			}
			return nil
		},
	}
	return cmd
}
