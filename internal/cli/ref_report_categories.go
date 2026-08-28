package cli

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var refReportCategoriesCmd = &cobra.Command{
	Use:   "report-categories",
	Short: "Report categories",
	Long:  "Commands for report category reference data.",
}

var refReportCategoriesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List report categories",
	Long:  "List available report categories.",
	RunE:  runRefReportCategoriesList,
}

func init() {
	refReportCategoriesCmd.AddCommand(refReportCategoriesListCmd)
}

func runRefReportCategoriesList(cmd *cobra.Command, args []string) error {
	client, err := loadClient()
	if err != nil {
		return err
	}

	items, err := client.GetReportCategories(cmd.Context())
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
