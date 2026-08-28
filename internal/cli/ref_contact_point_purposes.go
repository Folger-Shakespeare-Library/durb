package cli

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var refContactPointPurposesCmd = &cobra.Command{
	Use:   "contact-point-purposes",
	Short: "Contact point purposes",
	Long:  "Commands for contact point purpose reference data.",
}

var refContactPointPurposesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List contact point purposes",
	Long:  "List available contact point purposes.",
	RunE:  runRefContactPointPurposesList,
}

func init() {
	refContactPointPurposesCmd.AddCommand(refContactPointPurposesListCmd)
}

func runRefContactPointPurposesList(cmd *cobra.Command, args []string) error {
	client, err := loadClient()
	if err != nil {
		return err
	}

	items, err := client.GetContactPointPurposes(cmd.Context())
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
