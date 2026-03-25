package cli

import "github.com/spf13/cobra"

var reportCmd = &cobra.Command{
	Use:   "report",
	Short: "Work with reports",
	Long:  "Commands for managing Tessitura reports.",
}

func init() {
	reportCmd.AddCommand(reportGetCmd)
	reportCmd.AddCommand(reportListCmd)
	reportCmd.AddCommand(reportRequestCmd)
}
