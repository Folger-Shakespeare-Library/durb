package cli

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var refSeasonsCmd = &cobra.Command{
	Use:   "seasons",
	Short: "Seasons",
	Long:  "Commands for season reference data.",
}

var refSeasonsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List seasons",
	Long:  "List available seasons.",
	RunE:  runRefSeasonsList,
}

func init() {
	refSeasonsCmd.AddCommand(refSeasonsListCmd)
}

func runRefSeasonsList(cmd *cobra.Command, args []string) error {
	client, err := loadClient()
	if err != nil {
		return err
	}

	items, err := client.GetSeasons(cmd.Context())
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
