package cli

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var refInterestTypesCmd = &cobra.Command{
	Use:   "interest-types",
	Short: "Interest types",
	Long:  "Commands for interest type reference data.",
}

var refInterestTypesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List interest types",
	Long:  "List available interest types.",
	RunE:  runRefInterestTypesList,
}

func init() {
	refInterestTypesCmd.AddCommand(refInterestTypesListCmd)
}

func runRefInterestTypesList(cmd *cobra.Command, args []string) error {
	client, err := loadClient()
	if err != nil {
		return err
	}

	items, err := client.GetInterestTypes(cmd.Context())
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
