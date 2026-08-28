package cli

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var refAssociationTypesCmd = &cobra.Command{
	Use:   "association-types",
	Short: "Association types",
	Long:  "Commands for association type reference data.",
}

var refAssociationTypesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List association types",
	Long:  "List available association types.",
	RunE:  runRefAssociationTypesList,
}

func init() {
	refAssociationTypesCmd.AddCommand(refAssociationTypesListCmd)
}

func runRefAssociationTypesList(cmd *cobra.Command, args []string) error {
	client, err := loadClient()
	if err != nil {
		return err
	}

	items, err := client.GetAssociationTypes(cmd.Context())
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
