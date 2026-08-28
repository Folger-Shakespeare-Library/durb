package cli

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var refPerformanceSegmentTypesCmd = &cobra.Command{
	Use:   "performance-segment-types",
	Short: "Performance segment types",
	Long:  "Commands for performance segment type reference data.",
}

var refPerformanceSegmentTypesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List performance segment types",
	Long:  "List available performance segment types.",
	RunE:  runRefPerformanceSegmentTypesList,
}

func init() {
	refPerformanceSegmentTypesCmd.AddCommand(refPerformanceSegmentTypesListCmd)
}

func runRefPerformanceSegmentTypesList(cmd *cobra.Command, args []string) error {
	client, err := loadClient()
	if err != nil {
		return err
	}

	items, err := client.GetPerformanceSegmentTypes(cmd.Context())
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
