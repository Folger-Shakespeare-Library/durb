package cli

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var refMailIndicatorsCmd = &cobra.Command{
	Use:   "mail-indicators",
	Short: "Mail indicators",
	Long:  "Commands for mail indicator reference data.",
}

var refMailIndicatorsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List mail indicators",
	Long:  "List available mail indicator preferences.",
	RunE:  runRefMailIndicatorsList,
}

func init() {
	refMailIndicatorsCmd.AddCommand(refMailIndicatorsListCmd)
}

func runRefMailIndicatorsList(cmd *cobra.Command, args []string) error {
	client, err := loadClient()
	if err != nil {
		return err
	}

	items, err := client.GetMailIndicators(cmd.Context())
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
