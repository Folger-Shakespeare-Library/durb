package cli

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var refPlannedGivingCodesCmd = &cobra.Command{
	Use:   "planned-giving-codes",
	Short: "Planned giving codes",
	Long:  "Commands for planned giving code reference data.",
}

var refPlannedGivingCodesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List planned giving codes",
	Long:  "List available planned giving codes.",
	RunE:  runRefPlannedGivingCodesList,
}

func init() {
	refPlannedGivingCodesCmd.AddCommand(refPlannedGivingCodesListCmd)
}

func runRefPlannedGivingCodesList(cmd *cobra.Command, args []string) error {
	client, err := loadClient()
	if err != nil {
		return err
	}

	items, err := client.GetPlannedGivingCodes(cmd.Context())
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
