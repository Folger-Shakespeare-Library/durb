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
	refCmd.AddCommand(refAffiliationTypesCmd)
	refCmd.AddCommand(refAliasTypesCmd)
	refCmd.AddCommand(refAssociationTypesCmd)
	refCmd.AddCommand(refConstituentInactivesCmd)
	refCmd.AddCommand(refConstituentTypesCmd)
	refCmd.AddCommand(refContactPermissionCategoriesCmd)
	refCmd.AddCommand(refContactPermissionTypesCmd)
	refCmd.AddCommand(refDeliveryMethodsCmd)
	refCmd.AddCommand(refElectronicAddressTypesCmd)
	refCmd.AddCommand(refInactiveReasonsCmd)
	refCmd.AddCommand(refInterestTypesCmd)
	refCmd.AddCommand(refKeywordsCmd)
	refCmd.AddCommand(refLoginTypesCmd)
	refCmd.AddCommand(refMachineSettingsCmd)
	refCmd.AddCommand(refNoteTypesCmd)
	refCmd.AddCommand(refOrderCategoriesCmd)
	refCmd.AddCommand(refOriginalSourcesCmd)
	refCmd.AddCommand(refPaymentTypesCmd)
	refCmd.AddCommand(refPerformanceStatusesCmd)
	refCmd.AddCommand(refPerformanceTypesCmd)
	refCmd.AddCommand(refPriceCategoriesCmd)
	refCmd.AddCommand(refPriceTypeCategoriesCmd)
	refCmd.AddCommand(refPriceTypeGroupsCmd)
	refCmd.AddCommand(refReportCategoriesCmd)
	refCmd.AddCommand(refReportTypesCmd)
	refCmd.AddCommand(refSalesChannelsCmd)
	refCmd.AddCommand(refSeasonsCmd)
	refCmd.AddCommand(refSeatStatusesCmd)
}
