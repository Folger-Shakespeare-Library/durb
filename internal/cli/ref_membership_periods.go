package cli

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var refMembershipPeriodsCmd = &cobra.Command{
	Use:   "membership-periods",
	Short: "Membership periods",
	Long:  "Commands for membership period reference data.",
}

var refMembershipPeriodsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List membership periods",
	Long:  "List available membership periods.",
	RunE:  runRefMembershipPeriodsList,
}

func init() {
	refMembershipPeriodsCmd.AddCommand(refMembershipPeriodsListCmd)
}

func runRefMembershipPeriodsList(cmd *cobra.Command, args []string) error {
	client, err := loadClient()
	if err != nil {
		return err
	}

	items, err := client.GetMembershipPeriods(cmd.Context())
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
