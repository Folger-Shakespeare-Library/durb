package cli

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var refEventLevelsCmd = &cobra.Command{
	Use:   "event-levels",
	Short: "Event levels",
	Long:  "Commands for event level reference data.",
}

var refEventLevelsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List event levels",
	Long:  "List available event levels.",
	RunE:  runRefEventLevelsList,
}

func init() {
	refEventLevelsCmd.AddCommand(refEventLevelsListCmd)
}

func runRefEventLevelsList(cmd *cobra.Command, args []string) error {
	client, err := loadClient()
	if err != nil {
		return err
	}

	items, err := client.GetEventLevels(cmd.Context())
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
