package domain

import "github.com/Folger-Shakespeare-Library/durb/pkg/tessitura"

type ReportRef struct {
	ID          int     `json:"id"`
	Description *string `json:"description"`
}

type ReportParameter struct {
	ID            int     `json:"id"`
	ParameterName *string `json:"parameterName"`
	Description   *string `json:"description"`
	DataType      *int    `json:"dataType"`
	Required      bool    `json:"required"`
	DefaultValue  *string `json:"defaultValue"`
	MultiSelect   bool    `json:"multiSelect"`
	SequenceNumber *int   `json:"sequenceNumber"`
	Inactive      bool    `json:"inactive"`
}

// Report is the clean domain representation of a Tessitura report definition.
type Report struct {
	ID                       string             `json:"id"`
	Name                     *string            `json:"name"`
	Description              *string            `json:"description"`
	ReportPath               *string            `json:"reportPath"`
	Category                 *ReportRef         `json:"category"`
	ReportType               *ReportRef         `json:"reportType"`
	AllowSchedule            *bool              `json:"allowSchedule"`
	AllowQuery               *bool              `json:"allowQuery"`
	QueryStringAppend        *string            `json:"queryStringAppend"`
	ParameterWindow          *string            `json:"parameterWindow"`
	ParameterWindowIndicator *bool              `json:"parameterWindowIndicator"`
	PublicIndicator          *bool              `json:"publicIndicator"`
	WarningIndicator         *bool              `json:"warningIndicator"`
	UtilityIndicator         *bool              `json:"utilityIndicator"`
	Window                   *string            `json:"window"`
	ApplicationID            *string            `json:"applicationId"`
	Inactive                 bool               `json:"inactive"`
	LastRequestID            *int               `json:"lastRequestId"`
	Parameters               []*ReportParameter `json:"parameters,omitempty"`
	CreatedAt                *string            `json:"createdAt"`
	CreatedBy                *string            `json:"createdBy"`
	UpdatedAt                *string            `json:"updatedAt"`
	UpdatedBy                *string            `json:"updatedBy"`
}

// ReportFromAPI maps a raw Tessitura API report response to the domain type.
func ReportFromAPI(r *tessitura.APIReport) *Report {
	report := &Report{
		AllowSchedule:            r.AllowSchedule,
		AllowQuery:               r.AllowQuery,
		QueryStringAppend:        r.QueryStringAppend,
		ParameterWindow:          r.ParameterWindow,
		ParameterWindowIndicator: r.ParameterWindowIndicator,
		PublicIndicator:          r.PublicIndicator,
		WarningIndicator:         r.WarningIndicator,
		UtilityIndicator:         r.UtilityIndicator,
		Window:                   r.Window,
		ApplicationID:            r.ApplicationId,
		LastRequestID:            r.LastRequestId,
		CreatedAt:                r.CreatedDateTime,
		CreatedBy:                r.CreatedBy,
		UpdatedAt:                r.UpdatedDateTime,
		UpdatedBy:                r.UpdatedBy,
		Name:                     r.Name,
		Description:              r.Description,
		ReportPath:               r.ReportPath,
	}

	if r.Id != nil {
		report.ID = *r.Id
	}
	if r.Inactive != nil {
		report.Inactive = *r.Inactive
	}
	if r.Category != nil {
		ref := &ReportRef{Description: r.Category.Description}
		if r.Category.Id != nil {
			ref.ID = *r.Category.Id
		}
		report.Category = ref
	}
	if r.ReportType != nil {
		ref := &ReportRef{Description: r.ReportType.Description}
		if r.ReportType.Id != nil {
			ref.ID = *r.ReportType.Id
		}
		report.ReportType = ref
	}

	return report
}

// AttachDetail merges fields from a ReportDetail response into the domain Report.
// This adds Parameters and covers fields absent from the base Report endpoint.
func (r *Report) AttachDetail(d *tessitura.APIReportDetail) {
	if d == nil {
		return
	}

	for _, p := range d.Parameters {
		if p == nil {
			continue
		}
		param := &ReportParameter{
			ParameterName: p.ParameterName,
			Description:   p.Description,
			DataType:      p.DataType,
			DefaultValue:  p.DefaultValue,
			SequenceNumber: p.SequenceNumber,
		}
		if p.Id != nil {
			param.ID = *p.Id
		}
		if p.Required != nil {
			param.Required = *p.Required
		}
		if p.MultiSelect != nil {
			param.MultiSelect = *p.MultiSelect
		}
		if p.Inactive != nil {
			param.Inactive = *p.Inactive
		}
		r.Parameters = append(r.Parameters, param)
	}
}
