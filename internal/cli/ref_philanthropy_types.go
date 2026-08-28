package cli

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var refPhilanthropyTypesCmd = &cobra.Command{
	Use:   "philanthropy-types",
	Short: "Philanthropy types",
	Long:  "Commands for philanthropy type reference data.",
}

var refPhilanthropyTypesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List philanthropy types",
	Long:  "List available philanthropy types.",
	RunE:  runRefPhilanthropyTypesList,
}

func init() {
	refPhilanthropyTypesCmd.AddCommand(refPhilanthropyTypesListCmd)
}

func runRefPhilanthropyTypesList(cmd *cobra.Command, args []string) error {
	client, err := loadClient()
	if err != nil {
		return err
	}

	items, err := client.GetPhilanthropyTypes(cmd.Context())
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
