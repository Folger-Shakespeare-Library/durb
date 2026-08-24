package cli

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/Folger-Shakespeare-Library/durb/pkg/domain"
	"github.com/spf13/cobra"
)

var updateFlags struct {
	id             int
	email          string
	allowMarketing bool
}

var constituentUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a constituent's properties",
	Long: `Update properties on an existing constituent.

Currently supports updating the email marketing flag.
Use --allow-marketing (defaults to true) or --allow-marketing=false.
Requires --id and --email to identify the electronic address record.

Examples:
  tess constituent update --id 12345 --email "jane@example.com" --allow-marketing
  tess constituent update --id 12345 --email "jane@example.com" --allow-marketing=false`,
	RunE: runConstituentUpdate,
}

func init() {
	f := constituentUpdateCmd.Flags()
	f.IntVar(&updateFlags.id, "id", 0, "constituent ID (required)")
	f.StringVar(&updateFlags.email, "email", "", "email address to update (required)")
	f.BoolVar(&updateFlags.allowMarketing, "allow-marketing", true, "allow email marketing (use =false to disable)")

	constituentUpdateCmd.MarkFlagRequired("id")
	constituentUpdateCmd.MarkFlagRequired("email")
}

func runConstituentUpdate(cmd *cobra.Command, args []string) error {
	if !cmd.Flags().Changed("allow-marketing") {
		return fmt.Errorf("no update flags specified (e.g. --allow-marketing)")
	}

	client, err := loadClient()
	if err != nil {
		return err
	}

	addresses, err := client.GetElectronicAddresses(cmd.Context(), updateFlags.id, updateFlags.email)
	if err != nil {
		return fmt.Errorf("unable to fetch electronic addresses: %w", err)
	}

	var matchId int
	var matchUpdated *string
	var matchPrimary *bool
	found := false
	for _, addr := range addresses {
		if addr.Address != nil && strings.EqualFold(*addr.Address, updateFlags.email) {
			if addr.Id != nil {
				matchId = *addr.Id
			}
			matchUpdated = addr.UpdatedDateTime
			matchPrimary = addr.PrimaryIndicator
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("email %q not found on constituent %d", updateFlags.email, updateFlags.id)
	}

	primary := false
	if matchPrimary != nil {
		primary = *matchPrimary
	}

	body := map[string]interface{}{
		"Id":                    matchId,
		"Constituent":          map[string]int{"Id": updateFlags.id},
		"Address":              updateFlags.email,
		"ElectronicAddressType": map[string]int{"Id": 1},
		"PrimaryIndicator":     primary,
		"AllowMarketing":       updateFlags.allowMarketing,
		"UpdatedDateTime":      matchUpdated,
	}

	if err := client.UpdateElectronicAddress(cmd.Context(), matchId, body); err != nil {
		return fmt.Errorf("unable to update marketing flag: %w", err)
	}

	detail, err := client.GetConstituentDetail(cmd.Context(), strconv.Itoa(updateFlags.id))
	if err != nil {
		return fmt.Errorf("unable to fetch updated constituent: %w", err)
	}

	constituent := domain.ConstituentFromAPI(detail)
	out, err := json.MarshalIndent(constituent, "", "  ")
	if err != nil {
		return fmt.Errorf("unable to format output: %w", err)
	}

	fmt.Fprintln(cmd.OutOrStdout(), string(out))
	return nil
}
