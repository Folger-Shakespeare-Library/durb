package cli

import (
	"github.com/spf13/cobra"
)

var interestDisableFlags struct {
	constituentId int
	interestIds   string
}

var interestDisableCmd = &cobra.Command{
	Use:   "disable",
	Short: "Disable interests on a constituent",
	Long: `Disable one or more interests on a constituent by interest type ID.

Examples:
  tess interest disable --constituent-id 446106 --interest-type-ids 262
  tess interest disable --constituent-id 446106 --interest-type-ids 262,263,264`,
	RunE: runInterestDisable,
}

func init() {
	f := interestDisableCmd.Flags()
	f.IntVar(&interestDisableFlags.constituentId, "constituent-id", 0, "constituent ID (required)")
	f.StringVar(&interestDisableFlags.interestIds, "interest-type-ids", "", "comma-separated interest type IDs (required)")

	interestDisableCmd.MarkFlagRequired("constituent-id")
	interestDisableCmd.MarkFlagRequired("interest-type-ids")
}

func runInterestDisable(cmd *cobra.Command, args []string) error {
	return setInterests(cmd, interestDisableFlags.constituentId, interestDisableFlags.interestIds, false)
}
