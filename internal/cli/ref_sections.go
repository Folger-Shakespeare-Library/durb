package cli

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var refSectionsCmd = &cobra.Command{
	Use:   "sections",
	Short: "Sections",
	Long:  "Commands for venue section reference data.",
}

var refSectionsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List sections",
	Long:  "List available venue sections.",
	RunE:  runRefSectionsList,
}

func init() {
	refSectionsCmd.AddCommand(refSectionsListCmd)
}

func runRefSectionsList(cmd *cobra.Command, args []string) error {
	client, err := loadClient()
	if err != nil {
		return err
	}

	items, err := client.GetSections(cmd.Context())
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
