package cli

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var refPlanStatusesCmd = &cobra.Command{
	Use:   "plan-statuses",
	Short: "Plan statuses",
	Long:  "Commands for plan status reference data.",
}

var refPlanStatusesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List plan statuss",
	Long:  "List available plan statuss.",
	RunE:  runRefPlanStatusesList,
}

func init() {
	refPlanStatusesCmd.AddCommand(refPlanStatusesListCmd)
}

func runRefPlanStatusesList(cmd *cobra.Command, args []string) error {
	client, err := loadClient()
	if err != nil {
		return err
	}

	items, err := client.GetPlanStatuses(cmd.Context())
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
