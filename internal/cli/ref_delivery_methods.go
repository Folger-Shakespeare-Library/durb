package cli

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var refDeliveryMethodsCmd = &cobra.Command{
	Use:   "delivery-methods",
	Short: "Delivery methods",
	Long:  "Commands for delivery method reference data.",
}

var refDeliveryMethodsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List delivery methods",
	Long:  "List available delivery methods.",
	RunE:  runRefDeliveryMethodsList,
}

func init() {
	refDeliveryMethodsCmd.AddCommand(refDeliveryMethodsListCmd)
}

func runRefDeliveryMethodsList(cmd *cobra.Command, args []string) error {
	client, err := loadClient()
	if err != nil {
		return err
	}

	items, err := client.GetDeliveryMethods(cmd.Context())
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
