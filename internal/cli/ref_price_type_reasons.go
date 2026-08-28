package cli

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var refPriceTypeReasonsCmd = &cobra.Command{
	Use:   "price-type-reasons",
	Short: "Price type reasons",
	Long:  "Commands for price type reason reference data.",
}

var refPriceTypeReasonsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List price type reasons",
	Long:  "List available price type reasons.",
	RunE:  runRefPriceTypeReasonsList,
}

func init() {
	refPriceTypeReasonsCmd.AddCommand(refPriceTypeReasonsListCmd)
}

func runRefPriceTypeReasonsList(cmd *cobra.Command, args []string) error {
	client, err := loadClient()
	if err != nil {
		return err
	}

	items, err := client.GetPriceTypeReasons(cmd.Context())
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
