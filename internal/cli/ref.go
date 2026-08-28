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
	refCmd.AddCommand(refActivityStatusesCmd)
	refCmd.AddCommand(refActivityTypesCmd)
	refCmd.AddCommand(refConstituentInactivesCmd)
	refCmd.AddCommand(refConstituentTypesCmd)
	refCmd.AddCommand(refInactiveReasonsCmd)
	refCmd.AddCommand(refInterestTypesCmd)
	refCmd.AddCommand(refKeywordsCmd)
	refCmd.AddCommand(refMachineSettingsCmd)
	refCmd.AddCommand(refOriginalSourcesCmd)
	refCmd.AddCommand(refSeatStatusesCmd)
}
