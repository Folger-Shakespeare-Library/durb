package cli

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var refReportTypesCmd = &cobra.Command{
	Use:   "report-types",
	Short: "Report types",
	Long:  "Commands for report type reference data.",
}

var refReportTypesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List report types",
	Long:  "List available report types.",
	RunE:  runRefReportTypesList,
}

func init() {
	refReportTypesCmd.AddCommand(refReportTypesListCmd)
}

func runRefReportTypesList(cmd *cobra.Command, args []string) error {
	client, err := loadClient()
	if err != nil {
		return err
	}

	items, err := client.GetReportTypes(cmd.Context())
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
