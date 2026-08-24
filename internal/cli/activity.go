package cli

import (
	"github.com/spf13/cobra"
)

var activityCmd = &cobra.Command{
	Use:   "activity",
	Short: "Manage constituent activities",
}

func init() {
	activityCmd.AddCommand(activityCreateCmd)
	activityCmd.AddCommand(activityDeleteCmd)
	activityCmd.AddCommand(activityListCmd)
}
