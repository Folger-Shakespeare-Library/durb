package cli

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var refPlannedGivingPurposesCmd = &cobra.Command{
	Use:   "planned-giving-purposes",
	Short: "Planned giving purposes",
	Long:  "Commands for planned giving purpose reference data.",
}

var refPlannedGivingPurposesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List planned giving purposes",
	Long:  "List available planned giving purposes.",
	RunE:  runRefPlannedGivingPurposesList,
}

func init() {
	refPlannedGivingPurposesCmd.AddCommand(refPlannedGivingPurposesListCmd)
}

func runRefPlannedGivingPurposesList(cmd *cobra.Command, args []string) error {
	client, err := loadClient()
	if err != nil {
		return err
	}

	items, err := client.GetPlannedGivingPurposes(cmd.Context())
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
