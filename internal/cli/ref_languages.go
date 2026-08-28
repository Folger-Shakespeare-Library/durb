package cli

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var refLanguagesCmd = &cobra.Command{
	Use:   "languages",
	Short: "Languages",
	Long:  "Commands for language reference data.",
}

var refLanguagesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List languages",
	Long:  "List available languages.",
	RunE:  runRefLanguagesList,
}

func init() {
	refLanguagesCmd.AddCommand(refLanguagesListCmd)
}

func runRefLanguagesList(cmd *cobra.Command, args []string) error {
	client, err := loadClient()
	if err != nil {
		return err
	}

	items, err := client.GetLanguages(cmd.Context())
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
