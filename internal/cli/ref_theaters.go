package cli

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var refTheatersCmd = &cobra.Command{
	Use:   "theaters",
	Short: "Theaters",
	Long:  "Commands for theater (venue) reference data.",
}

var refTheatersListCmd = &cobra.Command{
	Use:   "list",
	Short: "List theaters",
	Long:  "List available theaters (venues).",
	RunE:  runRefTheatersList,
}

func init() {
	refTheatersCmd.AddCommand(refTheatersListCmd)
}

func runRefTheatersList(cmd *cobra.Command, args []string) error {
	client, err := loadClient()
	if err != nil {
		return err
	}

	items, err := client.GetTheaters(cmd.Context())
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
