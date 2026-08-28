package cli

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var refDonationLevelsCmd = &cobra.Command{
	Use:   "donation-levels",
	Short: "Donation levels",
	Long:  "Commands for donation level reference data.",
}

var refDonationLevelsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List donation levels",
	Long:  "List available donation levels.",
	RunE:  runRefDonationLevelsList,
}

func init() {
	refDonationLevelsCmd.AddCommand(refDonationLevelsListCmd)
}

func runRefDonationLevelsList(cmd *cobra.Command, args []string) error {
	client, err := loadClient()
	if err != nil {
		return err
	}

	items, err := client.GetDonationLevels(cmd.Context())
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
