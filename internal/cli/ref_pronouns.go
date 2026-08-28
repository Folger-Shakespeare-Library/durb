package cli

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var refPronounsCmd = &cobra.Command{
	Use:   "pronouns",
	Short: "Pronouns",
	Long:  "Commands for pronoun reference data.",
}

var refPronounsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List pronouns",
	Long:  "List available pronouns.",
	RunE:  runRefPronounsList,
}

func init() {
	refPronounsCmd.AddCommand(refPronounsListCmd)
}

func runRefPronounsList(cmd *cobra.Command, args []string) error {
	client, err := loadClient()
	if err != nil {
		return err
	}

	items, err := client.GetPronouns(cmd.Context())
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
