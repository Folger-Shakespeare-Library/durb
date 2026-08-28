package cli

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var refTNEWCustomizationPointsCmd = &cobra.Command{
	Use:   "tnew-customization-points",
	Short: "TNEW customization points",
	Long:  "Commands for TNEW customization point reference data.",
}

var refTNEWCustomizationPointsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List TNEW customization points",
	Long:  "List available TNEW customization points.",
	RunE:  runRefTNEWCustomizationPointsList,
}

func init() {
	refTNEWCustomizationPointsCmd.AddCommand(refTNEWCustomizationPointsListCmd)
}

func runRefTNEWCustomizationPointsList(cmd *cobra.Command, args []string) error {
	client, err := loadClient()
	if err != nil {
		return err
	}

	items, err := client.GetTNEWCustomizationPoints(cmd.Context())
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
