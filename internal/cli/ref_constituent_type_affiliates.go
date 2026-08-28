package cli

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var refConstituentTypeAffiliatesCmd = &cobra.Command{
	Use:   "constituent-type-affiliates",
	Short: "Constituent type affiliates",
	Long:  "Commands for constituent type affiliate reference data.",
}

var refConstituentTypeAffiliatesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List constituent type affiliates",
	Long:  "List available constituent type affiliate definitions.",
	RunE:  runRefConstituentTypeAffiliatesList,
}

func init() {
	refConstituentTypeAffiliatesCmd.AddCommand(refConstituentTypeAffiliatesListCmd)
}

func runRefConstituentTypeAffiliatesList(cmd *cobra.Command, args []string) error {
	client, err := loadClient()
	if err != nil {
		return err
	}

	items, err := client.GetConstituentTypeAffiliates(cmd.Context())
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
