package cli

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var refTNEWCustomizationsCmd = &cobra.Command{
	Use:   "tnew-customizations",
	Short: "TNEW customizations",
	Long:  "Commands for TNEW customization reference data.",
}

var refTNEWCustomizationsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List TNEW customizations",
	Long:  "List available TNEW customizations.",
	RunE:  runRefTNEWCustomizationsList,
}

func init() {
	refTNEWCustomizationsCmd.AddCommand(refTNEWCustomizationsListCmd)
}

func runRefTNEWCustomizationsList(cmd *cobra.Command, args []string) error {
	client, err := loadClient()
	if err != nil {
		return err
	}

	items, err := client.GetTNEWCustomizations(cmd.Context())
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
