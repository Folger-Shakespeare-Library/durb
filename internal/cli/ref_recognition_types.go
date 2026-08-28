package cli

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var refRecognitionTypesCmd = &cobra.Command{
	Use:   "recognition-types",
	Short: "Recognition types",
	Long:  "Commands for recognition type reference data.",
}

var refRecognitionTypesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List recognition types",
	Long:  "List available recognition types.",
	RunE:  runRefRecognitionTypesList,
}

func init() {
	refRecognitionTypesCmd.AddCommand(refRecognitionTypesListCmd)
}

func runRefRecognitionTypesList(cmd *cobra.Command, args []string) error {
	client, err := loadClient()
	if err != nil {
		return err
	}

	items, err := client.GetRecognitionTypes(cmd.Context())
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
