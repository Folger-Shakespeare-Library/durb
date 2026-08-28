package cli

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var refContactPointPurposeCategoriesCmd = &cobra.Command{
	Use:   "contact-point-purpose-categories",
	Short: "Contact point purpose categories",
	Long:  "Commands for contact point purpose-to-category mapping reference data.",
}

var refContactPointPurposeCategoriesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List contact point purpose categories",
	Long:  "List available contact point purpose-to-category mappings.",
	RunE:  runRefContactPointPurposeCategoriesList,
}

func init() {
	refContactPointPurposeCategoriesCmd.AddCommand(refContactPointPurposeCategoriesListCmd)
}

func runRefContactPointPurposeCategoriesList(cmd *cobra.Command, args []string) error {
	client, err := loadClient()
	if err != nil {
		return err
	}

	items, err := client.GetContactPointPurposeCategories(cmd.Context())
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
