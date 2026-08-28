package cli

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var refInterestCategoriesCmd = &cobra.Command{
	Use:   "interest-categories",
	Short: "Interest categories",
	Long:  "Commands for interest category reference data.",
}

var refInterestCategoriesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List interest categories",
	Long:  "List available interest categories (metadata for interest types).",
	RunE:  runRefInterestCategoriesList,
}

func init() {
	refInterestCategoriesCmd.AddCommand(refInterestCategoriesListCmd)
}

func runRefInterestCategoriesList(cmd *cobra.Command, args []string) error {
	client, err := loadClient()
	if err != nil {
		return err
	}

	items, err := client.GetInterestCategories(cmd.Context())
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
