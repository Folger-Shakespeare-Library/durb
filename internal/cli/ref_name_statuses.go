package cli

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var refNameStatusesCmd = &cobra.Command{
	Use:   "name-statuses",
	Short: "Name statuses",
	Long:  "Commands for name status reference data.",
}

var refNameStatusesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List name statuses",
	Long:  "List available name statuses for constituents.",
	RunE:  runRefNameStatusesList,
}

func init() {
	refNameStatusesCmd.AddCommand(refNameStatusesListCmd)
}

func runRefNameStatusesList(cmd *cobra.Command, args []string) error {
	client, err := loadClient()
	if err != nil {
		return err
	}

	items, err := client.GetNameStatuses(cmd.Context())
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
