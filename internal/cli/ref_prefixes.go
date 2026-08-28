package cli

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var refPrefixesCmd = &cobra.Command{
	Use:   "prefixes",
	Short: "Prefixes",
	Long:  "Commands for name prefix reference data.",
}

var refPrefixesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List prefixes",
	Long:  "List available name prefixes.",
	RunE:  runRefPrefixesList,
}

func init() {
	refPrefixesCmd.AddCommand(refPrefixesListCmd)
}

func runRefPrefixesList(cmd *cobra.Command, args []string) error {
	client, err := loadClient()
	if err != nil {
		return err
	}

	items, err := client.GetPrefixes(cmd.Context())
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
