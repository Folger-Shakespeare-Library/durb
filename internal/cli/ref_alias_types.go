package cli

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var refAliasTypesCmd = &cobra.Command{
	Use:   "alias-types",
	Short: "Alias types",
	Long:  "Commands for alias type reference data.",
}

var refAliasTypesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List alias types",
	Long:  "List available alias types.",
	RunE:  runRefAliasTypesList,
}

func init() {
	refAliasTypesCmd.AddCommand(refAliasTypesListCmd)
}

func runRefAliasTypesList(cmd *cobra.Command, args []string) error {
	client, err := loadClient()
	if err != nil {
		return err
	}

	items, err := client.GetAliasTypes(cmd.Context())
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
