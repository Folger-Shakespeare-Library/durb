package cli

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var refContactTypesCmd = &cobra.Command{
	Use:   "contact-types",
	Short: "Contact types",
	Long:  "Commands for contact type reference data.",
}

var refContactTypesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List contact types",
	Long:  "List available contact types.",
	RunE:  runRefContactTypesList,
}

func init() {
	refContactTypesCmd.AddCommand(refContactTypesListCmd)
}

func runRefContactTypesList(cmd *cobra.Command, args []string) error {
	client, err := loadClient()
	if err != nil {
		return err
	}

	items, err := client.GetContactTypes(cmd.Context())
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
