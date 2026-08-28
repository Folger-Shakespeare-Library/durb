package cli

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var refPackageTypesCmd = &cobra.Command{
	Use:   "package-types",
	Short: "Package types",
	Long:  "Commands for package type reference data.",
}

var refPackageTypesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List package types",
	Long:  "List available package types.",
	RunE:  runRefPackageTypesList,
}

func init() {
	refPackageTypesCmd.AddCommand(refPackageTypesListCmd)
}

func runRefPackageTypesList(cmd *cobra.Command, args []string) error {
	client, err := loadClient()
	if err != nil {
		return err
	}

	items, err := client.GetPackageTypes(cmd.Context())
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
