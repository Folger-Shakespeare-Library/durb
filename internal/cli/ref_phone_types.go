package cli

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var refPhoneTypesCmd = &cobra.Command{
	Use:   "phone-types",
	Short: "Phone types",
	Long:  "Commands for phone type reference data.",
}

var refPhoneTypesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List phone types",
	Long:  "List available phone types.",
	RunE:  runRefPhoneTypesList,
}

func init() {
	refPhoneTypesCmd.AddCommand(refPhoneTypesListCmd)
}

func runRefPhoneTypesList(cmd *cobra.Command, args []string) error {
	client, err := loadClient()
	if err != nil {
		return err
	}

	items, err := client.GetPhoneTypes(cmd.Context())
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
