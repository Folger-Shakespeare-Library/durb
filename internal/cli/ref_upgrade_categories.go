package cli

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var refUpgradeCategoriesCmd = &cobra.Command{
	Use:   "upgrade-categories",
	Short: "Upgrade categories",
	Long:  "Commands for upgrade category reference data.",
}

var refUpgradeCategoriesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List upgrade categories",
	Long:  "List available upgrade categories.",
	RunE:  runRefUpgradeCategoriesList,
}

func init() {
	refUpgradeCategoriesCmd.AddCommand(refUpgradeCategoriesListCmd)
}

func runRefUpgradeCategoriesList(cmd *cobra.Command, args []string) error {
	client, err := loadClient()
	if err != nil {
		return err
	}

	items, err := client.GetUpgradeCategories(cmd.Context())
	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(items, "", "  ")
	if err != nil {
		return fmt.Errorf("unable to format output: %w", err)
	}

	fmt.Fprintln(cmd.OutOrStdout(), string(data))
	return nil
}
