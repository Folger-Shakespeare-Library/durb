package cli

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var refPlanTypesCmd = &cobra.Command{
	Use:   "plan-types",
	Short: "Plan types",
	Long:  "Commands for plan type reference data.",
}

var refPlanTypesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List plan types",
	Long:  "List available plan types.",
	RunE:  runRefPlanTypesList,
}

func init() {
	refPlanTypesCmd.AddCommand(refPlanTypesListCmd)
}

func runRefPlanTypesList(cmd *cobra.Command, args []string) error {
	client, err := loadClient()
	if err != nil {
		return err
	}

	items, err := client.GetPlanTypes(cmd.Context())
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
