package cli

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var refContactPermissionTypesCmd = &cobra.Command{
	Use:   "contact-permission-types",
	Short: "Contact permission types",
	Long:  "Commands for contact permission type reference data.",
}

var refContactPermissionTypesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List contact permission types",
	Long:  "List available contact permission types (marketing consent).",
	RunE:  runRefContactPermissionTypesList,
}

func init() {
	refContactPermissionTypesCmd.AddCommand(refContactPermissionTypesListCmd)
}

func runRefContactPermissionTypesList(cmd *cobra.Command, args []string) error {
	client, err := loadClient()
	if err != nil {
		return err
	}

	items, err := client.GetContactPermissionTypes(cmd.Context())
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
