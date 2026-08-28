package cli

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var refPhoneIndicatorsCmd = &cobra.Command{
	Use:   "phone-indicators",
	Short: "Phone indicators",
	Long:  "Commands for phone indicator reference data.",
}

var refPhoneIndicatorsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List phone indicators",
	Long:  "List available phone indicator preferences.",
	RunE:  runRefPhoneIndicatorsList,
}

func init() {
	refPhoneIndicatorsCmd.AddCommand(refPhoneIndicatorsListCmd)
}

func runRefPhoneIndicatorsList(cmd *cobra.Command, args []string) error {
	client, err := loadClient()
	if err != nil {
		return err
	}

	items, err := client.GetPhoneIndicators(cmd.Context())
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
