package cli

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var refOriginalSourcesCmd = &cobra.Command{
	Use:   "original-sources",
	Short: "Original sources",
	Long:  "Commands for original source reference data.",
}

var refOriginalSourcesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List original sources",
	Long:  "List available original sources.",
	RunE:  runRefOriginalSourcesList,
}

func init() {
	refOriginalSourcesCmd.AddCommand(refOriginalSourcesListCmd)
}

func runRefOriginalSourcesList(cmd *cobra.Command, args []string) error {
	client, err := loadClient()
	if err != nil {
		return err
	}

	items, err := client.GetOriginalSources(cmd.Context())
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
