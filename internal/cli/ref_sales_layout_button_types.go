package cli

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var refSalesLayoutButtonTypesCmd = &cobra.Command{
	Use:   "sales-layout-button-types",
	Short: "Sales layout button types",
	Long:  "Commands for sales layout button type reference data.",
}

var refSalesLayoutButtonTypesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List sales layout button types",
	Long:  "List available sales layout button types.",
	RunE:  runRefSalesLayoutButtonTypesList,
}

func init() {
	refSalesLayoutButtonTypesCmd.AddCommand(refSalesLayoutButtonTypesListCmd)
}

func runRefSalesLayoutButtonTypesList(cmd *cobra.Command, args []string) error {
	client, err := loadClient()
	if err != nil {
		return err
	}

	items, err := client.GetSalesLayoutButtonTypes(cmd.Context())
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
