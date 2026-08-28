package cli

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var refElectronicAddressTypesCmd = &cobra.Command{
	Use:   "electronic-address-types",
	Short: "Electronic address types",
	Long:  "Commands for electronic address type reference data.",
}

var refElectronicAddressTypesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List electronic address types",
	Long:  "List available electronic address types (email, phone, web).",
	RunE:  runRefElectronicAddressTypesList,
}

func init() {
	refElectronicAddressTypesCmd.AddCommand(refElectronicAddressTypesListCmd)
}

func runRefElectronicAddressTypesList(cmd *cobra.Command, args []string) error {
	client, err := loadClient()
	if err != nil {
		return err
	}

	items, err := client.GetElectronicAddressTypes(cmd.Context())
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
