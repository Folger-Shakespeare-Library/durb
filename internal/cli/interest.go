package cli

import (
	"github.com/spf13/cobra"
)

var interestCmd = &cobra.Command{
	Use:   "interest",
	Short: "Manage constituent interests",
}

func init() {
	interestCmd.AddCommand(interestEnableCmd)
	interestCmd.AddCommand(interestDisableCmd)
	interestCmd.AddCommand(interestListCmd)
}
