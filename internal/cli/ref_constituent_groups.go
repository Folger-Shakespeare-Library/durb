package cli

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var refConstituentGroupsCmd = &cobra.Command{
	Use:   "constituent-groups",
	Short: "Constituent groups",
	Long:  "Commands for constituent group reference data.",
}

var refConstituentGroupsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List constituent groups",
	Long:  "List available constituent groups.",
	RunE:  runRefConstituentGroupsList,
}

func init() {
	refConstituentGroupsCmd.AddCommand(refConstituentGroupsListCmd)
}

func runRefConstituentGroupsList(cmd *cobra.Command, args []string) error {
	client, err := loadClient()
	if err != nil {
		return err
	}

	items, err := client.GetConstituentGroups(cmd.Context())
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
