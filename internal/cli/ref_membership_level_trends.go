package cli

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var refMembershipLevelTrendsCmd = &cobra.Command{
	Use:   "membership-level-trends",
	Short: "Membership level trends",
	Long:  "Commands for membership level trend reference data.",
}

var refMembershipLevelTrendsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List membership level trends",
	Long:  "List available membership level trends.",
	RunE:  runRefMembershipLevelTrendsList,
}

func init() {
	refMembershipLevelTrendsCmd.AddCommand(refMembershipLevelTrendsListCmd)
}

func runRefMembershipLevelTrendsList(cmd *cobra.Command, args []string) error {
	client, err := loadClient()
	if err != nil {
		return err
	}

	items, err := client.GetMembershipLevelTrends(cmd.Context())
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
