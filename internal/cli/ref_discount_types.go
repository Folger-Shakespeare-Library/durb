package cli

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var refDiscountTypesCmd = &cobra.Command{
	Use:   "discount-types",
	Short: "Discount types",
	Long:  "Commands for discount type reference data.",
}

var refDiscountTypesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List discount types",
	Long:  "List available discount types.",
	RunE:  runRefDiscountTypesList,
}

func init() {
	refDiscountTypesCmd.AddCommand(refDiscountTypesListCmd)
}

func runRefDiscountTypesList(cmd *cobra.Command, args []string) error {
	client, err := loadClient()
	if err != nil {
		return err
	}

	items, err := client.GetDiscountTypes(cmd.Context())
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
