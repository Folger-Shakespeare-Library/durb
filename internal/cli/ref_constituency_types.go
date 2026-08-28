package cli

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var refConstituencyTypesCmd = &cobra.Command{
	Use:   "constituency-types",
	Short: "Constituency types",
	Long:  "Commands for constituency type reference data.",
}

var refConstituencyTypesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List constituency types",
	Long:  "List available constituency types.",
	RunE:  runRefConstituencyTypesList,
}

func init() {
	refConstituencyTypesCmd.AddCommand(refConstituencyTypesListCmd)
}

func runRefConstituencyTypesList(cmd *cobra.Command, args []string) error {
	client, err := loadClient()
	if err != nil {
		return err
	}

	items, err := client.GetConstituencyTypes(cmd.Context())
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
