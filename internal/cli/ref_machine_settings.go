package cli

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var refMachineSettingsCmd = &cobra.Command{
	Use:   "machine-settings",
	Short: "Machine settings",
	Long:  "Commands for machine settings reference data.",
}

var refMachineSettingsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List machine settings",
	Long:  "List available machine settings.",
	RunE:  runRefMachineSettingsList,
}

func init() {
	refMachineSettingsCmd.AddCommand(refMachineSettingsListCmd)
}

func runRefMachineSettingsList(cmd *cobra.Command, args []string) error {
	client, err := loadClient()
	if err != nil {
		return err
	}

	items, err := client.GetMachineSettings(cmd.Context())
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
