package cli

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var refContactPermissionCategoriesCmd = &cobra.Command{
	Use:   "contact-permission-categories",
	Short: "Contact permission categories",
	Long:  "Commands for contact permission category reference data.",
}

var refContactPermissionCategoriesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List contact permission categories",
	Long:  "List available contact permission categories.",
	RunE:  runRefContactPermissionCategoriesList,
}

func init() {
	refContactPermissionCategoriesCmd.AddCommand(refContactPermissionCategoriesListCmd)
}

func runRefContactPermissionCategoriesList(cmd *cobra.Command, args []string) error {
	client, err := loadClient()
	if err != nil {
		return err
	}

	items, err := client.GetContactPermissionCategories(cmd.Context())
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
