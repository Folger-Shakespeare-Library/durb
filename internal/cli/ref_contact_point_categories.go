package cli

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var refContactPointCategoriesCmd = &cobra.Command{
	Use:   "contact-point-categories",
	Short: "Contact point categories",
	Long:  "Commands for contact point category reference data.",
}

var refContactPointCategoriesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List contact point categories",
	Long:  "List available contact point categories.",
	RunE:  runRefContactPointCategoriesList,
}

func init() {
	refContactPointCategoriesCmd.AddCommand(refContactPointCategoriesListCmd)
}

func runRefContactPointCategoriesList(cmd *cobra.Command, args []string) error {
	client, err := loadClient()
	if err != nil {
		return err
	}

	items, err := client.GetContactPointCategories(cmd.Context())
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
