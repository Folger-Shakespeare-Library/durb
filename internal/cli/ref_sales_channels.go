package cli

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var refSalesChannelsCmd = &cobra.Command{
	Use:   "sales-channels",
	Short: "Sales channels",
	Long:  "Commands for sales channel reference data.",
}

var refSalesChannelsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List sales channels",
	Long:  "List available sales channels.",
	RunE:  runRefSalesChannelsList,
}

func init() {
	refSalesChannelsCmd.AddCommand(refSalesChannelsListCmd)
}

func runRefSalesChannelsList(cmd *cobra.Command, args []string) error {
	client, err := loadClient()
	if err != nil {
		return err
	}

	items, err := client.GetSalesChannels(cmd.Context())
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
