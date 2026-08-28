package cli

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var refMembershipStandingsCmd = &cobra.Command{
	Use:   "membership-standings",
	Short: "Membership standings",
	Long:  "Commands for membership standing reference data.",
}

var refMembershipStandingsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List membership standings",
	Long:  "List available membership standings.",
	RunE:  runRefMembershipStandingsList,
}

func init() {
	refMembershipStandingsCmd.AddCommand(refMembershipStandingsListCmd)
}

func runRefMembershipStandingsList(cmd *cobra.Command, args []string) error {
	client, err := loadClient()
	if err != nil {
		return err
	}

	items, err := client.GetMembershipStandings(cmd.Context())
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
