package cli

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var refStatesCmd = &cobra.Command{
	Use:   "states",
	Short: "States",
	Long:  "Commands for state/province reference data.",
}

var refStatesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List states",
	Long:  "List available states and provinces.",
	RunE:  runRefStatesList,
}

func init() {
	refStatesCmd.AddCommand(refStatesListCmd)
}

func runRefStatesList(cmd *cobra.Command, args []string) error {
	client, err := loadClient()
	if err != nil {
		return err
	}

	items, err := client.GetStates(cmd.Context())
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
