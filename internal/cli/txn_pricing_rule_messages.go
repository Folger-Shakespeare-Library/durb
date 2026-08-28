package cli

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var txnPricingRuleMessagesCmd = &cobra.Command{
	Use:   "pricing-rule-messages",
	Short: "Pricing rule messages",
	Long:  "Commands for pricing rule messages.",
}

var txnPricingRuleMessagesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List pricing rule messages",
	Long:  "List available pricing rule messages.",
	RunE:  runTxnPricingRuleMessagesList,
}

func init() {
	txnPricingRuleMessagesCmd.AddCommand(txnPricingRuleMessagesListCmd)
}

func runTxnPricingRuleMessagesList(cmd *cobra.Command, args []string) error {
	client, err := loadClient()
	if err != nil {
		return err
	}

	items, err := client.GetPricingRuleMessages(cmd.Context())
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
