package cli

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var refPriceCategoriesCmd = &cobra.Command{
	Use:   "price-categories",
	Short: "Price categories",
	Long:  "Commands for price category reference data.",
}

var refPriceCategoriesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List price categories",
	Long:  "List available price categories.",
	RunE:  runRefPriceCategoriesList,
}

func init() {
	refPriceCategoriesCmd.AddCommand(refPriceCategoriesListCmd)
}

func runRefPriceCategoriesList(cmd *cobra.Command, args []string) error {
	client, err := loadClient()
	if err != nil {
		return err
	}

	items, err := client.GetPriceCategories(cmd.Context())
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
