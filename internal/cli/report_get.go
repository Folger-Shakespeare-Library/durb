package cli

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/Folger-Shakespeare-Library/durb/pkg/config"
	"github.com/Folger-Shakespeare-Library/durb/pkg/domain"
	"github.com/Folger-Shakespeare-Library/durb/pkg/tessitura"
	"github.com/spf13/cobra"
)

var reportGetCmd = &cobra.Command{
	Use:   "get <id> [id...]",
	Short: "Fetch one or more reports by ID",
	Long: `Fetch one or more reports by ID. Always returns a JSON array.

IDs can be passed as arguments, via stdin (one per line), or both.
Multiple IDs are fetched concurrently.

Examples:
  tess report get 42
  tess report get 42 43
  echo "42" | tess report get`,
	Args: cobra.ArbitraryArgs,
	RunE: runReportGet,
}

func runReportGet(cmd *cobra.Command, args []string) error {
	ids := append([]string{}, args...)

	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			line := strings.TrimSpace(scanner.Text())
			if line != "" {
				ids = append(ids, line)
			}
		}
	}

	if len(ids) == 0 {
		return fmt.Errorf("at least one report ID is required")
	}

	cfg, err := config.Load()
	if err != nil {
		return err
	}
	if err := cfg.Validate(); err != nil {
		return err
	}

	client := tessitura.NewClient(cfg)

	apiResults, err := client.GetReportsBatch(cmd.Context(), ids)
	if err != nil {
		return err
	}

	results := make([]*domain.Report, len(apiResults))
	for i, r := range apiResults {
		report := domain.ReportFromAPI(r.Base)
		report.AttachDetail(r.Detail)
		results[i] = report
	}

	out, err := json.MarshalIndent(results, "", "  ")
	if err != nil {
		return fmt.Errorf("unable to format output: %w", err)
	}

	fmt.Fprintln(cmd.OutOrStdout(), string(out))
	return nil
}
