package cli

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var refSubLineItemStatusesCmd = &cobra.Command{
	Use:   "sub-line-item-statuses",
	Short: "Sub-line item statuses",
	Long:  "Commands for sub-line item status reference data.",
}

var refSubLineItemStatusesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List sub-line item statuses",
	Long:  "List available sub-line item statuses.",
	RunE:  runRefSubLineItemStatusesList,
}

func init() {
	refSubLineItemStatusesCmd.AddCommand(refSubLineItemStatusesListCmd)
}

func runRefSubLineItemStatusesList(cmd *cobra.Command, args []string) error {
	client, err := loadClient()
	if err != nil {
		return err
	}

	items, err := client.GetSubLineItemStatuses(cmd.Context())
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
