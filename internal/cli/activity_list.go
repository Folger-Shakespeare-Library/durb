package cli

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var activityListFlags struct {
	constituentId  int
	activityTypeId int
}

var activityListCmd = &cobra.Command{
	Use:   "list",
	Short: "List activities for a constituent",
	Long: `List special activities for a constituent. Optionally filter by activity type.

Examples:
  tess activity list --constituent-id 446106
  tess activity list --constituent-id 446106 --activity-type-id 22`,
	RunE: runActivityList,
}

func init() {
	f := activityListCmd.Flags()
	f.IntVar(&activityListFlags.constituentId, "constituent-id", 0, "constituent ID (required)")
	f.IntVar(&activityListFlags.activityTypeId, "activity-type-id", 0, "filter by activity type ID")

	activityListCmd.MarkFlagRequired("constituent-id")
}

func runActivityList(cmd *cobra.Command, args []string) error {
	client, err := loadClient()
	if err != nil {
		return err
	}

	activities, err := client.GetActivities(cmd.Context(), activityListFlags.constituentId, activityListFlags.activityTypeId)
	if err != nil {
		return fmt.Errorf("unable to fetch activities: %w", err)
	}

	out, err := json.MarshalIndent(activities, "", "  ")
	if err != nil {
		return fmt.Errorf("unable to format output: %w", err)
	}

	fmt.Fprintln(cmd.OutOrStdout(), string(out))
	return nil
}
