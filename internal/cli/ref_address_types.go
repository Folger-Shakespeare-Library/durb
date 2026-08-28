package cli

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var refAddressTypesCmd = &cobra.Command{
	Use:   "address-types",
	Short: "Address types",
	Long:  "Commands for address type reference data.",
}

var refAddressTypesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List address types",
	Long:  "List available address types.",
	RunE:  runRefAddressTypesList,
}

func init() {
	refAddressTypesCmd.AddCommand(refAddressTypesListCmd)
}

func runRefAddressTypesList(cmd *cobra.Command, args []string) error {
	client, err := loadClient()
	if err != nil {
		return err
	}

	items, err := client.GetAddressTypes(cmd.Context())
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
