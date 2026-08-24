package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var attributeSetFlags struct {
	constituentId   int
	attributeTypeId int
	value           string
}

var attributeSetCmd = &cobra.Command{
	Use:   "set",
	Short: "Set an attribute value on a constituent",
	Long: `Set an attribute on a constituent by attribute type ID.

Creates the attribute if no record exists for that keyword on the
constituent. Updates the value if one already exists.

Examples:
  tess attribute set --constituent-id 446106 --attribute-type-id 651 --value 3`,
	RunE: runAttributeSet,
}

func init() {
	f := attributeSetCmd.Flags()
	f.IntVar(&attributeSetFlags.constituentId, "constituent-id", 0, "constituent ID (required)")
	f.IntVar(&attributeSetFlags.attributeTypeId, "attribute-type-id", 0, "attribute type / keyword ID (required)")
	f.StringVar(&attributeSetFlags.value, "value", "", "attribute value (required)")

	attributeSetCmd.MarkFlagRequired("constituent-id")
	attributeSetCmd.MarkFlagRequired("attribute-type-id")
	attributeSetCmd.MarkFlagRequired("value")
}

func runAttributeSet(cmd *cobra.Command, args []string) error {
	client, err := loadClient()
	if err != nil {
		return err
	}

	existing, err := client.GetAttributes(cmd.Context(), attributeSetFlags.constituentId, attributeSetFlags.attributeTypeId)
	if err != nil {
		return fmt.Errorf("unable to fetch attributes: %w", err)
	}

	if len(existing) > 0 {
		attr := existing[0]
		body := map[string]interface{}{
			"Id":              attr.Id,
			"Constituent":    map[string]int{"Id": attributeSetFlags.constituentId},
			"Keyword":        map[string]int{"Id": attributeSetFlags.attributeTypeId},
			"Value":          attributeSetFlags.value,
			"UpdatedDateTime": attr.UpdatedDateTime,
		}
		return client.UpdateAttribute(cmd.Context(), *attr.Id, body)
	}

	body := map[string]interface{}{
		"Constituent": map[string]int{"Id": attributeSetFlags.constituentId},
		"Keyword":     map[string]int{"Id": attributeSetFlags.attributeTypeId},
		"Value":       attributeSetFlags.value,
	}
	return client.CreateAttribute(cmd.Context(), body)
}
