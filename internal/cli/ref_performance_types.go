package cli

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var refPerformanceTypesCmd = &cobra.Command{
	Use:   "performance-types",
	Short: "Performance types",
	Long:  "Commands for performance type reference data.",
}

var refPerformanceTypesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List performance types",
	Long:  "List available performance types.",
	RunE:  runRefPerformanceTypesList,
}

func init() {
	refPerformanceTypesCmd.AddCommand(refPerformanceTypesListCmd)
}

func runRefPerformanceTypesList(cmd *cobra.Command, args []string) error {
	client, err := loadClient()
	if err != nil {
		return err
	}

	items, err := client.GetPerformanceTypes(cmd.Context())
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
