package domain

import "github.com/Folger-Shakespeare-Library/durb/pkg/tessitura"

// ConstituentSearchResult is a summary record returned from a search.
type ConstituentSearchResult struct {
	ID              int     `json:"id"`
	DisplayName     *string `json:"displayName"`
	FirstName       *string `json:"firstName"`
	LastName        *string `json:"lastName"`
	Prefix          *string `json:"prefix"`
	Suffix          *string `json:"suffix"`
	ConstituentType *string `json:"constituentType"`
	Inactive        bool    `json:"inactive"`
	Street1         *string `json:"street1,omitempty"`
	City            *string `json:"city,omitempty"`
	State           *string `json:"state,omitempty"`
	PostalCode      *string `json:"postalCode,omitempty"`
	Country         *string `json:"country,omitempty"`
	Phone           *string `json:"phone,omitempty"`
	Email           *string `json:"email,omitempty"`
}

func SearchResultsFromAPI(resp *tessitura.APIConstituentSearchResponse) []ConstituentSearchResult {
	var results []ConstituentSearchResult
	seen := make(map[int]bool)
	for _, s := range resp.ConstituentSummaries {
		id := derefInt(s.Id)
		if seen[id] {
			continue
		}
		seen[id] = true
		results = append(results, searchResultFromAPI(s))
	}
	if results == nil {
		results = []ConstituentSearchResult{}
	}
	return results
}

func searchResultFromAPI(s tessitura.APIConstituentSummary) ConstituentSearchResult {
	return ConstituentSearchResult{
		ID:              derefInt(s.Id),
		DisplayName:     s.DisplayName,
		FirstName:       s.FirstName,
		LastName:        s.LastName,
		Prefix:          s.Prefix,
		Suffix:          s.Suffix,
		ConstituentType: s.TypeDescription,
		Inactive:        s.Inactive != nil && *s.Inactive != "" && *s.Inactive != "0",
		Street1:         s.Street1,
		City:            s.City,
		State:           s.State,
		PostalCode:      s.PostalCode,
		Country:         s.Country,
		Phone:           s.Phone,
		Email:           s.Email,
	}
}
