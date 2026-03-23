package cli

import (
	"github.com/spf13/cobra"
)

var constituentCmd = &cobra.Command{
	Use:     "constituent",
	Aliases: []string{"con"},
	Short:   "Work with constituents",
	Long:    "Commands for managing Tessitura constituent records.",
}

func init() {
	constituentCmd.AddCommand(constituentGetCmd)
	constituentCmd.AddCommand(constituentSearchCmd)
}
