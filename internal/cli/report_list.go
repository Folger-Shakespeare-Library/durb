package cli

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/Folger-Shakespeare-Library/durb/pkg/config"
	"github.com/Folger-Shakespeare-Library/durb/pkg/domain"
	"github.com/Folger-Shakespeare-Library/durb/pkg/tessitura"
	"github.com/spf13/cobra"
)

var (
	reportListTypeIds      []string
	reportListCategoryIds  []string
	reportListIncludeInactive bool
)

var reportListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all reports",
	Long: `List all reports. Optionally filter by type or category ID.

Examples:
  tess report list
  tess report list --type-ids 6
  tess report list --category-ids 9
  tess report list --type-ids 6 --category-ids 9,12
  tess report list --include-inactive`,
	Args: cobra.NoArgs,
	RunE: runReportList,
}

func init() {
	reportListCmd.Flags().StringSliceVar(&reportListTypeIds, "type-ids", nil, "filter by report type ID (comma-delimited)")
	reportListCmd.Flags().StringSliceVar(&reportListCategoryIds, "category-ids", nil, "filter by category ID (comma-delimited)")
	reportListCmd.Flags().BoolVar(&reportListIncludeInactive, "include-inactive", false, "include inactive reports")
}

func runReportList(cmd *cobra.Command, args []string) error {
	cfg, err := config.Load()
	if err != nil {
		return err
	}
	if err := cfg.Validate(); err != nil {
		return err
	}

	client := tessitura.NewClient(cfg)

	apiResults, err := client.GetReports(cmd.Context(), strings.Join(reportListTypeIds, ","), strings.Join(reportListCategoryIds, ","))
	if err != nil {
		return err
	}

	var results []*domain.Report
	for _, r := range apiResults {
		report := domain.ReportFromAPI(r)
		if !reportListIncludeInactive && report.Inactive {
			continue
		}
		results = append(results, report)
	}

	out, err := json.MarshalIndent(results, "", "  ")
	if err != nil {
		return fmt.Errorf("unable to format output: %w", err)
	}

	fmt.Fprintln(cmd.OutOrStdout(), string(out))
	return nil
}
