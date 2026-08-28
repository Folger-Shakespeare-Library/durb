package cli

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var refPricingRuleCategoriesCmd = &cobra.Command{
	Use:   "pricing-rule-categories",
	Short: "Pricing rule categories",
	Long:  "Commands for pricing rule category reference data.",
}

var refPricingRuleCategoriesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List pricing rule categories",
	Long:  "List available pricing rule categories.",
	RunE:  runRefPricingRuleCategoriesList,
}

func init() {
	refPricingRuleCategoriesCmd.AddCommand(refPricingRuleCategoriesListCmd)
}

func runRefPricingRuleCategoriesList(cmd *cobra.Command, args []string) error {
	client, err := loadClient()
	if err != nil {
		return err
	}

	items, err := client.GetPricingRuleCategories(cmd.Context())
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
