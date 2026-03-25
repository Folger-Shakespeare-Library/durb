package tessitura

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
)

type APIReportRequest struct {
	Id               *int    `json:"Id"`
	DeletedIndicator *bool   `json:"DeletedIndicator"`
	EmailBody        *string `json:"EmailBody"`
	Name             *string `json:"Name"`
	EmailRecipients  *string `json:"EmailRecipients"`
	EmailSubject     *string `json:"EmailSubject"`
	EndDateTime      *string `json:"EndDateTime"`
	OutputOption     *string `json:"OutputOption"`
	PublicIndicator  *bool   `json:"PublicIndicator"`
	ReportId         *string `json:"ReportId"`
	ReportType       *int    `json:"ReportType"`
	RequestDateTime  *string `json:"RequestDateTime"`
	ResultCode       *string `json:"ResultCode"`
	ResultText       *string `json:"ResultText"`
	Type             *string `json:"Type"`
	QueueStatus      *string `json:"QueueStatus"`
	ScheduleId       *int    `json:"ScheduleId"`
	HeaderRequestId  *int    `json:"HeaderRequestId"`
	UserGroupId      *string `json:"UserGroupId"`
	UserId           *string `json:"UserId"`
	CreatedDateTime  *string `json:"CreatedDateTime"`
	CreateLocation   *string `json:"CreateLocation"`
	CreatedBy        *string `json:"CreatedBy"`
	UpdatedDateTime  *string `json:"UpdatedDateTime"`
	UpdatedBy        *string `json:"UpdatedBy"`
}

type APIReportRequestParameter struct {
	RequestId   *int    `json:"RequestId"`
	ReportId    *string `json:"ReportId"`
	ParameterId *int    `json:"ParameterId"`
	Value       *string `json:"Value"`
}

type APIReportRequestDetail struct {
	Id               *int                         `json:"Id"`
	DeletedIndicator *bool                        `json:"DeletedIndicator"`
	EmailBody        *string                      `json:"EmailBody"`
	Name             *string                      `json:"Name"`
	EmailRecipients  *string                      `json:"EmailRecipients"`
	EmailSubject     *string                      `json:"EmailSubject"`
	EndDateTime      *string                      `json:"EndDateTime"`
	OutputOption     *string                      `json:"OutputOption"`
	PublicIndicator  *bool                        `json:"PublicIndicator"`
	ReportId         *string                      `json:"ReportId"`
	ReportType       *int                         `json:"ReportType"`
	RequestDateTime  *string                      `json:"RequestDateTime"`
	ResultCode       *string                      `json:"ResultCode"`
	ResultText       *string                      `json:"ResultText"`
	Type             *string                      `json:"Type"`
	QueueStatus      *string                      `json:"QueueStatus"`
	ScheduleId       *int                         `json:"ScheduleId"`
	HeaderRequestId  *int                         `json:"HeaderRequestId"`
	UserGroupId      *string                      `json:"UserGroupId"`
	UserId           *string                      `json:"UserId"`
	CreatedDateTime  *string                      `json:"CreatedDateTime"`
	CreateLocation   *string                      `json:"CreateLocation"`
	CreatedBy        *string                      `json:"CreatedBy"`
	UpdatedDateTime  *string                      `json:"UpdatedDateTime"`
	UpdatedBy        *string                      `json:"UpdatedBy"`
	Parameters       []*APIReportRequestParameter `json:"Parameters"`
}

type APIEntitySummary struct {
	Id               *int    `json:"Id"`
	Description      *string `json:"Description"`
	AliasDescription *string `json:"AliasDescription"`
}

type APIReportSummary struct {
	Id           *string `json:"Id"`
	Name         *string `json:"Name"`
	Window       *string `json:"Window"`
	ReportTypeId *int    `json:"ReportTypeId"`
}

type APIReportResult struct {
	Id              *int              `json:"Id"`
	Report          *APIReportSummary `json:"Report"`
	ReportCategory  *APIEntitySummary `json:"ReportCategory"`
	Schedule        *APIEntitySummary `json:"Schedule"`
	RequestDateTime *string           `json:"RequestDateTime"`
	EndDateTime     *string           `json:"EndDateTime"`
	OutputOption    *string           `json:"OutputOption"`
	IsPrinted       *bool             `json:"IsPrinted"`
	UserId          *string           `json:"UserId"`
	IsPublic        *bool             `json:"IsPublic"`
	IsDeleted       *bool             `json:"IsDeleted"`
	ResultCode      *string           `json:"ResultCode"`
	ResultText      *string           `json:"ResultText"`
	Type            *string           `json:"Type"`
	QueueStatus     *string           `json:"QueueStatus"`
	LastRequest     *int              `json:"LastRequest"`
}

type APIReportResultsResponse struct {
	TotalCount    *int               `json:"TotalCount"`
	Page          *int               `json:"Page"`
	PageSize      *int               `json:"PageSize"`
	ReportResults []*APIReportResult `json:"ReportResults"`
}

// ReportRequestResult holds both the base request and its detail, fetched via batch.
type ReportRequestResult struct {
	Base   *APIReportRequest
	Detail *APIReportRequestDetail
}

// GetReportResults fetches scheduled report results with optional filters.
// Returns the results page and the total count.
func (c *Client) GetReportResults(ctx context.Context, params ReportResultsParams) ([]*APIReportResult, int, error) {
	path := "/Reporting/ReportRequests/Results?"
	sep := ""
	add := func(key, val string) {
		if val != "" {
			path += sep + key + "=" + val
			sep = "&"
		}
	}

	add("scheduleName", params.ScheduleName)
	add("reportId", params.ReportId)
	add("startDate", params.StartDate)
	add("endDate", params.EndDate)
	if params.IncludePublic {
		add("includePublic", "true")
	}
	if params.MyReportsOnly {
		add("myReportsOnly", "true")
	}
	if params.RecentResultsOnly {
		add("recentResultsOnly", "true")
	}
	if params.IncludeErrors {
		add("includeErrors", "true")
	}
	if params.IncludeDeleted {
		add("includeDeleted", "true")
	}
	add("page", fmt.Sprintf("%d", params.Page))
	add("pageSize", fmt.Sprintf("%d", params.PageSize))

	data, err := c.Get(ctx, path)
	if err != nil {
		return nil, 0, err
	}

	var resp APIReportResultsResponse
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, 0, fmt.Errorf("unable to parse report results: %w", err)
	}

	total := 0
	if resp.TotalCount != nil {
		total = *resp.TotalCount
	}
	return resp.ReportResults, total, nil
}

// ReportResultsParams holds query parameters for GetReportResults.
type ReportResultsParams struct {
	ScheduleName      string
	ReportId          string
	StartDate         string
	EndDate           string
	IncludePublic     bool
	MyReportsOnly     bool
	RecentResultsOnly bool
	IncludeErrors     bool
	IncludeDeleted    bool
	Page              int
	PageSize          int
}

// GetReportRequests fetches all report requests. Pass activeOnly=true to return
// only scheduled or currently running requests.
func (c *Client) GetReportRequests(ctx context.Context, activeOnly bool) ([]*APIReportRequest, error) {
	path := "/Reporting/ReportRequests"
	if activeOnly {
		path += "?activeOnly=true"
	}

	data, err := c.Get(ctx, path)
	if err != nil {
		return nil, err
	}

	var results []*APIReportRequest
	if err := json.Unmarshal(data, &results); err != nil {
		return nil, fmt.Errorf("unable to parse report requests: %w", err)
	}
	return results, nil
}

// GetReportRequest fetches a single report request by ID, batching the base
// request and its detail in a single HTTP call. Request IDs: 1=base, 2=detail.
func (c *Client) GetReportRequest(ctx context.Context, id string) (*ReportRequestResult, error) {
	batchResp, err := c.Batch(ctx, []BatchRequestItem{
		{HttpMethod: "GET", Id: 1, Uri: "/Reporting/ReportRequests/" + id},
		{HttpMethod: "GET", Id: 2, Uri: "/Reporting/ReportRequests/" + id + "/Details"},
	})
	if err != nil {
		return nil, err
	}

	result := &ReportRequestResult{}
	for _, item := range batchResp.Responses {
		switch item.RequestId {
		case 1:
			var r APIReportRequest
			if err := json.Unmarshal(item.ResponseObject, &r); err != nil {
				return nil, fmt.Errorf("unable to parse report request %s: %w", id, err)
			}
			result.Base = &r
		case 2:
			var d APIReportRequestDetail
			if err := json.Unmarshal(item.ResponseObject, &d); err != nil {
				return nil, fmt.Errorf("unable to parse report request detail %s: %w", id, err)
			}
			result.Detail = &d
		}
	}

	return result, nil
}

// GetReportRequestsBatch fetches multiple report requests concurrently.
// Results are returned in the same order as the input IDs.
func (c *Client) GetReportRequestsBatch(ctx context.Context, ids []string) ([]*ReportRequestResult, error) {
	results := make([]*ReportRequestResult, len(ids))
	errs := make([]error, len(ids))

	var wg sync.WaitGroup
	for i, id := range ids {
		wg.Add(1)
		go func(idx int, rid string) {
			defer wg.Done()
			req, err := c.GetReportRequest(ctx, rid)
			results[idx] = req
			errs[idx] = err
		}(i, id)
	}
	wg.Wait()

	for i, err := range errs {
		if err != nil {
			return nil, fmt.Errorf("report request %s: %w", ids[i], err)
		}
	}

	return results, nil
}
