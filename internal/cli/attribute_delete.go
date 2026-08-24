package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var attributeDeleteFlags struct {
	attributeId int
}

var attributeDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete an attribute record",
	Long: `Delete a specific attribute record by its ID.

Examples:
  tess attribute delete --attribute-id 98765`,
	RunE: runAttributeDelete,
}

func init() {
	attributeDeleteCmd.Flags().IntVar(&attributeDeleteFlags.attributeId, "attribute-id", 0, "attribute record ID (required)")
	attributeDeleteCmd.MarkFlagRequired("attribute-id")
}

func runAttributeDelete(cmd *cobra.Command, args []string) error {
	client, err := loadClient()
	if err != nil {
		return err
	}

	if err := client.DeleteAttribute(cmd.Context(), attributeDeleteFlags.attributeId); err != nil {
		return fmt.Errorf("unable to delete attribute: %w", err)
	}

	return nil
}
