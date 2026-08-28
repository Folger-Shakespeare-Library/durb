package cli

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var refContactPointCategoryPurposesCmd = &cobra.Command{
	Use:   "contact-point-category-purposes",
	Short: "Contact point category purposes",
	Long:  "Commands for contact point category-to-purpose mapping reference data.",
}

var refContactPointCategoryPurposesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List contact point category purposes",
	Long:  "List available contact point category-to-purpose mappings.",
	RunE:  runRefContactPointCategoryPurposesList,
}

func init() {
	refContactPointCategoryPurposesCmd.AddCommand(refContactPointCategoryPurposesListCmd)
}

func runRefContactPointCategoryPurposesList(cmd *cobra.Command, args []string) error {
	client, err := loadClient()
	if err != nil {
		return err
	}

	items, err := client.GetContactPointCategoryPurposes(cmd.Context())
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
