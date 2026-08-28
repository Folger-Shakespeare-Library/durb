package cli

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var refPaymentTypesCmd = &cobra.Command{
	Use:   "payment-types",
	Short: "Payment types",
	Long:  "Commands for payment type reference data.",
}

var refPaymentTypesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List payment types",
	Long:  "List available payment types.",
	RunE:  runRefPaymentTypesList,
}

func init() {
	refPaymentTypesCmd.AddCommand(refPaymentTypesListCmd)
}

func runRefPaymentTypesList(cmd *cobra.Command, args []string) error {
	client, err := loadClient()
	if err != nil {
		return err
	}

	items, err := client.GetPaymentTypes(cmd.Context())
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
