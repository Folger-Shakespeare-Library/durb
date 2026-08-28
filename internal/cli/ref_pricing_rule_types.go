package cli

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var refPricingRuleTypesCmd = &cobra.Command{
	Use:   "pricing-rule-types",
	Short: "Pricing rule types",
	Long:  "Commands for pricing rule type reference data.",
}

var refPricingRuleTypesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List pricing rule types",
	Long:  "List available pricing rule types.",
	RunE:  runRefPricingRuleTypesList,
}

func init() {
	refPricingRuleTypesCmd.AddCommand(refPricingRuleTypesListCmd)
}

func runRefPricingRuleTypesList(cmd *cobra.Command, args []string) error {
	client, err := loadClient()
	if err != nil {
		return err
	}

	items, err := client.GetPricingRuleTypes(cmd.Context())
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
