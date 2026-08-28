package cli

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var txnPricingRuleSetsCmd = &cobra.Command{
	Use:   "pricing-rule-sets",
	Short: "Pricing rule sets",
	Long:  "Commands for pricing rule sets.",
}

var txnPricingRuleSetsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List pricing rule sets",
	Long:  "List available pricing rule sets.",
	RunE:  runTxnPricingRuleSetsList,
}

func init() {
	txnPricingRuleSetsCmd.AddCommand(txnPricingRuleSetsListCmd)
}

func runTxnPricingRuleSetsList(cmd *cobra.Command, args []string) error {
	client, err := loadClient()
	if err != nil {
		return err
	}

	items, err := client.GetPricingRuleSets(cmd.Context())
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
