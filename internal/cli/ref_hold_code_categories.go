package cli

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var refHoldCodeCategoriesCmd = &cobra.Command{
	Use:   "hold-code-categories",
	Short: "Hold code categories",
	Long:  "Commands for hold code category reference data.",
}

var refHoldCodeCategoriesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List hold code categories",
	Long:  "List available hold code categories.",
	RunE:  runRefHoldCodeCategoriesList,
}

func init() {
	refHoldCodeCategoriesCmd.AddCommand(refHoldCodeCategoriesListCmd)
}

func runRefHoldCodeCategoriesList(cmd *cobra.Command, args []string) error {
	client, err := loadClient()
	if err != nil {
		return err
	}

	items, err := client.GetHoldCodeCategories(cmd.Context())
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
