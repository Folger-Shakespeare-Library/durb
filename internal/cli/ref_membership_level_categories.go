package cli

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var refMembershipLevelCategoriesCmd = &cobra.Command{
	Use:   "membership-level-categories",
	Short: "Membership level categories",
	Long:  "Commands for membership level category reference data.",
}

var refMembershipLevelCategoriesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List membership level categories",
	Long:  "List available membership level categories.",
	RunE:  runRefMembershipLevelCategoriesList,
}

func init() {
	refMembershipLevelCategoriesCmd.AddCommand(refMembershipLevelCategoriesListCmd)
}

func runRefMembershipLevelCategoriesList(cmd *cobra.Command, args []string) error {
	client, err := loadClient()
	if err != nil {
		return err
	}

	items, err := client.GetMembershipLevelCategories(cmd.Context())
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
