package cli

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var refContributionImportSetsCmd = &cobra.Command{
	Use:   "contribution-import-sets",
	Short: "Contribution import sets",
	Long:  "Commands for contribution import set reference data.",
}

var refContributionImportSetsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List contribution import sets",
	Long:  "List available contribution import sets.",
	RunE:  runRefContributionImportSetsList,
}

func init() {
	refContributionImportSetsCmd.AddCommand(refContributionImportSetsListCmd)
}

func runRefContributionImportSetsList(cmd *cobra.Command, args []string) error {
	client, err := loadClient()
	if err != nil {
		return err
	}

	items, err := client.GetContributionImportSets(cmd.Context())
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
