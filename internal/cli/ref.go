package cli

import (
	"github.com/spf13/cobra"
)

var refCmd = &cobra.Command{
	Use:   "ref",
	Short: "Reference data lookups",
	Long:  "Commands for listing Tessitura reference data (lookup tables).",
}

func init() {
	refCmd.AddCommand(refConstituentInactivesCmd)
	refCmd.AddCommand(refSeatStatusesCmd)
}
