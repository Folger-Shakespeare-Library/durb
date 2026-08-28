package cli

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var refContributionDesignationsCmd = &cobra.Command{
	Use:   "contribution-designations",
	Short: "Contribution designations",
	Long:  "Commands for contribution designation reference data.",
}

var refContributionDesignationsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List contribution designations",
	Long:  "List available contribution designations.",
	RunE:  runRefContributionDesignationsList,
}

func init() {
	refContributionDesignationsCmd.AddCommand(refContributionDesignationsListCmd)
}

func runRefContributionDesignationsList(cmd *cobra.Command, args []string) error {
	client, err := loadClient()
	if err != nil {
		return err
	}

	items, err := client.GetContributionDesignations(cmd.Context())
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
