package cli

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var refPlannedGivingFundingsCmd = &cobra.Command{
	Use:   "planned-giving-fundings",
	Short: "Planned giving fundings",
	Long:  "Commands for planned giving funding reference data.",
}

var refPlannedGivingFundingsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List planned giving fundings",
	Long:  "List available planned giving fundings.",
	RunE:  runRefPlannedGivingFundingsList,
}

func init() {
	refPlannedGivingFundingsCmd.AddCommand(refPlannedGivingFundingsListCmd)
}

func runRefPlannedGivingFundingsList(cmd *cobra.Command, args []string) error {
	client, err := loadClient()
	if err != nil {
		return err
	}

	items, err := client.GetPlannedGivingFundings(cmd.Context())
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
