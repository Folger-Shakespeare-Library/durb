package cli

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var refPriceTypeGroupsCmd = &cobra.Command{
	Use:   "price-type-groups",
	Short: "Price type groups",
	Long:  "Commands for price type group reference data.",
}

var refPriceTypeGroupsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List price type groups",
	Long:  "List available price type groups.",
	RunE:  runRefPriceTypeGroupsList,
}

func init() {
	refPriceTypeGroupsCmd.AddCommand(refPriceTypeGroupsListCmd)
}

func runRefPriceTypeGroupsList(cmd *cobra.Command, args []string) error {
	client, err := loadClient()
	if err != nil {
		return err
	}

	items, err := client.GetPriceTypeGroups(cmd.Context())
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
