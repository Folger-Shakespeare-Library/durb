package domain

import "github.com/Folger-Shakespeare-Library/durb/pkg/tessitura"

// ReportResultRef is a reference sub-object used in ReportResult (category, schedule).
type ReportResultRef struct {
	ID               int     `json:"id"`
	Description      *string `json:"description"`
	AliasDescription *string `json:"aliasDescription,omitempty"`
}

// ReportResultReportRef is the report sub-object in a ReportResult.
type ReportResultReportRef struct {
	ID           *string `json:"id"`
	Name         *string `json:"name"`
	Window       *string `json:"window,omitempty"`
	ReportTypeID *int    `json:"reportTypeId"`
}

// ReportResult is the domain representation of a scheduled report result —
// a combined entity from ReportRequest, ReportSchedule, and Report.
type ReportResult struct {
	ID              int                    `json:"id"`
	Report          *ReportResultReportRef `json:"report"`
	Category        *ReportResultRef       `json:"category"`
	Schedule        *ReportResultRef       `json:"schedule"`
	RequestDateTime *string                `json:"requestDateTime"`
	EndDateTime     *string                `json:"endDateTime"`
	OutputOption    *string                `json:"outputOption"`
	IsPrinted       bool                   `json:"isPrinted"`
	UserID          *string                `json:"userId"`
	IsPublic        bool                   `json:"isPublic"`
	IsDeleted       bool                   `json:"isDeleted"`
	ResultCode      *string                `json:"resultCode"`
	ResultText      *string                `json:"resultText"`
	Type            *string                `json:"type"`
	QueueStatus     *string                `json:"queueStatus"`
	LastRequest     *int                   `json:"lastRequest"`
}

// ReportResultFromAPI maps a raw Tessitura API report result to the domain type.
func ReportResultFromAPI(r *tessitura.APIReportResult) *ReportResult {
	result := &ReportResult{
		RequestDateTime: r.RequestDateTime,
		EndDateTime:     r.EndDateTime,
		OutputOption:    r.OutputOption,
		UserID:          r.UserId,
		ResultCode:      r.ResultCode,
		ResultText:      r.ResultText,
		Type:            r.Type,
		QueueStatus:     r.QueueStatus,
		LastRequest:     r.LastRequest,
	}
	if r.Id != nil {
		result.ID = *r.Id
	}
	if r.IsPrinted != nil {
		result.IsPrinted = *r.IsPrinted
	}
	if r.IsPublic != nil {
		result.IsPublic = *r.IsPublic
	}
	if r.IsDeleted != nil {
		result.IsDeleted = *r.IsDeleted
	}
	if r.Report != nil {
		result.Report = &ReportResultReportRef{
			ID:           r.Report.Id,
			Name:         r.Report.Name,
			Window:       r.Report.Window,
			ReportTypeID: r.Report.ReportTypeId,
		}
	}
	if r.ReportCategory != nil {
		ref := &ReportResultRef{Description: r.ReportCategory.Description, AliasDescription: r.ReportCategory.AliasDescription}
		if r.ReportCategory.Id != nil {
			ref.ID = *r.ReportCategory.Id
		}
		result.Category = ref
	}
	if r.Schedule != nil {
		ref := &ReportResultRef{Description: r.Schedule.Description, AliasDescription: r.Schedule.AliasDescription}
		if r.Schedule.Id != nil {
			ref.ID = *r.Schedule.Id
		}
		result.Schedule = ref
	}
	return result
}

// ReportRequestParameter holds a single parameter value for a report request.
type ReportRequestParameter struct {
	ParameterID *int    `json:"parameterId"`
	ReportID    *string `json:"reportId"`
	Value       *string `json:"value"`
}

// ReportRequest is the full domain representation of a Tessitura report request.
type ReportRequest struct {
	ID              int                       `json:"id"`
	Name            *string                   `json:"name"`
	ReportID        *string                   `json:"reportId"`
	ReportType      *int                      `json:"reportType"`
	Type            *string                   `json:"type"`
	QueueStatus     *string                   `json:"queueStatus"`
	ResultCode      *string                   `json:"resultCode"`
	ResultText      *string                   `json:"resultText"`
	RequestDateTime *string                   `json:"requestDateTime"`
	EndDateTime     *string                   `json:"endDateTime"`
	OutputOption    *string                   `json:"outputOption"`
	PublicIndicator bool                      `json:"publicIndicator"`
	DeletedIndicator bool                     `json:"deletedIndicator"`
	ScheduleID      *int                      `json:"scheduleId"`
	HeaderRequestID *int                      `json:"headerRequestId"`
	UserID          *string                   `json:"userId"`
	UserGroupID     *string                   `json:"userGroupId"`
	EmailRecipients *string                   `json:"emailRecipients"`
	EmailSubject    *string                   `json:"emailSubject"`
	EmailBody       *string                   `json:"emailBody"`
	Parameters      []*ReportRequestParameter `json:"parameters,omitempty"`
	CreatedAt       *string                   `json:"createdAt"`
	CreatedBy       *string                   `json:"createdBy"`
	UpdatedAt       *string                   `json:"updatedAt"`
	UpdatedBy       *string                   `json:"updatedBy"`
}

// ReportRequestFromAPI maps a raw Tessitura API report request to the domain type.
func ReportRequestFromAPI(r *tessitura.APIReportRequest) *ReportRequest {
	req := &ReportRequest{
		Name:            r.Name,
		ReportID:        r.ReportId,
		ReportType:      r.ReportType,
		Type:            r.Type,
		QueueStatus:     r.QueueStatus,
		ResultCode:      r.ResultCode,
		ResultText:      r.ResultText,
		RequestDateTime: r.RequestDateTime,
		EndDateTime:     r.EndDateTime,
		OutputOption:    r.OutputOption,
		ScheduleID:      r.ScheduleId,
		HeaderRequestID: r.HeaderRequestId,
		UserID:          r.UserId,
		UserGroupID:     r.UserGroupId,
		EmailRecipients: r.EmailRecipients,
		EmailSubject:    r.EmailSubject,
		EmailBody:       r.EmailBody,
		CreatedAt:       r.CreatedDateTime,
		CreatedBy:       r.CreatedBy,
		UpdatedAt:       r.UpdatedDateTime,
		UpdatedBy:       r.UpdatedBy,
	}
	if r.Id != nil {
		req.ID = *r.Id
	}
	if r.PublicIndicator != nil {
		req.PublicIndicator = *r.PublicIndicator
	}
	if r.DeletedIndicator != nil {
		req.DeletedIndicator = *r.DeletedIndicator
	}
	return req
}

// AttachRequestDetail merges parameter values from a ReportRequestDetail response
// into the domain ReportRequest.
func (r *ReportRequest) AttachRequestDetail(d *tessitura.APIReportRequestDetail) {
	if d == nil {
		return
	}
	for _, p := range d.Parameters {
		if p == nil {
			continue
		}
		r.Parameters = append(r.Parameters, &ReportRequestParameter{
			ParameterID: p.ParameterId,
			ReportID:    p.ReportId,
			Value:       p.Value,
		})
	}
}
