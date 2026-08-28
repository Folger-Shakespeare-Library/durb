package cli

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var refLoginTypesCmd = &cobra.Command{
	Use:   "login-types",
	Short: "Login types",
	Long:  "Commands for login type reference data.",
}

var refLoginTypesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List login types",
	Long:  "List available login types.",
	RunE:  runRefLoginTypesList,
}

func init() {
	refLoginTypesCmd.AddCommand(refLoginTypesListCmd)
}

func runRefLoginTypesList(cmd *cobra.Command, args []string) error {
	client, err := loadClient()
	if err != nil {
		return err
	}

	items, err := client.GetLoginTypes(cmd.Context())
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
