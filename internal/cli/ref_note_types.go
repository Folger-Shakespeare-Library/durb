package cli

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var refNoteTypesCmd = &cobra.Command{
	Use:   "note-types",
	Short: "Note types",
	Long:  "Commands for note type reference data.",
}

var refNoteTypesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List note types",
	Long:  "List available note types.",
	RunE:  runRefNoteTypesList,
}

func init() {
	refNoteTypesCmd.AddCommand(refNoteTypesListCmd)
}

func runRefNoteTypesList(cmd *cobra.Command, args []string) error {
	client, err := loadClient()
	if err != nil {
		return err
	}

	items, err := client.GetNoteTypes(cmd.Context())
	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(items, "", "  ")
	if err != nil {
		return fmt.Errorf("unable to format output: %w", err)
	}

	fmt.Fprintln(cmd.OutOrStdout(), string(data))
	return nil
}
