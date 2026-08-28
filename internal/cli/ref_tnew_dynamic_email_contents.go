package cli

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var refTNEWDynamicEmailContentsCmd = &cobra.Command{
	Use:   "tnew-dynamic-email-contents",
	Short: "TNEW dynamic email contents",
	Long:  "Commands for TNEW dynamic email content reference data.",
}

var refTNEWDynamicEmailContentsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List TNEW dynamic email contents",
	Long:  "List available TNEW dynamic email contents.",
	RunE:  runRefTNEWDynamicEmailContentsList,
}

func init() {
	refTNEWDynamicEmailContentsCmd.AddCommand(refTNEWDynamicEmailContentsListCmd)
}

func runRefTNEWDynamicEmailContentsList(cmd *cobra.Command, args []string) error {
	client, err := loadClient()
	if err != nil {
		return err
	}

	items, err := client.GetTNEWDynamicEmailContents(cmd.Context())
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
