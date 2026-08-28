package cli

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var refConstituentProtectionTypesCmd = &cobra.Command{
	Use:   "constituent-protection-types",
	Short: "Constituent protection types",
	Long:  "Commands for constituent protection type reference data.",
}

var refConstituentProtectionTypesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List constituent protection types",
	Long:  "List available constituent protection types.",
	RunE:  runRefConstituentProtectionTypesList,
}

func init() {
	refConstituentProtectionTypesCmd.AddCommand(refConstituentProtectionTypesListCmd)
}

func runRefConstituentProtectionTypesList(cmd *cobra.Command, args []string) error {
	client, err := loadClient()
	if err != nil {
		return err
	}

	items, err := client.GetConstituentProtectionTypes(cmd.Context())
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
