package cli

import (
	"github.com/spf13/cobra"
)

var crmCmd = &cobra.Command{
	Use:   "crm",
	Short: "CRM commands",
	Long:  "Commands for Tessitura CRM resources (constituents, activities, interests, attributes).",
}

func init() {
	crmCmd.AddCommand(activityCmd)
	crmCmd.AddCommand(attributeCmd)
	crmCmd.AddCommand(constituentCmd)
	crmCmd.AddCommand(interestCmd)
}
