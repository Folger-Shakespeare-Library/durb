package cli

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var refSeatStatusesCmd = &cobra.Command{
	Use:   "seat-statuses",
	Short: "Seat statuses",
	Long:  "Commands for seat status reference data.",
}

var refSeatStatusesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List seat statuses",
	Long:  "List available seat statuses.",
	RunE:  runRefSeatStatusesList,
}

func init() {
	refSeatStatusesCmd.AddCommand(refSeatStatusesListCmd)
}

func runRefSeatStatusesList(cmd *cobra.Command, args []string) error {
	client, err := loadClient()
	if err != nil {
		return err
	}

	items, err := client.GetSeatStatuses(cmd.Context())
	if err != nil {
		return err
	}

	type seatStatus struct {
		Id             int    `json:"id"`
		Description    string `json:"description"`
		StatusCode     string `json:"statusCode"`
		StatusLegend   string `json:"statusLegend"`
		StatusPriority int    `json:"statusPriority"`
		Inactive       bool   `json:"inactive"`
	}

	var out []seatStatus
	for _, item := range items {
		r := seatStatus{
			Inactive:       item.Inactive,
			StatusPriority: item.StatusPriority,
		}
		if item.Id != nil {
			r.Id = *item.Id
		}
		if item.Description != nil {
			r.Description = *item.Description
		}
		if item.StatusCode != nil {
			r.StatusCode = *item.StatusCode
		}
		if item.StatusLegend != nil {
			r.StatusLegend = *item.StatusLegend
		}
		out = append(out, r)
	}

	data, err := json.MarshalIndent(out, "", "  ")
	if err != nil {
		return fmt.Errorf("unable to format output: %w", err)
	}

	fmt.Fprintln(cmd.OutOrStdout(), string(data))
	return nil
}
