package cli

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var refKeywordConstituentTypesCmd = &cobra.Command{
	Use:   "keyword-constituent-types",
	Short: "Keyword constituent types",
	Long:  "Commands for keyword-to-constituent-type mapping reference data.",
}

var refKeywordConstituentTypesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List keyword constituent types",
	Long:  "List available keyword-to-constituent-type mappings.",
	RunE:  runRefKeywordConstituentTypesList,
}

func init() {
	refKeywordConstituentTypesCmd.AddCommand(refKeywordConstituentTypesListCmd)
}

func runRefKeywordConstituentTypesList(cmd *cobra.Command, args []string) error {
	client, err := loadClient()
	if err != nil {
		return err
	}

	items, err := client.GetKeywordConstituentTypes(cmd.Context())
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
