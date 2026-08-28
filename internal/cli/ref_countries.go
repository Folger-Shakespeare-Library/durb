package cli

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var refCountriesCmd = &cobra.Command{
	Use:   "countries",
	Short: "Countries",
	Long:  "Commands for country reference data.",
}

var refCountriesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List countries",
	Long:  "List available countries.",
	RunE:  runRefCountriesList,
}

func init() {
	refCountriesCmd.AddCommand(refCountriesListCmd)
}

func runRefCountriesList(cmd *cobra.Command, args []string) error {
	client, err := loadClient()
	if err != nil {
		return err
	}

	items, err := client.GetCountries(cmd.Context())
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
