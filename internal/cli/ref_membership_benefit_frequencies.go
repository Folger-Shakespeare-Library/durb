package cli

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var refMembershipBenefitFrequenciesCmd = &cobra.Command{
	Use:   "membership-benefit-frequencies",
	Short: "Membership benefit frequencies",
	Long:  "Commands for membership benefit frequency reference data.",
}

var refMembershipBenefitFrequenciesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List membership benefit frequencies",
	Long:  "List available membership benefit frequencies.",
	RunE:  runRefMembershipBenefitFrequenciesList,
}

func init() {
	refMembershipBenefitFrequenciesCmd.AddCommand(refMembershipBenefitFrequenciesListCmd)
}

func runRefMembershipBenefitFrequenciesList(cmd *cobra.Command, args []string) error {
	client, err := loadClient()
	if err != nil {
		return err
	}

	items, err := client.GetMembershipBenefitFrequencies(cmd.Context())
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
