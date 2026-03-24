package tessitura

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
)

type APIReportCategorySummary struct {
	Id          *int    `json:"Id"`
	Description *string `json:"Description"`
}

type APIReportTypeSummary struct {
	Id          *int    `json:"Id"`
	Description *string `json:"Description"`
}

type APIReport struct {
	Id                     *string                   `json:"Id"`
	Name                   *string                   `json:"Name"`
	Description            *string                   `json:"Description"`
	ReportPath             *string                   `json:"ReportPath"`
	Category               *APIReportCategorySummary `json:"Category"`
	ReportType             *APIReportTypeSummary     `json:"ReportType"`
	AllowSchedule          *bool                     `json:"AllowSchedule"`
	AllowQuery             *bool                     `json:"AllowQuery"`
	QueryStringAppend      *string                   `json:"QueryStringAppend"`
	ParameterWindow        *string                   `json:"ParameterWindow"`
	ParameterWindowIndicator *bool                   `json:"ParameterWindowIndicator"`
	PublicIndicator        *bool                     `json:"PublicIndicator"`
	WarningIndicator       *bool                     `json:"WarningIndicator"`
	UtilityIndicator       *bool                     `json:"UtilityIndicator"`
	Window                 *string                   `json:"Window"`
	ApplicationId          *string                   `json:"ApplicationId"`
	Inactive               *bool                     `json:"Inactive"`
	LastRequestId          *int                      `json:"LastRequestId"`
	CreatedDateTime        *string                   `json:"CreatedDateTime"`
	CreateLocation         *string                   `json:"CreateLocation"`
	CreatedBy              *string                   `json:"CreatedBy"`
	UpdatedDateTime        *string                   `json:"UpdatedDateTime"`
	UpdatedBy              *string                   `json:"UpdatedBy"`
}

type APIReportParameter struct {
	Id             *int    `json:"Id"`
	ParameterName  *string `json:"ParameterName"`
	Description    *string `json:"Description"`
	DataType       *int    `json:"DataType"`
	Required       *bool   `json:"Required"`
	DefaultValue   *string `json:"DefaultValue"`
	MultiSelect    *bool   `json:"MultiSelect"`
	SequenceNumber *int    `json:"SequenceNumber"`
	Inactive       *bool   `json:"Inactive"`
}

type APIReportDetail struct {
	Id                       *string                   `json:"Id"`
	Name                     *string                   `json:"Name"`
	Description              *string                   `json:"Description"`
	ReportPath               *string                   `json:"ReportPath"`
	Category                 *APIReportCategorySummary `json:"Category"`
	ReportType               *APIReportTypeSummary     `json:"ReportType"`
	AllowSchedule            *bool                     `json:"AllowSchedule"`
	ParameterWindow          *string                   `json:"ParameterWindow"`
	ParameterWindowIndicator *bool                     `json:"ParameterWindowIndicator"`
	PublicIndicator          *bool                     `json:"PublicIndicator"`
	WarningIndicator         *bool                     `json:"WarningIndicator"`
	UtilityIndicator         *bool                     `json:"UtilityIndicator"`
	Window                   *string                   `json:"Window"`
	ApplicationId            *string                   `json:"ApplicationId"`
	Inactive                 *bool                     `json:"Inactive"`
	LastRequestId            *int                      `json:"LastRequestId"`
	CreatedDateTime          *string                   `json:"CreatedDateTime"`
	CreateLocation           *string                   `json:"CreateLocation"`
	CreatedBy                *string                   `json:"CreatedBy"`
	UpdatedDateTime          *string                   `json:"UpdatedDateTime"`
	UpdatedBy                *string                   `json:"UpdatedBy"`
	Parameters               []*APIReportParameter     `json:"Parameters"`
}

// ReportResult holds both the base report and its detail, fetched via batch.
type ReportResult struct {
	Base   *APIReport
	Detail *APIReportDetail
}

// GetReports fetches all reports, optionally filtered by type and/or category IDs
// (comma-delimited strings, e.g. "1,2,3").
func (c *Client) GetReports(ctx context.Context, typeIds, categoryIds string) ([]*APIReport, error) {
	path := "/Reporting/Reports"
	sep := "?"
	if typeIds != "" {
		path += sep + "typeIds=" + typeIds
		sep = "&"
	}
	if categoryIds != "" {
		path += sep + "categoryIds=" + categoryIds
	}

	data, err := c.Get(ctx, path)
	if err != nil {
		return nil, err
	}

	var reports []*APIReport
	if err := json.Unmarshal(data, &reports); err != nil {
		return nil, fmt.Errorf("unable to parse reports: %w", err)
	}
	return reports, nil
}

// GetReport fetches a single report by ID, batching the base report and its
// detail in a single HTTP call. Request IDs: 1=base, 2=detail.
func (c *Client) GetReport(ctx context.Context, id string) (*ReportResult, error) {
	batchResp, err := c.Batch(ctx, []BatchRequestItem{
		{HttpMethod: "GET", Id: 1, Uri: "/Reporting/Reports/" + id},
		{HttpMethod: "GET", Id: 2, Uri: "/Reporting/Reports/" + id + "/Details"},
	})
	if err != nil {
		return nil, err
	}

	result := &ReportResult{}
	for _, item := range batchResp.Responses {
		switch item.RequestId {
		case 1:
			var r APIReport
			if err := json.Unmarshal(item.ResponseObject, &r); err != nil {
				return nil, fmt.Errorf("unable to parse report %s: %w", id, err)
			}
			result.Base = &r
		case 2:
			var d APIReportDetail
			if err := json.Unmarshal(item.ResponseObject, &d); err != nil {
				return nil, fmt.Errorf("unable to parse report detail %s: %w", id, err)
			}
			result.Detail = &d
		}
	}

	return result, nil
}

// GetReportsBatch fetches multiple reports concurrently, each with full detail.
// Results are returned in the same order as the input IDs.
func (c *Client) GetReportsBatch(ctx context.Context, ids []string) ([]*ReportResult, error) {
	results := make([]*ReportResult, len(ids))
	errs := make([]error, len(ids))

	var wg sync.WaitGroup
	for i, id := range ids {
		wg.Add(1)
		go func(idx int, rid string) {
			defer wg.Done()
			report, err := c.GetReport(ctx, rid)
			results[idx] = report
			errs[idx] = err
		}(i, id)
	}
	wg.Wait()

	for i, err := range errs {
		if err != nil {
			return nil, fmt.Errorf("report %s: %w", ids[i], err)
		}
	}

	return results, nil
}
