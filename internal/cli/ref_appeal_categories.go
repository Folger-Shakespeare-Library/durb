package cli

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var refAppealCategoriesCmd = &cobra.Command{
	Use:   "appeal-categories",
	Short: "Appeal categories",
	Long:  "Commands for appeal category reference data.",
}

var refAppealCategoriesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List appeal categorys",
	Long:  "List available appeal categorys.",
	RunE:  runRefAppealCategoriesList,
}

func init() {
	refAppealCategoriesCmd.AddCommand(refAppealCategoriesListCmd)
}

func runRefAppealCategoriesList(cmd *cobra.Command, args []string) error {
	client, err := loadClient()
	if err != nil {
		return err
	}

	items, err := client.GetAppealCategories(cmd.Context())
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
