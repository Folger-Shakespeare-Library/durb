package cli

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var refWebContentTypesCmd = &cobra.Command{
	Use:   "web-content-types",
	Short: "Web content types",
	Long:  "Commands for web content type reference data.",
}

var refWebContentTypesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List web content types",
	Long:  "List available web content types.",
	RunE:  runRefWebContentTypesList,
}

func init() {
	refWebContentTypesCmd.AddCommand(refWebContentTypesListCmd)
}

func runRefWebContentTypesList(cmd *cobra.Command, args []string) error {
	client, err := loadClient()
	if err != nil {
		return err
	}

	items, err := client.GetWebContentTypes(cmd.Context())
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
