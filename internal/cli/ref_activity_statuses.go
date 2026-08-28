package cli

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var refActivityStatusesCmd = &cobra.Command{
	Use:   "activity-statuses",
	Short: "Activity statuses",
	Long:  "Commands for special activity status reference data.",
}

var refActivityStatusesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List activity statuses",
	Long:  "List available special activity statuses.",
	RunE:  runRefActivityStatusesList,
}

func init() {
	refActivityStatusesCmd.AddCommand(refActivityStatusesListCmd)
}

func runRefActivityStatusesList(cmd *cobra.Command, args []string) error {
	client, err := loadClient()
	if err != nil {
		return err
	}

	items, err := client.GetSpecialActivityStatuses(cmd.Context())
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
