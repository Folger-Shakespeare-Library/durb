package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var activityDeleteFlags struct {
	activityId int
}

var activityDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete an activity record",
	Long: `Delete a specific activity record by its ID.

Examples:
  tess activity delete --activity-id 98765`,
	RunE: runActivityDelete,
}

func init() {
	activityDeleteCmd.Flags().IntVar(&activityDeleteFlags.activityId, "activity-id", 0, "activity record ID (required)")
	activityDeleteCmd.MarkFlagRequired("activity-id")
}

func runActivityDelete(cmd *cobra.Command, args []string) error {
	client, err := loadClient()
	if err != nil {
		return err
	}

	if err := client.DeleteActivity(cmd.Context(), activityDeleteFlags.activityId); err != nil {
		return fmt.Errorf("unable to delete activity: %w", err)
	}

	return nil
}
