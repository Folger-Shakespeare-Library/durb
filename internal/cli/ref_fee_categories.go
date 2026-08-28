package cli

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var refFeeCategoriesCmd = &cobra.Command{
	Use:   "fee-categories",
	Short: "Fee categories",
	Long:  "Commands for fee category reference data.",
}

var refFeeCategoriesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List fee categories",
	Long:  "List available fee categories.",
	RunE:  runRefFeeCategoriesList,
}

func init() {
	refFeeCategoriesCmd.AddCommand(refFeeCategoriesListCmd)
}

func runRefFeeCategoriesList(cmd *cobra.Command, args []string) error {
	client, err := loadClient()
	if err != nil {
		return err
	}

	items, err := client.GetFeeCategories(cmd.Context())
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
