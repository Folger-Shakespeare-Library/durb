package cli

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var txnPricingRulesCmd = &cobra.Command{
	Use:   "pricing-rules",
	Short: "Pricing rules",
	Long:  "Commands for pricing rules.",
}

var txnPricingRulesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List pricing rules",
	Long:  "List available pricing rules.",
	RunE:  runTxnPricingRulesList,
}

func init() {
	txnPricingRulesCmd.AddCommand(txnPricingRulesListCmd)
}

func runTxnPricingRulesList(cmd *cobra.Command, args []string) error {
	client, err := loadClient()
	if err != nil {
		return err
	}

	items, err := client.GetPricingRules(cmd.Context())
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
