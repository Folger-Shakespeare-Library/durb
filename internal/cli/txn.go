package cli

import (
	"github.com/spf13/cobra"
)

var txnCmd = &cobra.Command{
	Use:   "txn",
	Short: "Transaction data",
	Long:  "Commands for Tessitura transaction data (pricing rules, rule sets).",
}

func init() {
	txnCmd.AddCommand(txnPricingRuleMessagesCmd)
	txnCmd.AddCommand(txnPricingRulesCmd)
	txnCmd.AddCommand(txnPricingRuleSetPricingRulesCmd)
	txnCmd.AddCommand(txnPricingRuleSetsCmd)
}
