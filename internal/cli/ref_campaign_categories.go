package cli

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var refCampaignCategoriesCmd = &cobra.Command{
	Use:   "campaign-categories",
	Short: "Campaign categories",
	Long:  "Commands for campaign category reference data.",
}

var refCampaignCategoriesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List campaign categorys",
	Long:  "List available campaign categorys.",
	RunE:  runRefCampaignCategoriesList,
}

func init() {
	refCampaignCategoriesCmd.AddCommand(refCampaignCategoriesListCmd)
}

func runRefCampaignCategoriesList(cmd *cobra.Command, args []string) error {
	client, err := loadClient()
	if err != nil {
		return err
	}

	items, err := client.GetCampaignCategories(cmd.Context())
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
