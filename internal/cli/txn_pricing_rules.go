package cli

import (
	"encoding/json"
	"fmt"

	"github.com/Folger-Shakespeare-Library/durb/pkg/tessitura"
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

var (
	txnPricingRulesListPerformanceIDs string
	txnPricingRulesListPackageIDs     string
	txnPricingRulesListOrderDate      string
	txnPricingRulesListModeOfSaleID   int
)

func init() {
	txnPricingRulesListCmd.Flags().StringVar(&txnPricingRulesListPerformanceIDs, "performance-ids", "", "filter by performance IDs (comma-delimited)")
	txnPricingRulesListCmd.Flags().StringVar(&txnPricingRulesListPackageIDs, "package-ids", "", "filter by package IDs (comma-delimited)")
	txnPricingRulesListCmd.Flags().StringVar(&txnPricingRulesListOrderDate, "order-date", "", "filter by order date")
	txnPricingRulesListCmd.Flags().IntVar(&txnPricingRulesListModeOfSaleID, "mode-of-sale-id", 0, "filter by mode of sale ID")
	txnPricingRulesCmd.AddCommand(txnPricingRulesListCmd)
}

func runTxnPricingRulesList(cmd *cobra.Command, args []string) error {
	client, err := loadClient()
	if err != nil {
		return err
	}

	var params *tessitura.GetPricingRulesParams
	if cmd.Flags().Changed("performance-ids") || cmd.Flags().Changed("package-ids") || cmd.Flags().Changed("order-date") || cmd.Flags().Changed("mode-of-sale-id") {
		params = &tessitura.GetPricingRulesParams{
			PerformanceIDs: txnPricingRulesListPerformanceIDs,
			PackageIDs:     txnPricingRulesListPackageIDs,
			OrderDate:      txnPricingRulesListOrderDate,
		}
		if cmd.Flags().Changed("mode-of-sale-id") {
			params.ModeOfSaleID = &txnPricingRulesListModeOfSaleID
		}
	}

	items, err := client.GetPricingRules(cmd.Context(), params)
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
