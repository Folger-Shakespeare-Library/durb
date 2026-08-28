package cli

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var refPlannedGivingStatusesCmd = &cobra.Command{
	Use:   "planned-giving-statuses",
	Short: "Planned giving statuses",
	Long:  "Commands for planned giving status reference data.",
}

var refPlannedGivingStatusesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List planned giving statuss",
	Long:  "List available planned giving statuss.",
	RunE:  runRefPlannedGivingStatusesList,
}

func init() {
	refPlannedGivingStatusesCmd.AddCommand(refPlannedGivingStatusesListCmd)
}

func runRefPlannedGivingStatusesList(cmd *cobra.Command, args []string) error {
	client, err := loadClient()
	if err != nil {
		return err
	}

	items, err := client.GetPlannedGivingStatuses(cmd.Context())
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
