package cli

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var interestEnableFlags struct {
	constituentId int
	interestIds   string
}

var interestEnableCmd = &cobra.Command{
	Use:   "enable",
	Short: "Enable interests on a constituent",
	Long: `Enable one or more interests on a constituent by interest type ID.

Creates the interest assignment if it doesn't exist, or updates it if it does.

Examples:
  tess interest enable --constituent-id 446106 --interest-type-ids 262
  tess interest enable --constituent-id 446106 --interest-type-ids 262,263,264`,
	RunE: runInterestEnable,
}

func init() {
	f := interestEnableCmd.Flags()
	f.IntVar(&interestEnableFlags.constituentId, "constituent-id", 0, "constituent ID (required)")
	f.StringVar(&interestEnableFlags.interestIds, "interest-type-ids", "", "comma-separated interest type IDs (required)")

	interestEnableCmd.MarkFlagRequired("constituent-id")
	interestEnableCmd.MarkFlagRequired("interest-type-ids")
}

func runInterestEnable(cmd *cobra.Command, args []string) error {
	return setInterests(cmd, interestEnableFlags.constituentId, interestEnableFlags.interestIds, true)
}

func setInterests(cmd *cobra.Command, constituentId int, interestIdsStr string, selected bool) error {
	typeIds, err := parseIntList(interestIdsStr)
	if err != nil {
		return fmt.Errorf("invalid --interest-type-ids: %w", err)
	}

	client, err := loadClient()
	if err != nil {
		return err
	}

	existing, err := client.GetInterests(cmd.Context(), constituentId)
	if err != nil {
		return fmt.Errorf("unable to fetch interests: %w", err)
	}

	byTypeId := make(map[int]struct {
		id              int
		updatedDateTime *string
	})
	for _, interest := range existing {
		if interest.InterestType != nil && interest.InterestType.Id != nil && interest.Id != nil {
			byTypeId[*interest.InterestType.Id] = struct {
				id              int
				updatedDateTime *string
			}{*interest.Id, interest.UpdatedDateTime}
		}
	}

	for _, typeId := range typeIds {
		if match, ok := byTypeId[typeId]; ok {
			body := map[string]interface{}{
				"Id":              match.id,
				"Constituent":    map[string]int{"Id": constituentId},
				"InterestType":   map[string]int{"Id": typeId},
				"Selected":       selected,
				"UpdatedDateTime": match.updatedDateTime,
			}
			if err := client.UpdateInterest(cmd.Context(), match.id, body); err != nil {
				return fmt.Errorf("unable to update interest type %d: %w", typeId, err)
			}
		} else {
			body := map[string]interface{}{
				"Constituent":  map[string]int{"Id": constituentId},
				"InterestType": map[string]int{"Id": typeId},
				"Selected":     selected,
			}
			if err := client.CreateInterest(cmd.Context(), body); err != nil {
				return fmt.Errorf("unable to create interest type %d: %w", typeId, err)
			}
		}
	}

	return nil
}

func parseIntList(s string) ([]int, error) {
	parts := strings.Split(s, ",")
	result := make([]int, 0, len(parts))
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p == "" {
			continue
		}
		n, err := strconv.Atoi(p)
		if err != nil {
			return nil, fmt.Errorf("%q is not a valid integer", p)
		}
		result = append(result, n)
	}
	if len(result) == 0 {
		return nil, fmt.Errorf("no IDs provided")
	}
	return result, nil
}
