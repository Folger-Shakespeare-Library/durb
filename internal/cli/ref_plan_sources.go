package cli

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var refPlanSourcesCmd = &cobra.Command{
	Use:   "plan-sources",
	Short: "Plan sources",
	Long:  "Commands for plan source reference data.",
}

var refPlanSourcesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List plan sources",
	Long:  "List available plan sources.",
	RunE:  runRefPlanSourcesList,
}

func init() {
	refPlanSourcesCmd.AddCommand(refPlanSourcesListCmd)
}

func runRefPlanSourcesList(cmd *cobra.Command, args []string) error {
	client, err := loadClient()
	if err != nil {
		return err
	}

	items, err := client.GetPlanSources(cmd.Context())
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
