package cli

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var refMembershipStatusesCmd = &cobra.Command{
	Use:   "membership-statuses",
	Short: "Membership statuses",
	Long:  "Commands for membership status reference data.",
}

var refMembershipStatusesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List membership statuses",
	Long:  "List available membership statuses.",
	RunE:  runRefMembershipStatusesList,
}

func init() {
	refMembershipStatusesCmd.AddCommand(refMembershipStatusesListCmd)
}

func runRefMembershipStatusesList(cmd *cobra.Command, args []string) error {
	client, err := loadClient()
	if err != nil {
		return err
	}

	items, err := client.GetMembershipStatuses(cmd.Context())
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
