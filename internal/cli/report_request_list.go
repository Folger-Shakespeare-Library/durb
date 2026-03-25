package cli

import (
	"encoding/json"
	"fmt"

	"github.com/Folger-Shakespeare-Library/durb/pkg/config"
	"github.com/Folger-Shakespeare-Library/durb/pkg/domain"
	"github.com/Folger-Shakespeare-Library/durb/pkg/tessitura"
	"github.com/spf13/cobra"
)

var reportRequestListIncludeInactive bool

var reportRequestListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all report requests",
	Long: `List all report requests. Active requests only by default.

Examples:
  tess report request list
  tess report request list --include-inactive`,
	Args: cobra.NoArgs,
	RunE: runReportRequestList,
}

func init() {
	reportRequestListCmd.Flags().BoolVar(&reportRequestListIncludeInactive, "include-inactive", false, "include completed and cancelled requests")
}

func runReportRequestList(cmd *cobra.Command, args []string) error {
	cfg, err := config.Load()
	if err != nil {
		return err
	}
	if err := cfg.Validate(); err != nil {
		return err
	}

	client := tessitura.NewClient(cfg)

	activeOnly := !reportRequestListIncludeInactive
	apiResults, err := client.GetReportRequests(cmd.Context(), activeOnly)
	if err != nil {
		return err
	}

	results := make([]*domain.ReportRequest, 0, len(apiResults))
	for _, r := range apiResults {
		results = append(results, domain.ReportRequestFromAPI(r))
	}

	out, err := json.MarshalIndent(results, "", "  ")
	if err != nil {
		return fmt.Errorf("unable to format output: %w", err)
	}

	fmt.Fprintln(cmd.OutOrStdout(), string(out))
	return nil
}
