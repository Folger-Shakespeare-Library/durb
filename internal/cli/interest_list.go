package cli

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var interestListFlags struct {
	constituentId int
}

var interestListCmd = &cobra.Command{
	Use:   "list",
	Short: "List interests for a constituent",
	Long: `List all interest assignments for a constituent.

Examples:
  tess interest list --constituent-id 446106`,
	RunE: runInterestList,
}

func init() {
	interestListCmd.Flags().IntVar(&interestListFlags.constituentId, "constituent-id", 0, "constituent ID (required)")
	interestListCmd.MarkFlagRequired("constituent-id")
}

func runInterestList(cmd *cobra.Command, args []string) error {
	client, err := loadClient()
	if err != nil {
		return err
	}

	interests, err := client.GetInterests(cmd.Context(), interestListFlags.constituentId)
	if err != nil {
		return fmt.Errorf("unable to fetch interests: %w", err)
	}

	out, err := json.MarshalIndent(interests, "", "  ")
	if err != nil {
		return fmt.Errorf("unable to format output: %w", err)
	}

	fmt.Fprintln(cmd.OutOrStdout(), string(out))
	return nil
}
