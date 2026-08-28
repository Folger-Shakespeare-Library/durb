package cli

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var refMembershipBenefitTypesCmd = &cobra.Command{
	Use:   "membership-benefit-types",
	Short: "Membership benefit types",
	Long:  "Commands for membership benefit type reference data.",
}

var refMembershipBenefitTypesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List membership benefit types",
	Long:  "List available membership benefit types.",
	RunE:  runRefMembershipBenefitTypesList,
}

func init() {
	refMembershipBenefitTypesCmd.AddCommand(refMembershipBenefitTypesListCmd)
}

func runRefMembershipBenefitTypesList(cmd *cobra.Command, args []string) error {
	client, err := loadClient()
	if err != nil {
		return err
	}

	items, err := client.GetMembershipBenefitTypes(cmd.Context())
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
