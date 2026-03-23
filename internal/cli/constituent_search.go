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
	searchLastName           string
	searchFirstName          string
	searchStreet             string
	searchPostalCode         string
	searchID                 string
	searchEmail              string
	searchPhone              string
	searchOrderNo            string
	searchWebLogin           string
	searchCustomerServiceNo  string
	searchGroups             string
	searchIncludeAffiliates  bool
)

var constituentSearchCmd = &cobra.Command{
	Use:   "search [query]",
	Short: "Search for constituents",
	Long: `Search for constituents by free text, structured fields, or advanced criteria.

Free-text search:
  tess constituent search Smith
  tess constituent search "Jane Smith"

Basic search (combinable):
  tess constituent search --last-name Smith
  tess constituent search --first-name Jane --last-name Smith
  tess constituent search --street "Main Street"
  tess constituent search --postal-code 10001
  tess constituent search --id 12345

Advanced search (one at a time):
  tess constituent search --email user@example.com
  tess constituent search --phone 5551234567
  tess constituent search --order-no 99999
  tess constituent search --web-login jsmith
  tess constituent search --customer-service-no 11111

Filter by constituent group (works with any search mode):
  tess constituent search Smith --groups individuals
  tess constituent search --last-name Smith --groups individuals,households`,
	RunE: runConstituentSearch,
}

func init() {
	// Basic search flags.
	constituentSearchCmd.Flags().StringVar(&searchLastName, "last-name", "", "last name")
	constituentSearchCmd.Flags().StringVar(&searchFirstName, "first-name", "", "first name")
	constituentSearchCmd.Flags().StringVar(&searchStreet, "street", "", "street address")
	constituentSearchCmd.Flags().StringVar(&searchPostalCode, "postal-code", "", "postal/ZIP code")
	constituentSearchCmd.Flags().StringVar(&searchID, "id", "", "constituent ID")

	// Advanced search flags.
	constituentSearchCmd.Flags().StringVar(&searchEmail, "email", "", "email address")
	constituentSearchCmd.Flags().StringVar(&searchPhone, "phone", "", "phone number")
	constituentSearchCmd.Flags().StringVar(&searchOrderNo, "order-no", "", "order number")
	constituentSearchCmd.Flags().StringVar(&searchWebLogin, "web-login", "", "web login username")
	constituentSearchCmd.Flags().StringVar(&searchCustomerServiceNo, "customer-service-no", "", "customer service number")

	// Filter flags.
	constituentSearchCmd.Flags().StringVar(&searchGroups, "groups", "", "filter by constituent groups: individuals, organizations, households (comma-separated)")
	constituentSearchCmd.Flags().BoolVar(&searchIncludeAffiliates, "include-affiliations", false, "include affiliated constituents in results")
}

func runConstituentSearch(cmd *cobra.Command, args []string) error {
	params := tessitura.SearchParams{
		LastName:          searchLastName,
		FirstName:         searchFirstName,
		Street:            searchStreet,
		PostalCode:        searchPostalCode,
		ID:                searchID,
		Email:             searchEmail,
		Phone:             searchPhone,
		OrderNo:           searchOrderNo,
		WebLogin:          searchWebLogin,
		CustomerServiceNo: searchCustomerServiceNo,
		IncludeAffiliates: searchIncludeAffiliates,
		ConstituentGroups: searchGroups,
	}

	if len(args) > 0 {
		params.Query = strings.Join(args, " ")
	}

	hasQuery := params.Query != ""
	hasBasic := params.LastName != "" || params.FirstName != "" || params.Street != "" ||
		params.PostalCode != "" || params.ID != ""
	hasAdvanced := params.IsAdvanced()

	if !hasQuery && !hasBasic && !hasAdvanced {
		return fmt.Errorf("provide a search query or use flags (see --help for options)")
	}

	// Count how many search modes are active.
	modes := 0
	if hasQuery {
		modes++
	}
	if hasBasic {
		modes++
	}
	if hasAdvanced {
		modes++
	}
	if modes > 1 {
		return fmt.Errorf("cannot combine free-text, basic (--last-name, etc.), and advanced (--email, --phone, etc.) searches — these use different search modes")
	}

	cfg, err := config.Load()
	if err != nil {
		return err
	}
	if err := cfg.Validate(); err != nil {
		return err
	}

	client := tessitura.NewClient(cfg)

	resp, err := client.SearchConstituents(cmd.Context(), params)
	if err != nil {
		return err
	}

	results := domain.SearchResultsFromAPI(resp)

	out, err := json.MarshalIndent(results, "", "  ")
	if err != nil {
		return fmt.Errorf("unable to format output: %w", err)
	}

	fmt.Fprintln(cmd.OutOrStdout(), string(out))
	return nil
}
