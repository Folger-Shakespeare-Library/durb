package cli

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var refWalletTemplatesCmd = &cobra.Command{
	Use:   "wallet-templates",
	Short: "Wallet templates",
	Long:  "Commands for wallet template reference data.",
}

var refWalletTemplatesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List wallet templates",
	Long:  "List available wallet templates.",
	RunE:  runRefWalletTemplatesList,
}

func init() {
	refWalletTemplatesCmd.AddCommand(refWalletTemplatesListCmd)
}

func runRefWalletTemplatesList(cmd *cobra.Command, args []string) error {
	client, err := loadClient()
	if err != nil {
		return err
	}

	items, err := client.GetWalletTemplates(cmd.Context())
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
