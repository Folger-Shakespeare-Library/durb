package cli

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var refRelationshipCategoriesCmd = &cobra.Command{
	Use:   "relationship-categories",
	Short: "Relationship categories",
	Long:  "Commands for relationship category reference data.",
}

var refRelationshipCategoriesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List relationship categories",
	Long:  "List available relationship categories for constituents.",
	RunE:  runRefRelationshipCategoriesList,
}

func init() {
	refRelationshipCategoriesCmd.AddCommand(refRelationshipCategoriesListCmd)
}

func runRefRelationshipCategoriesList(cmd *cobra.Command, args []string) error {
	client, err := loadClient()
	if err != nil {
		return err
	}

	items, err := client.GetRelationshipCategories(cmd.Context())
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
