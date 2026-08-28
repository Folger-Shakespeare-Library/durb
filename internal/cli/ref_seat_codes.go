package cli

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var refSeatCodesCmd = &cobra.Command{
	Use:   "seat-codes",
	Short: "Seat codes",
	Long:  "Commands for seat code reference data.",
}

var refSeatCodesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List seat codes",
	Long:  "List available seat codes.",
	RunE:  runRefSeatCodesList,
}

func init() {
	refSeatCodesCmd.AddCommand(refSeatCodesListCmd)
}

func runRefSeatCodesList(cmd *cobra.Command, args []string) error {
	client, err := loadClient()
	if err != nil {
		return err
	}

	items, err := client.GetSeatCodes(cmd.Context())
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
