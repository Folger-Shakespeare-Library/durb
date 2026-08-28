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
	refCmd.AddCommand(refAddressTypesCmd)
	refCmd.AddCommand(refAffiliationTypesCmd)
	refCmd.AddCommand(refAliasTypesCmd)
	refCmd.AddCommand(refAssociationTypesCmd)
	refCmd.AddCommand(refBusinessUnitsCmd)
	refCmd.AddCommand(refConstituencyTypesCmd)
	refCmd.AddCommand(refConstituentGroupsCmd)
	refCmd.AddCommand(refConstituentInactivesCmd)
	refCmd.AddCommand(refConstituentTypesCmd)
	refCmd.AddCommand(refContactPermissionCategoriesCmd)
	refCmd.AddCommand(refContactPermissionTypesCmd)
	refCmd.AddCommand(refCountriesCmd)
	refCmd.AddCommand(refDeliveryMethodsCmd)
	refCmd.AddCommand(refElectronicAddressTypesCmd)
	refCmd.AddCommand(refGendersCmd)
	refCmd.AddCommand(refInactiveReasonsCmd)
	refCmd.AddCommand(refInterestCategoriesCmd)
	refCmd.AddCommand(refInterestTypesCmd)
	refCmd.AddCommand(refKeywordCategoriesCmd)
	refCmd.AddCommand(refKeywordsCmd)
	refCmd.AddCommand(refLanguagesCmd)
	refCmd.AddCommand(refLoginTypesCmd)
	refCmd.AddCommand(refMachineSettingsCmd)
	refCmd.AddCommand(refNoteTypesCmd)
	refCmd.AddCommand(refOrderCategoriesCmd)
	refCmd.AddCommand(refOriginalSourcesCmd)
	refCmd.AddCommand(refPaymentTypesCmd)
	refCmd.AddCommand(refPerformanceStatusesCmd)
	refCmd.AddCommand(refPerformanceTypesCmd)
	refCmd.AddCommand(refPrefixesCmd)
	refCmd.AddCommand(refPriceCategoriesCmd)
	refCmd.AddCommand(refPriceTypeCategoriesCmd)
	refCmd.AddCommand(refPriceTypeGroupsCmd)
	refCmd.AddCommand(refPronounsCmd)
	refCmd.AddCommand(refReportCategoriesCmd)
	refCmd.AddCommand(refReportTypesCmd)
	refCmd.AddCommand(refSalesChannelsCmd)
	refCmd.AddCommand(refSeatCodesCmd)
	refCmd.AddCommand(refSeatStatusesCmd)
	refCmd.AddCommand(refSeasonsCmd)
	refCmd.AddCommand(refSectionsCmd)
	refCmd.AddCommand(refStatesCmd)
	refCmd.AddCommand(refSuffixesCmd)
	refCmd.AddCommand(refTheatersCmd)
	refCmd.AddCommand(refUserGroupsCmd)
}
