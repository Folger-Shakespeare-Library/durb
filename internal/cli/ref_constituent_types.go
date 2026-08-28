package cli

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var refConstituentTypesCmd = &cobra.Command{
	Use:   "constituent-types",
	Short: "Constituent types",
	Long:  "Commands for constituent type reference data.",
}

var refConstituentTypesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List constituent types",
	Long:  "List available constituent types.",
	RunE:  runRefConstituentTypesList,
}

func init() {
	refConstituentTypesCmd.AddCommand(refConstituentTypesListCmd)
}

func runRefConstituentTypesList(cmd *cobra.Command, args []string) error {
	client, err := loadClient()
	if err != nil {
		return err
	}

	items, err := client.GetConstituentTypes(cmd.Context())
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
