package cli

import (
	"github.com/spf13/cobra"
)

var attributeCmd = &cobra.Command{
	Use:   "attribute",
	Short: "Manage constituent attributes",
}

func init() {
	attributeCmd.AddCommand(attributeDeleteCmd)
	attributeCmd.AddCommand(attributeListCmd)
	attributeCmd.AddCommand(attributeSetCmd)
}
