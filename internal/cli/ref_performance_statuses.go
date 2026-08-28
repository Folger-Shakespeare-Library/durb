package cli

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var refPerformanceStatusesCmd = &cobra.Command{
	Use:   "performance-statuses",
	Short: "Performance statuses",
	Long:  "Commands for performance status reference data.",
}

var refPerformanceStatusesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List performance statuses",
	Long:  "List available performance statuses.",
	RunE:  runRefPerformanceStatusesList,
}

func init() {
	refPerformanceStatusesCmd.AddCommand(refPerformanceStatusesListCmd)
}

func runRefPerformanceStatusesList(cmd *cobra.Command, args []string) error {
	client, err := loadClient()
	if err != nil {
		return err
	}

	items, err := client.GetPerformanceStatuses(cmd.Context())
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
