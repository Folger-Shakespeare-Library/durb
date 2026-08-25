package cli

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var setStatusFlags struct {
	id     int
	status string
	reason string
}

var constituentSetStatusCmd = &cobra.Command{
	Use:   "set-status",
	Short: "Change a constituent's inactive status",
	Long: `Set the inactive status on a constituent (e.g. Active, Inactive, Purge, Merged).

Looks up the status and reason by description from reference data, so the
exact available values depend on your Tessitura installation.

The --reason flag is required when inactivating a constituent.

Examples:
  tess constituent set-status --id 12345 --status "Inactive" --reason "Duplicate"
  tess constituent set-status --id 12345 --status "Active"`,
	RunE: runConstituentSetStatus,
}

func init() {
	f := constituentSetStatusCmd.Flags()
	f.IntVar(&setStatusFlags.id, "id", 0, "constituent ID (required)")
	f.StringVar(&setStatusFlags.status, "status", "", "status description, e.g. Active, Inactive (required)")
	f.StringVar(&setStatusFlags.reason, "reason", "", "inactive reason description (required when inactivating)")

	constituentSetStatusCmd.MarkFlagRequired("id")
	constituentSetStatusCmd.MarkFlagRequired("status")
}

func runConstituentSetStatus(cmd *cobra.Command, args []string) error {
	client, err := loadClient()
	if err != nil {
		return err
	}

	statuses, err := client.GetConstituentInactiveStatuses(cmd.Context())
	if err != nil {
		return fmt.Errorf("fetching inactive statuses: %w", err)
	}

	inactiveId := -999
	for _, s := range statuses {
		if s.Inactive != nil && *s.Inactive {
			continue
		}
		if s.Description != nil && strings.EqualFold(*s.Description, setStatusFlags.status) {
			if s.Id != nil {
				inactiveId = *s.Id
			}
			break
		}
	}
	if inactiveId == -999 {
		var valid []string
		for _, s := range statuses {
			if s.Inactive != nil && *s.Inactive {
				continue
			}
			if s.Description != nil {
				valid = append(valid, *s.Description)
			}
		}
		return fmt.Errorf("unknown status %q; valid values: %s", setStatusFlags.status, strings.Join(valid, ", "))
	}

	isActive := inactiveId == 1
	if isActive && setStatusFlags.reason != "" {
		return fmt.Errorf("--reason cannot be used when setting status to Active")
	}
	if !isActive && setStatusFlags.reason == "" {
		return fmt.Errorf("--reason is required when setting status to %q", setStatusFlags.status)
	}

	var reasonId *int
	if setStatusFlags.reason != "" {
		reasons, err := client.GetConstituentInactiveReasons(cmd.Context())
		if err != nil {
			return fmt.Errorf("fetching inactive reasons: %w", err)
		}

		found := false
		for _, r := range reasons {
			if r.Description != nil && strings.EqualFold(*r.Description, setStatusFlags.reason) {
				if r.Id != nil {
					reasonId = r.Id
					found = true
				}
				break
			}
		}
		if !found {
			var valid []string
			for _, r := range reasons {
				if r.Description != nil {
					valid = append(valid, *r.Description)
				}
			}
			return fmt.Errorf("unknown reason %q; valid values: %s", setStatusFlags.reason, strings.Join(valid, ", "))
		}
	}

	idStr := fmt.Sprintf("%d", setStatusFlags.id)
	return client.SetConstituentStatus(cmd.Context(), idStr, inactiveId, reasonId)
}
