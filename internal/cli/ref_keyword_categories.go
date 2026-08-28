package cli

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var refKeywordCategoriesCmd = &cobra.Command{
	Use:   "keyword-categories",
	Short: "Keyword categories",
	Long:  "Commands for keyword category reference data.",
}

var refKeywordCategoriesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List keyword categories",
	Long:  "List available keyword categories (metadata for attribute types).",
	RunE:  runRefKeywordCategoriesList,
}

func init() {
	refKeywordCategoriesCmd.AddCommand(refKeywordCategoriesListCmd)
}

func runRefKeywordCategoriesList(cmd *cobra.Command, args []string) error {
	client, err := loadClient()
	if err != nil {
		return err
	}

	items, err := client.GetKeywordCategories(cmd.Context())
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
