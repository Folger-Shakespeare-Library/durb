package cli

import "github.com/spf13/cobra"

var reportRequestCmd = &cobra.Command{
	Use:   "request",
	Short: "Work with report requests",
	Long:  "Commands for managing Tessitura report requests.",
}

func init() {
	reportRequestCmd.AddCommand(reportRequestGetCmd)
	reportRequestCmd.AddCommand(reportRequestListCmd)
	reportRequestCmd.AddCommand(reportRequestResultsCmd)
}
