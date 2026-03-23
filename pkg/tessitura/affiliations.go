package tessitura

import (
	"context"
	"encoding/json"
)

type APIAffiliationTypeSummary struct {
	Id          *int    `json:"Id"`
	Description *string `json:"Description"`
	Inactive    *bool   `json:"Inactive"`
	UseTitle    *bool   `json:"UseTitle"`
	UseSalary   *bool   `json:"UseSalary"`
}

type APIEntity struct {
	Id int `json:"Id"`
}

type APIAffiliation struct {
	Id                          *int                       `json:"Id"`
	AffiliationType             *APIAffiliationTypeSummary `json:"AffiliationType"`
	IndividualConstituent       *APIEntity                 `json:"IndividualConstituent"`
	GroupConstituent            *APIEntity                 `json:"GroupConstituent"`
	IndividualConstituentName   *string                    `json:"IndividualConstituentName"`
	GroupConstituentName        *string                    `json:"GroupConstituentName"`
	Title                       *string                    `json:"Title"`
	Note                        *string                    `json:"Note"`
	PrimaryIndicator            bool                       `json:"PrimaryIndicator"`
	Inactive                    *bool                      `json:"Inactive"`
	StartDate                   *string                    `json:"StartDate"`
	EndDate                     *string                    `json:"EndDate"`
	CreatedDateTime             *string                    `json:"CreatedDateTime"`
	UpdatedDateTime             *string                    `json:"UpdatedDateTime"`
}

// GetAffiliations fetches all affiliations for a constituent, checking both
// the individual side (person → org) and group side (org → person).
func (c *Client) GetAffiliations(ctx context.Context, constituentID string) ([]APIAffiliation, error) {
	// Try as individual first (most common for person records).
	data, err := c.Get(ctx, "/CRM/Affiliations?individualConstituentId="+constituentID)
	if err != nil {
		return nil, err
	}

	var affiliations []APIAffiliation
	if err := json.Unmarshal(data, &affiliations); err != nil {
		return nil, err
	}

	// Also check as group constituent (for orgs/households).
	data, err = c.Get(ctx, "/CRM/Affiliations?groupConstituentId="+constituentID)
	if err != nil {
		return nil, err
	}

	var groupAffiliations []APIAffiliation
	if err := json.Unmarshal(data, &groupAffiliations); err != nil {
		return nil, err
	}

	affiliations = append(affiliations, groupAffiliations...)
	return affiliations, nil
}
