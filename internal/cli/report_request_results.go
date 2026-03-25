package cli

import (
	"encoding/json"
	"fmt"

	"github.com/Folger-Shakespeare-Library/durb/pkg/config"
	"github.com/Folger-Shakespeare-Library/durb/pkg/domain"
	"github.com/Folger-Shakespeare-Library/durb/pkg/tessitura"
	"github.com/spf13/cobra"
)

var reportRequestResultsFlags struct {
	scheduleName      string
	reportId          string
	startDate         string
	endDate           string
	includePublic     bool
	myReportsOnly     bool
	recentOnly        bool
	includeErrors     bool
	includeDeleted    bool
	page              int
	pageSize          int
}

var reportRequestResultsCmd = &cobra.Command{
	Use:   "results",
	Short: "List scheduled report results",
	Long: `List scheduled report results. Returns a combined view of report request,
schedule, and report definition data.

Examples:
  tess report request results
  tess report request results --report-id perfseatingbook
  tess report request results --start-date 2025-01-01 --end-date 2025-12-31
  tess report request results --schedule-name "Daily Seating"
  tess report request results --my-reports-only
  tess report request results --recent-only
  tess report request results --page 2 --page-size 50`,
	Args: cobra.NoArgs,
	RunE: runReportRequestResults,
}

func init() {
	f := reportRequestResultsCmd.Flags()
	f.StringVar(&reportRequestResultsFlags.scheduleName, "schedule-name", "", "filter by schedule name")
	f.StringVar(&reportRequestResultsFlags.reportId, "report-id", "", "filter by report ID")
	f.StringVar(&reportRequestResultsFlags.startDate, "start-date", "", "filter by start date (YYYY-MM-DD)")
	f.StringVar(&reportRequestResultsFlags.endDate, "end-date", "", "filter by end date (YYYY-MM-DD)")
	f.BoolVar(&reportRequestResultsFlags.includePublic, "include-public", false, "include public report results")
	f.BoolVar(&reportRequestResultsFlags.myReportsOnly, "my-reports-only", false, "only return results owned by the current user")
	f.BoolVar(&reportRequestResultsFlags.recentOnly, "recent-only", false, "only return recent results")
	f.BoolVar(&reportRequestResultsFlags.includeErrors, "include-errors", false, "include errored results")
	f.BoolVar(&reportRequestResultsFlags.includeDeleted, "include-deleted", false, "include results whose output has been deleted")
	f.IntVar(&reportRequestResultsFlags.page, "page", 1, "page number")
	f.IntVar(&reportRequestResultsFlags.pageSize, "page-size", 100, "results per page")
}

func runReportRequestResults(cmd *cobra.Command, args []string) error {
	cfg, err := config.Load()
	if err != nil {
		return err
	}
	if err := cfg.Validate(); err != nil {
		return err
	}

	client := tessitura.NewClient(cfg)

	params := tessitura.ReportResultsParams{
		ScheduleName:      reportRequestResultsFlags.scheduleName,
		ReportId:          reportRequestResultsFlags.reportId,
		StartDate:         reportRequestResultsFlags.startDate,
		EndDate:           reportRequestResultsFlags.endDate,
		IncludePublic:     reportRequestResultsFlags.includePublic,
		MyReportsOnly:     reportRequestResultsFlags.myReportsOnly,
		RecentResultsOnly: reportRequestResultsFlags.recentOnly,
		IncludeErrors:     reportRequestResultsFlags.includeErrors,
		IncludeDeleted:    reportRequestResultsFlags.includeDeleted,
		Page:              reportRequestResultsFlags.page,
		PageSize:          reportRequestResultsFlags.pageSize,
	}

	apiResults, total, err := client.GetReportResults(cmd.Context(), params)
	if err != nil {
		return err
	}

	results := make([]*domain.ReportResult, 0, len(apiResults))
	for _, r := range apiResults {
		results = append(results, domain.ReportResultFromAPI(r))
	}

	out, err := json.MarshalIndent(results, "", "  ")
	if err != nil {
		return fmt.Errorf("unable to format output: %w", err)
	}

	fmt.Fprintln(cmd.OutOrStdout(), string(out))

	// Print pagination info to stderr if there are more pages
	if total > reportRequestResultsFlags.page*reportRequestResultsFlags.pageSize {
		remaining := total - reportRequestResultsFlags.page*reportRequestResultsFlags.pageSize
		fmt.Fprintf(cmd.ErrOrStderr(), "%d more result(s) available (use --page to paginate)\n", remaining)
	}

	return nil
}
