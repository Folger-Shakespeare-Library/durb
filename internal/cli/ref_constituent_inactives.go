package cli

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var refConstituentInactivesCmd = &cobra.Command{
	Use:   "constituent-inactives",
	Short: "Constituent inactive statuses",
	Long:  "Commands for constituent inactive status reference data.",
}

var refConstituentInactivesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List constituent inactive statuses",
	Long:  "List available inactive status types for constituents.",
	RunE:  runRefConstituentInactivesList,
}

func init() {
	refConstituentInactivesCmd.AddCommand(refConstituentInactivesListCmd)
}

func runRefConstituentInactivesList(cmd *cobra.Command, args []string) error {
	client, err := loadClient()
	if err != nil {
		return err
	}

	items, err := client.GetConstituentInactiveStatuses(cmd.Context())
	if err != nil {
		return err
	}

	type refItem struct {
		Id          int    `json:"id"`
		Description string `json:"description"`
		Inactive    bool   `json:"inactive"`
	}

	var out []refItem
	for _, item := range items {
		r := refItem{}
		if item.Id != nil {
			r.Id = *item.Id
		}
		if item.Description != nil {
			r.Description = *item.Description
		}
		if item.Inactive != nil {
			r.Inactive = *item.Inactive
		}
		out = append(out, r)
	}

	data, err := json.MarshalIndent(out, "", "  ")
	if err != nil {
		return fmt.Errorf("unable to format output: %w", err)
	}

	fmt.Fprintln(cmd.OutOrStdout(), string(data))
	return nil
}
