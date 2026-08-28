package cli

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var refPricingRuleMessageTypesCmd = &cobra.Command{
	Use:   "pricing-rule-message-types",
	Short: "Pricing rule message types",
	Long:  "Commands for pricing rule message type reference data.",
}

var refPricingRuleMessageTypesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List pricing rule message types",
	Long:  "List available pricing rule message types.",
	RunE:  runRefPricingRuleMessageTypesList,
}

func init() {
	refPricingRuleMessageTypesCmd.AddCommand(refPricingRuleMessageTypesListCmd)
}

func runRefPricingRuleMessageTypesList(cmd *cobra.Command, args []string) error {
	client, err := loadClient()
	if err != nil {
		return err
	}

	items, err := client.GetPricingRuleMessageTypes(cmd.Context())
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
