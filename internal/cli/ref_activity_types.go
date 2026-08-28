package cli

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var refActivityTypesCmd = &cobra.Command{
	Use:   "activity-types",
	Short: "Activity types",
	Long:  "Commands for special activity type reference data.",
}

var refActivityTypesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List activity types",
	Long:  "List available special activity types.",
	RunE:  runRefActivityTypesList,
}

func init() {
	refActivityTypesCmd.AddCommand(refActivityTypesListCmd)
}

func runRefActivityTypesList(cmd *cobra.Command, args []string) error {
	client, err := loadClient()
	if err != nil {
		return err
	}

	items, err := client.GetSpecialActivityTypes(cmd.Context())
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
