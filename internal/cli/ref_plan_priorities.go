package cli

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var refPlanPrioritiesCmd = &cobra.Command{
	Use:   "plan-priorities",
	Short: "Plan priorities",
	Long:  "Commands for plan priority reference data.",
}

var refPlanPrioritiesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List plan prioritys",
	Long:  "List available plan prioritys.",
	RunE:  runRefPlanPrioritiesList,
}

func init() {
	refPlanPrioritiesCmd.AddCommand(refPlanPrioritiesListCmd)
}

func runRefPlanPrioritiesList(cmd *cobra.Command, args []string) error {
	client, err := loadClient()
	if err != nil {
		return err
	}

	items, err := client.GetPlanPriorities(cmd.Context())
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
