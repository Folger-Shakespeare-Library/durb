package cli

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var refPlannedGivingOnFilesCmd = &cobra.Command{
	Use:   "planned-giving-on-files",
	Short: "Planned giving on-file statuses",
	Long:  "Commands for planned giving on-file status reference data.",
}

var refPlannedGivingOnFilesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List planned giving on-file statuss",
	Long:  "List available planned giving on-file statuss.",
	RunE:  runRefPlannedGivingOnFilesList,
}

func init() {
	refPlannedGivingOnFilesCmd.AddCommand(refPlannedGivingOnFilesListCmd)
}

func runRefPlannedGivingOnFilesList(cmd *cobra.Command, args []string) error {
	client, err := loadClient()
	if err != nil {
		return err
	}

	items, err := client.GetPlannedGivingOnFiles(cmd.Context())
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
