package cli

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var refSeasonTypesCmd = &cobra.Command{
	Use:   "season-types",
	Short: "Season types",
	Long:  "Commands for season type reference data.",
}

var refSeasonTypesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List season types",
	Long:  "List available season types.",
	RunE:  runRefSeasonTypesList,
}

func init() {
	refSeasonTypesCmd.AddCommand(refSeasonTypesListCmd)
}

func runRefSeasonTypesList(cmd *cobra.Command, args []string) error {
	client, err := loadClient()
	if err != nil {
		return err
	}

	items, err := client.GetSeasonTypes(cmd.Context())
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
