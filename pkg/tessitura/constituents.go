package tessitura

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
)

// Raw API response structs — these mirror the Tessitura JSON shape exactly.
// They exist only to unmarshal the API response; all consumer code should
// use the domain types in pkg/domain instead.

type APIRef struct {
	Id          *int    `json:"Id"`
	Description *string `json:"Description"`
}

type APIInactiveSummary struct {
	Id          *int    `json:"Id"`
	Description *string `json:"Description"`
	Inactive    *bool   `json:"Inactive"`
}

type APIConstituentTypeSummary struct {
	Id               *int    `json:"Id"`
	Description      *string `json:"Description"`
	ConstituentGroup *APIRef `json:"ConstituentGroup"`
}

type APIStateSummary struct {
	Id          *int    `json:"Id"`
	Description *string `json:"Description"`
	StateCode   *string `json:"StateCode"`
}

type APICountrySummary struct {
	Id          *int    `json:"Id"`
	Description *string `json:"Description"`
	IsoAlpha2   *string `json:"IsoAlpha2Code"`
}

type APIAddress struct {
	Id               *int              `json:"Id"`
	Constituent      *APIEntity        `json:"Constituent"`
	AddressType      *APIRef           `json:"AddressType"`
	Street1          *string           `json:"Street1"`
	Street2          *string           `json:"Street2"`
	Street3          *string           `json:"Street3"`
	City             *string           `json:"City"`
	State            *APIStateSummary  `json:"State"`
	PostalCode       *string           `json:"PostalCode"`
	Country          *APICountrySummary `json:"Country"`
	PrimaryIndicator *bool             `json:"PrimaryIndicator"`
	Inactive         *bool             `json:"Inactive"`
	StartDate        *string           `json:"StartDate"`
	EndDate          *string           `json:"EndDate"`
	Months            *string           `json:"Months"`
	IsFromAffiliation bool              `json:"IsFromAffiliation"`
	CreatedDateTime   *string           `json:"CreatedDateTime"`
	UpdatedDateTime   *string           `json:"UpdatedDateTime"`
}

type APIElectronicAddress struct {
	Id                    *int       `json:"Id"`
	Constituent           *APIEntity `json:"Constituent"`
	ElectronicAddressType *APIRef    `json:"ElectronicAddressType"`
	Address               *string `json:"Address"`
	PrimaryIndicator      *bool   `json:"PrimaryIndicator"`
	Inactive              *bool   `json:"Inactive"`
	IsEmail               *bool   `json:"IsEmail"`
	AllowHtmlFormat       *bool   `json:"AllowHtmlFormat"`
	AllowMarketing        *bool   `json:"AllowMarketing"`
	StartDate             *string `json:"StartDate"`
	EndDate               *string `json:"EndDate"`
	Months                *string `json:"Months"`
	IsFromAffiliation     bool    `json:"IsFromAffiliation"`
	CreatedDateTime       *string `json:"CreatedDateTime"`
	UpdatedDateTime       *string `json:"UpdatedDateTime"`
}

type APIPhone struct {
	Id                     *int              `json:"Id"`
	Constituent            *APIEntity        `json:"Constituent"`
	PhoneType              *APIRef           `json:"PhoneType"`
	PhoneNumber            *string           `json:"PhoneNumber"`
	PhoneFormatted         *string           `json:"PhoneFormatted"`
	InternationalMobile    *string           `json:"InternationalMobilePhone"`
	Country                *APICountrySummary `json:"Country"`
	PrimaryIndicator       *bool             `json:"PrimaryIndicator"`
	Inactive               *bool             `json:"Inactive"`
	IsMobile               *bool             `json:"IsMobile"`
	AllowTelemarketing     *bool             `json:"AllowTelemarketing"`
	CreatedDateTime        *string           `json:"CreatedDateTime"`
	UpdatedDateTime        *string           `json:"UpdatedDateTime"`
}

type APISalutation struct {
	Id                  *int       `json:"Id"`
	Constituent         *APIEntity `json:"Constituent"`
	SalutationType      *APIRef    `json:"SalutationType"`
	BusinessTitle       *string    `json:"BusinessTitle"`
	EnvelopeSalutation1 *string    `json:"EnvelopeSalutation1"`
	EnvelopeSalutation2 *string    `json:"EnvelopeSalutation2"`
	LetterSalutation    *string    `json:"LetterSalutation"`
	DefaultIndicator    bool       `json:"DefaultIndicator"`
	CreatedDateTime     *string    `json:"CreatedDateTime"`
	UpdatedDateTime     *string    `json:"UpdatedDateTime"`
}

type APIConstituentDetail struct {
	Id                   *int                          `json:"Id"`
	FirstName            *string                       `json:"FirstName"`
	MiddleName           *string                       `json:"MiddleName"`
	LastName             *string                       `json:"LastName"`
	DisplayName          *string                       `json:"DisplayName"`
	SortName             *string                       `json:"SortName"`
	Prefix               *APIRef                       `json:"Prefix"`
	Suffix               *APIRef                       `json:"Suffix"`
	Gender               *APIRef                       `json:"Gender"`
	Pronoun              *APIRef                       `json:"Pronoun"`
	ConstituentType      *APIConstituentTypeSummary    `json:"ConstituentType"`
	Inactive             *APIInactiveSummary            `json:"Inactive"`
	InactiveReason       *APIRef                       `json:"InactiveReason"`
	ProtectionType       *APIRef                       `json:"ProtectionType"`
	OriginalSource       *APIRef                       `json:"OriginalSource"`
	NameStatus           *APIRef                       `json:"NameStatus"`
	MailIndicator        *APIRef                       `json:"MailIndicator"`
	EmarketIndicator     *APIRef                       `json:"EmarketIndicator"`
	PhoneIndicator       *APIRef                       `json:"PhoneIndicator"`
	LastActivityDate     *string                       `json:"LastActivityDate"`
	LastGiftDate         *string                       `json:"LastGiftDate"`
	LastTicketDate       *string                       `json:"LastTicketDate"`
	CreatedDateTime      *string                       `json:"CreatedDateTime"`
	CreatedBy            *string                       `json:"CreatedBy"`
	UpdatedDateTime      *string                       `json:"UpdatedDateTime"`
	UpdatedBy            *string                       `json:"UpdatedBy"`
	Addresses            []APIAddress                  `json:"Addresses"`
	ElectronicAddresses  []APIElectronicAddress        `json:"ElectronicAddresses"`
	PhoneNumbers         []APIPhone                    `json:"PhoneNumbers"`
	Salutations          []APISalutation               `json:"Salutations"`
}

// GetConstituentDetail fetches the full detail view for a constituent.
// This includes addresses, electronic addresses, phones, and salutations.
func (c *Client) GetConstituentDetail(ctx context.Context, id string) (*APIConstituentDetail, error) {
	data, err := c.Get(ctx, "/CRM/Constituents/"+id+"/Detail")
	if err != nil {
		return nil, err
	}

	var detail APIConstituentDetail
	if err := json.Unmarshal(data, &detail); err != nil {
		return nil, err
	}
	return &detail, nil
}

// ConstituentResult holds all API data fetched for a single constituent.
type ConstituentResult struct {
	Detail       *APIConstituentDetail
	Affiliations []APIAffiliation
	Associations []APIAssociation
	Notes        []APINote
}

// GetConstituentFull fetches detail and optionally affiliations, associations,
// and/or notes for a single constituent. When any extras are requested, it uses
// a single batch call to fetch everything in one HTTP request.
func (c *Client) GetConstituentFull(ctx context.Context, id string, includeAffiliations, includeAssociations, includeNotes bool) (*ConstituentResult, error) {
	if !includeAffiliations && !includeAssociations && !includeNotes {
		// Simple case: just the detail endpoint
		detail, err := c.GetConstituentDetail(ctx, id)
		if err != nil {
			return nil, err
		}
		return &ConstituentResult{Detail: detail}, nil
	}

	// Batch the detail + any requested extras into one HTTP call.
	// Fixed request IDs: 1=detail, 2=affiliations (individual), 3=affiliations (group), 4=notes, 5=associations
	items := []BatchRequestItem{
		{HttpMethod: "GET", Id: 1, Uri: "/CRM/Constituents/" + id + "/Detail"},
	}
	if includeAffiliations {
		items = append(items, BatchRequestItem{HttpMethod: "GET", Id: 2, Uri: "/CRM/Affiliations?individualConstituentId=" + id})
		items = append(items, BatchRequestItem{HttpMethod: "GET", Id: 3, Uri: "/CRM/Affiliations?groupConstituentId=" + id})
	}
	if includeNotes {
		items = append(items, BatchRequestItem{HttpMethod: "GET", Id: 4, Uri: "/CRM/Notes?constituentId=" + id})
	}
	if includeAssociations {
		items = append(items, BatchRequestItem{HttpMethod: "GET", Id: 5, Uri: "/CRM/Associations?constituentId=" + id})
	}

	batchResp, err := c.Batch(ctx, items)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch constituent %s: %w", id, err)
	}

	result := &ConstituentResult{}

	for _, resp := range batchResp.Responses {
		if resp.StatusCode < 200 || resp.StatusCode >= 300 {
			return nil, fmt.Errorf("request %d for constituent %s failed with status %d",
				resp.RequestId, id, resp.StatusCode)
		}

		switch resp.RequestId {
		case 1:
			var detail APIConstituentDetail
			if err := json.Unmarshal(resp.ResponseObject, &detail); err != nil {
				return nil, fmt.Errorf("unable to parse detail for constituent %s: %w", id, err)
			}
			result.Detail = &detail

		case 2:
			var affiliations []APIAffiliation
			if err := json.Unmarshal(resp.ResponseObject, &affiliations); err != nil {
				return nil, fmt.Errorf("unable to parse affiliations for constituent %s: %w", id, err)
			}
			result.Affiliations = append(result.Affiliations, affiliations...)

		case 3:
			var affiliations []APIAffiliation
			if err := json.Unmarshal(resp.ResponseObject, &affiliations); err != nil {
				return nil, fmt.Errorf("unable to parse group affiliations for constituent %s: %w", id, err)
			}
			result.Affiliations = append(result.Affiliations, affiliations...)

		case 4:
			var notes []APINote
			if err := json.Unmarshal(resp.ResponseObject, &notes); err != nil {
				return nil, fmt.Errorf("unable to parse notes for constituent %s: %w", id, err)
			}
			result.Notes = notes

		case 5:
			var associations []APIAssociation
			if err := json.Unmarshal(resp.ResponseObject, &associations); err != nil {
				return nil, fmt.Errorf("unable to parse associations for constituent %s: %w", id, err)
			}
			result.Associations = associations
		}
	}

	if result.Detail == nil {
		return nil, fmt.Errorf("no detail returned for constituent %s", id)
	}

	return result, nil
}

// GetConstituentsBatch fetches multiple constituents concurrently using
// goroutines. Each constituent gets its own HTTP call (or batch call if
// extras are included). Results are returned in the same order as the input IDs.
func (c *Client) GetConstituentsBatch(ctx context.Context, ids []string, includeAffiliations, includeAssociations, includeNotes bool) ([]*ConstituentResult, error) {
	results := make([]*ConstituentResult, len(ids))
	errs := make([]error, len(ids))

	var wg sync.WaitGroup
	for i, id := range ids {
		wg.Add(1)
		go func(idx int, cid string) {
			defer wg.Done()
			result, err := c.GetConstituentFull(ctx, cid, includeAffiliations, includeAssociations, includeNotes)
			results[idx] = result
			errs[idx] = err
		}(i, id)
	}
	wg.Wait()

	// Return the first error encountered
	for i, err := range errs {
		if err != nil {
			return nil, fmt.Errorf("constituent %s: %w", ids[i], err)
		}
	}

	return results, nil
}
