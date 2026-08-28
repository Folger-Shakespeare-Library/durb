package cli

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var refDesignationCodesCmd = &cobra.Command{
	Use:   "designation-codes",
	Short: "Designation codes",
	Long:  "Commands for designation code reference data.",
}

var refDesignationCodesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List designation codes",
	Long:  "List available designation codes.",
	RunE:  runRefDesignationCodesList,
}

func init() {
	refDesignationCodesCmd.AddCommand(refDesignationCodesListCmd)
}

func runRefDesignationCodesList(cmd *cobra.Command, args []string) error {
	client, err := loadClient()
	if err != nil {
		return err
	}

	items, err := client.GetDesignationCodes(cmd.Context())
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
