package cli

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Folger-Shakespeare-Library/durb/pkg/config"
	"github.com/Folger-Shakespeare-Library/durb/pkg/domain"
	"github.com/Folger-Shakespeare-Library/durb/pkg/tessitura"
	"github.com/spf13/cobra"
)

var includeFlags []string

var constituentGetCmd = &cobra.Command{
	Use:   "get <id> [id...]",
	Short: "Fetch one or more constituents by ID",
	Long: `Fetch one or more constituents by ID. Always returns a JSON array.

IDs can be passed as arguments, via stdin (one per line), or both.
Multiple IDs are fetched concurrently.

Examples:
  tess constituent get 12345
  tess constituent get 12345 67890
  echo "12345" | tess constituent get
  tess constituent search "Smith" | jq -r '.[].id' | tess constituent get`,
	Args: cobra.ArbitraryArgs,
	RunE: runConstituentGet,
}

func init() {
	constituentGetCmd.Flags().StringSliceVar(&includeFlags, "with", nil,
		`attach related data: affiliations, notes (or "all")`)
}

func runConstituentGet(cmd *cobra.Command, args []string) error {
	ids := append([]string{}, args...)

	// Read IDs from stdin if piped
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
		return fmt.Errorf("at least one constituent ID is required")
	}

	// Validate all IDs upfront
	for _, id := range ids {
		n, err := strconv.Atoi(id)
		if err != nil || n <= 0 {
			return fmt.Errorf("invalid constituent ID %q: must be a positive integer", id)
		}
	}

	cfg, err := config.Load()
	if err != nil {
		return err
	}
	if err := cfg.Validate(); err != nil {
		return err
	}

	client := tessitura.NewClient(cfg)
	includes := parseIncludes(includeFlags)
	includeAffiliations := includes["affiliations"]
	includeNotes := includes["notes"]

	// Fetch all constituents concurrently (each one gets its own HTTP call,
	// or batch call if extras are included)
	apiResults, err := client.GetConstituentsBatch(cmd.Context(), ids, includeAffiliations, includeNotes)
	if err != nil {
		return err
	}

	// Build domain objects in order
	var results []*domain.Constituent
	for _, r := range apiResults {
		constituent := domain.ConstituentFromAPI(r.Detail)

		if includeAffiliations {
			constituent.AttachAffiliations(r.Affiliations)
		}
		if includeNotes {
			constituent.AttachNotes(r.Notes)
		}

		results = append(results, constituent)
	}

	out, err := json.MarshalIndent(results, "", "  ")
	if err != nil {
		return fmt.Errorf("unable to format output: %w", err)
	}

	fmt.Fprintln(cmd.OutOrStdout(), string(out))
	return nil
}

// parseIncludes normalizes the --include flag values into a lookup map.
// Supports "all" to enable everything.
func parseIncludes(flags []string) map[string]bool {
	m := make(map[string]bool)
	for _, f := range flags {
		for _, part := range strings.Split(f, ",") {
			key := strings.TrimSpace(strings.ToLower(part))
			if key == "all" {
				m["affiliations"] = true
				m["notes"] = true
				return m
			}
			m[key] = true
		}
	}
	return m
}
