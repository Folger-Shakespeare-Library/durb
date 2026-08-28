package cli

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var refPlannedGivingSourcesCmd = &cobra.Command{
	Use:   "planned-giving-sources",
	Short: "Planned giving sources",
	Long:  "Commands for planned giving source reference data.",
}

var refPlannedGivingSourcesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List planned giving sources",
	Long:  "List available planned giving sources.",
	RunE:  runRefPlannedGivingSourcesList,
}

func init() {
	refPlannedGivingSourcesCmd.AddCommand(refPlannedGivingSourcesListCmd)
}

func runRefPlannedGivingSourcesList(cmd *cobra.Command, args []string) error {
	client, err := loadClient()
	if err != nil {
		return err
	}

	items, err := client.GetPlannedGivingSources(cmd.Context())
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
