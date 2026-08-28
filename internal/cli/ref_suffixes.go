package cli

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var refSuffixesCmd = &cobra.Command{
	Use:   "suffixes",
	Short: "Suffixes",
	Long:  "Commands for name suffix reference data.",
}

var refSuffixesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List suffixes",
	Long:  "List available name suffixes.",
	RunE:  runRefSuffixesList,
}

func init() {
	refSuffixesCmd.AddCommand(refSuffixesListCmd)
}

func runRefSuffixesList(cmd *cobra.Command, args []string) error {
	client, err := loadClient()
	if err != nil {
		return err
	}

	items, err := client.GetSuffixes(cmd.Context())
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
