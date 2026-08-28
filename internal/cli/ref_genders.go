package cli

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var refGendersCmd = &cobra.Command{
	Use:   "genders",
	Short: "Genders",
	Long:  "Commands for gender reference data.",
}

var refGendersListCmd = &cobra.Command{
	Use:   "list",
	Short: "List genders",
	Long:  "List available genders.",
	RunE:  runRefGendersList,
}

func init() {
	refGendersCmd.AddCommand(refGendersListCmd)
}

func runRefGendersList(cmd *cobra.Command, args []string) error {
	client, err := loadClient()
	if err != nil {
		return err
	}

	items, err := client.GetGenders(cmd.Context())
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
