package cli

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var refAffiliationTypesCmd = &cobra.Command{
	Use:   "affiliation-types",
	Short: "Affiliation types",
	Long:  "Commands for affiliation type reference data.",
}

var refAffiliationTypesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List affiliation types",
	Long:  "List available affiliation types.",
	RunE:  runRefAffiliationTypesList,
}

func init() {
	refAffiliationTypesCmd.AddCommand(refAffiliationTypesListCmd)
}

func runRefAffiliationTypesList(cmd *cobra.Command, args []string) error {
	client, err := loadClient()
	if err != nil {
		return err
	}

	items, err := client.GetAffiliationTypes(cmd.Context())
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
