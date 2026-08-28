package cli

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var txnPricingRuleSetPricingRulesCmd = &cobra.Command{
	Use:   "pricing-rule-set-pricing-rules",
	Short: "Pricing rule set pricing rules",
	Long:  "Commands for pricing rule set pricing rule mappings.",
}

var txnPricingRuleSetPricingRulesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List pricing rule set pricing rules",
	Long:  "List pricing rule set to pricing rule mappings.",
	RunE:  runTxnPricingRuleSetPricingRulesList,
}

func init() {
	txnPricingRuleSetPricingRulesCmd.AddCommand(txnPricingRuleSetPricingRulesListCmd)
}

func runTxnPricingRuleSetPricingRulesList(cmd *cobra.Command, args []string) error {
	client, err := loadClient()
	if err != nil {
		return err
	}

	items, err := client.GetPricingRuleSetPricingRules(cmd.Context())
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
