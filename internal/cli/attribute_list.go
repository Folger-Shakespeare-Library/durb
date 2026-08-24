package cli

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var attributeListFlags struct {
	constituentId   int
	attributeTypeId int
}

var attributeListCmd = &cobra.Command{
	Use:   "list",
	Short: "List attributes on a constituent",
	Long: `List attributes on a constituent, optionally filtered by attribute type ID.

Examples:
  tess attribute list --constituent-id 446106
  tess attribute list --constituent-id 446106 --attribute-type-id 651`,
	RunE: runAttributeList,
}

func init() {
	f := attributeListCmd.Flags()
	f.IntVar(&attributeListFlags.constituentId, "constituent-id", 0, "constituent ID (required)")
	f.IntVar(&attributeListFlags.attributeTypeId, "attribute-type-id", 0, "filter by attribute type / keyword ID")

	attributeListCmd.MarkFlagRequired("constituent-id")
}

func runAttributeList(cmd *cobra.Command, args []string) error {
	client, err := loadClient()
	if err != nil {
		return err
	}

	attrs, err := client.GetAttributes(cmd.Context(), attributeListFlags.constituentId, attributeListFlags.attributeTypeId)
	if err != nil {
		return fmt.Errorf("unable to fetch attributes: %w", err)
	}

	out, err := json.MarshalIndent(attrs, "", "  ")
	if err != nil {
		return fmt.Errorf("unable to format output: %w", err)
	}

	fmt.Fprintln(cmd.OutOrStdout(), string(out))
	return nil
}
