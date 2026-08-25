package cli

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/Folger-Shakespeare-Library/durb/pkg/tessitura"
	"github.com/spf13/cobra"
)

var activityCreateFlags struct {
	constituentId  int
	activityTypeId int
	statusId       int
	datetime       string
	notes          string
	unique         bool
}

var activityCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create an activity on a constituent",
	Long: `Create a special activity record on a constituent.

Accepts a full ISO 8601 datetime or a bare date (YYYY-MM-DD), which defaults
to midnight in the system's local timezone.

Use --unique to skip creation if an activity with the same type and datetime
already exists on the constituent.

Examples:
  tess activity create --constituent-id 446106 --activity-type-id 22 --status-id 9 --datetime "2025-01-15T10:00:00-05:00"
  tess activity create --constituent-id 446106 --activity-type-id 22 --status-id 9 --datetime "2025-01-15"
  tess activity create --constituent-id 446106 --activity-type-id 22 --status-id 9 --datetime "2025-01-15" --unique`,
	RunE: runActivityCreate,
}

func init() {
	f := activityCreateCmd.Flags()
	f.IntVar(&activityCreateFlags.constituentId, "constituent-id", 0, "constituent ID (required)")
	f.IntVar(&activityCreateFlags.activityTypeId, "activity-type-id", 0, "activity type ID (required)")
	f.IntVar(&activityCreateFlags.statusId, "status-id", 0, "activity status ID (required)")
	f.StringVar(&activityCreateFlags.datetime, "datetime", "", "date or datetime in ISO 8601 format (required)")
	f.StringVar(&activityCreateFlags.notes, "notes", "", "activity notes")
	f.BoolVar(&activityCreateFlags.unique, "unique", false, "skip if an activity with the same type and datetime already exists")

	activityCreateCmd.MarkFlagRequired("constituent-id")
	activityCreateCmd.MarkFlagRequired("activity-type-id")
	activityCreateCmd.MarkFlagRequired("status-id")
	activityCreateCmd.MarkFlagRequired("datetime")
}

func normalizeDateTime(input string) (string, error) {
	input = strings.TrimSpace(input)

	if !strings.Contains(input, "T") {
		d, err := time.Parse("2006-01-02", input)
		if err != nil {
			return "", fmt.Errorf("invalid date %q: expected YYYY-MM-DD", input)
		}
		return d.Format(time.RFC3339), nil
	}

	for _, layout := range []string{time.RFC3339, "2006-01-02T15:04:05-0700"} {
		t, err := time.Parse(layout, input)
		if err == nil {
			return t.Format(time.RFC3339), nil
		}
	}
	return "", fmt.Errorf("invalid datetime %q: expected ISO 8601 format", input)
}

func runActivityCreate(cmd *cobra.Command, args []string) error {
	dt, err := normalizeDateTime(activityCreateFlags.datetime)
	if err != nil {
		return err
	}

	client, err := loadClient()
	if err != nil {
		return err
	}

	if activityCreateFlags.unique {
		existing, err := client.GetActivities(cmd.Context(), activityCreateFlags.constituentId, activityCreateFlags.activityTypeId)
		if err != nil {
			return fmt.Errorf("unable to check existing activities: %w", err)
		}

		inputTime, _ := time.Parse(time.RFC3339, dt)
		for _, a := range existing {
			if a.SpecialActivityDateTime == nil {
				continue
			}
			existingTime, err := time.Parse(time.RFC3339, *a.SpecialActivityDateTime)
			if err != nil {
				continue
			}
			if inputTime.Equal(existingTime) {
				out, err := json.MarshalIndent(a, "", "  ")
				if err != nil {
					return fmt.Errorf("unable to format output: %w", err)
				}
				fmt.Fprintln(cmd.OutOrStdout(), string(out))
				return nil
			}
		}
	}

	params := tessitura.CreateActivityParams{
		ConstituentId:  activityCreateFlags.constituentId,
		ActivityTypeId: activityCreateFlags.activityTypeId,
		StatusId:       activityCreateFlags.statusId,
		DateTime:       dt,
		Notes:          activityCreateFlags.notes,
	}

	activity, err := client.CreateActivity(cmd.Context(), params)
	if err != nil {
		return err
	}

	out, err := json.MarshalIndent(activity, "", "  ")
	if err != nil {
		return fmt.Errorf("unable to format output: %w", err)
	}

	fmt.Fprintln(cmd.OutOrStdout(), string(out))
	return nil
}
