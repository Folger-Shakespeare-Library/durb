package cli

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var refPlannedGivingGiftTypesCmd = &cobra.Command{
	Use:   "planned-giving-gift-types",
	Short: "Planned giving gift types",
	Long:  "Commands for planned giving gift type reference data.",
}

var refPlannedGivingGiftTypesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List planned giving gift types",
	Long:  "List available planned giving gift types.",
	RunE:  runRefPlannedGivingGiftTypesList,
}

func init() {
	refPlannedGivingGiftTypesCmd.AddCommand(refPlannedGivingGiftTypesListCmd)
}

func runRefPlannedGivingGiftTypesList(cmd *cobra.Command, args []string) error {
	client, err := loadClient()
	if err != nil {
		return err
	}

	items, err := client.GetPlannedGivingGiftTypes(cmd.Context())
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
