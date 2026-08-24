package cli

import (
	"encoding/json"
	"fmt"

	"github.com/Folger-Shakespeare-Library/durb/pkg/domain"
	"github.com/Folger-Shakespeare-Library/durb/pkg/tessitura"
	"github.com/spf13/cobra"
)

var createFlags struct {
	firstName         string
	lastName          string
	email             string
	constituentTypeId int
	originalSourceId  int
	postalCode        string
	street            string
}

var constituentCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new constituent",
	Long: `Create a new constituent in Tessitura.

Requires --original-source-id and --constituent-type-id (no defaults for write operations).
Auto-generates SortName as "LastName, FirstName".

Examples:
  tess constituent create --first "Jane" --last "Smith" --email "jane@example.com" --constituent-type-id 1 --original-source-id 33
  tess constituent create --first "Jane" --last "Smith" --email "jane@example.com" --constituent-type-id 1 --original-source-id 33 --street "N/A" --postal-code 20003`,
	RunE: runConstituentCreate,
}

func init() {
	f := constituentCreateCmd.Flags()
	f.StringVar(&createFlags.firstName, "first", "", "first name (required, max 20 chars)")
	f.StringVar(&createFlags.lastName, "last", "", "last name (required, max 55 chars)")
	f.StringVar(&createFlags.email, "email", "", "email address (required)")
	f.IntVar(&createFlags.constituentTypeId, "constituent-type-id", 0, "constituent type ID (required, e.g. 1 = Individual)")
	f.IntVar(&createFlags.originalSourceId, "original-source-id", 0, "original source ID (required)")
	f.StringVar(&createFlags.street, "street", "", "street address line 1")
	f.StringVar(&createFlags.postalCode, "postal-code", "", "postal/ZIP code (max 10 chars)")

	constituentCreateCmd.MarkFlagRequired("first")
	constituentCreateCmd.MarkFlagRequired("last")
	constituentCreateCmd.MarkFlagRequired("email")
	constituentCreateCmd.MarkFlagRequired("original-source-id")
	constituentCreateCmd.MarkFlagRequired("constituent-type-id")
}

func runConstituentCreate(cmd *cobra.Command, args []string) error {
	client, err := loadClient()
	if err != nil {
		return err
	}

	firstName := createFlags.firstName
	if len(firstName) > 20 {
		firstName = firstName[:20]
	}

	lastName := createFlags.lastName
	if len(lastName) > 55 {
		lastName = lastName[:55]
	}

	postalCode := createFlags.postalCode
	if len(postalCode) > 10 {
		postalCode = postalCode[:10]
	}

	params := tessitura.CreateConstituentParams{
		FirstName:         firstName,
		LastName:          lastName,
		Email:             createFlags.email,
		ConstituentTypeId: createFlags.constituentTypeId,
		OriginalSourceId:  createFlags.originalSourceId,
		Street:            createFlags.street,
		PostalCode:        postalCode,
	}

	detail, err := client.CreateConstituent(cmd.Context(), params)
	if err != nil {
		return err
	}

	constituent := domain.ConstituentFromAPI(detail)

	out, err := json.MarshalIndent(constituent, "", "  ")
	if err != nil {
		return fmt.Errorf("unable to format output: %w", err)
	}

	fmt.Fprintln(cmd.OutOrStdout(), string(out))
	return nil
}
