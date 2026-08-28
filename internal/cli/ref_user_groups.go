package cli

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var refUserGroupsCmd = &cobra.Command{
	Use:   "user-groups",
	Short: "User groups",
	Long:  "Commands for user group reference data.",
}

var refUserGroupsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List user groups",
	Long:  "List available user groups.",
	RunE:  runRefUserGroupsList,
}

func init() {
	refUserGroupsCmd.AddCommand(refUserGroupsListCmd)
}

func runRefUserGroupsList(cmd *cobra.Command, args []string) error {
	client, err := loadClient()
	if err != nil {
		return err
	}

	items, err := client.GetUserGroups(cmd.Context())
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
