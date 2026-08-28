package cli

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var refBusinessUnitsCmd = &cobra.Command{
	Use:   "business-units",
	Short: "Business units",
	Long:  "Commands for business unit reference data.",
}

var refBusinessUnitsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List business units",
	Long:  "List available business units.",
	RunE:  runRefBusinessUnitsList,
}

func init() {
	refBusinessUnitsCmd.AddCommand(refBusinessUnitsListCmd)
}

func runRefBusinessUnitsList(cmd *cobra.Command, args []string) error {
	client, err := loadClient()
	if err != nil {
		return err
	}

	items, err := client.GetBusinessUnits(cmd.Context())
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
