package cli

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var refContactLogActivityTypesCmd = &cobra.Command{
	Use:   "contact-log-activity-types",
	Short: "Contact log activity types",
	Long:  "Commands for contact log activity type reference data.",
}

var refContactLogActivityTypesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List contact log activity types",
	Long:  "List available contact log activity types.",
	RunE:  runRefContactLogActivityTypesList,
}

func init() {
	refContactLogActivityTypesCmd.AddCommand(refContactLogActivityTypesListCmd)
}

func runRefContactLogActivityTypesList(cmd *cobra.Command, args []string) error {
	client, err := loadClient()
	if err != nil {
		return err
	}

	items, err := client.GetContactLogActivityTypes(cmd.Context())
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
