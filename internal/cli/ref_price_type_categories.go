package cli

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var refPriceTypeCategoriesCmd = &cobra.Command{
	Use:   "price-type-categories",
	Short: "Price type categories",
	Long:  "Commands for price type category reference data.",
}

var refPriceTypeCategoriesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List price type categories",
	Long:  "List available price type categories.",
	RunE:  runRefPriceTypeCategoriesList,
}

func init() {
	refPriceTypeCategoriesCmd.AddCommand(refPriceTypeCategoriesListCmd)
}

func runRefPriceTypeCategoriesList(cmd *cobra.Command, args []string) error {
	client, err := loadClient()
	if err != nil {
		return err
	}

	items, err := client.GetPriceTypeCategories(cmd.Context())
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
