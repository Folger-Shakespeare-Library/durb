package cli

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var refOrderCategoriesCmd = &cobra.Command{
	Use:   "order-categories",
	Short: "Order categories",
	Long:  "Commands for order category reference data.",
}

var refOrderCategoriesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List order categories",
	Long:  "List available order categories.",
	RunE:  runRefOrderCategoriesList,
}

func init() {
	refOrderCategoriesCmd.AddCommand(refOrderCategoriesListCmd)
}

func runRefOrderCategoriesList(cmd *cobra.Command, args []string) error {
	client, err := loadClient()
	if err != nil {
		return err
	}

	items, err := client.GetOrderCategories(cmd.Context())
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
