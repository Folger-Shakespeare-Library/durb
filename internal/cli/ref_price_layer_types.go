package cli

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var refPriceLayerTypesCmd = &cobra.Command{
	Use:   "price-layer-types",
	Short: "Price layer types",
	Long:  "Commands for price layer type reference data.",
}

var refPriceLayerTypesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List price layer types",
	Long:  "List available price layer types.",
	RunE:  runRefPriceLayerTypesList,
}

func init() {
	refPriceLayerTypesCmd.AddCommand(refPriceLayerTypesListCmd)
}

func runRefPriceLayerTypesList(cmd *cobra.Command, args []string) error {
	client, err := loadClient()
	if err != nil {
		return err
	}

	items, err := client.GetPriceLayerTypes(cmd.Context())
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
