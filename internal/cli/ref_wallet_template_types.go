package cli

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var refWalletTemplateTypesCmd = &cobra.Command{
	Use:   "wallet-template-types",
	Short: "Wallet template types",
	Long:  "Commands for wallet template type reference data.",
}

var refWalletTemplateTypesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List wallet template types",
	Long:  "List available wallet template types.",
	RunE:  runRefWalletTemplateTypesList,
}

func init() {
	refWalletTemplateTypesCmd.AddCommand(refWalletTemplateTypesListCmd)
}

func runRefWalletTemplateTypesList(cmd *cobra.Command, args []string) error {
	client, err := loadClient()
	if err != nil {
		return err
	}

	items, err := client.GetWalletTemplateTypes(cmd.Context())
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
