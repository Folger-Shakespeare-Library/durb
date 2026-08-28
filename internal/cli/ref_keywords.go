package cli

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var refKeywordsCmd = &cobra.Command{
	Use:   "keywords",
	Short: "Keywords (attribute types)",
	Long:  "Commands for keyword reference data. Keywords are attribute types in Tessitura.",
}

var refKeywordsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List keywords",
	Long:  "List available keywords (attribute types).",
	RunE:  runRefKeywordsList,
}

func init() {
	refKeywordsCmd.AddCommand(refKeywordsListCmd)
}

func runRefKeywordsList(cmd *cobra.Command, args []string) error {
	client, err := loadClient()
	if err != nil {
		return err
	}

	items, err := client.GetKeywords(cmd.Context())
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
