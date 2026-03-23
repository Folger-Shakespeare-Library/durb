package tessitura

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
)

type APIConstituentSummary struct {
	Id              *int    `json:"Id"`
	DisplayName     *string `json:"DisplayName"`
	FirstName       *string `json:"FirstName"`
	MiddleName      *string `json:"MiddleName"`
	LastName        *string `json:"LastName"`
	Prefix          *string `json:"Prefix"`
	Suffix          *string `json:"Suffix"`
	Street1         *string `json:"Street1"`
	Street2         *string `json:"Street2"`
	City            *string `json:"City"`
	State           *string `json:"State"`
	PostalCode      *string `json:"PostalCode"`
	Country         *string `json:"Country"`
	Phone           *string `json:"Phone"`
	Email           *string `json:"Email"`
	ConstituentType *string `json:"ConstituentType"`
	TypeDescription *string `json:"TypeDescription"`
	Inactive        *string `json:"Inactive"`
}

type APISearchMetadata struct {
	Total *int `json:"Total"`
}

type APIConstituentSearchResponse struct {
	ConstituentSummaries []APIConstituentSummary `json:"ConstituentSummaries"`
	SearchMetadata       *APISearchMetadata      `json:"SearchMetadata"`
}

// SearchParams holds the query parameters for constituent search.
type SearchParams struct {
	// Free-text search (uses fluent type).
	Query string

	// Structured fields (uses basic type).
	LastName   string
	FirstName  string
	Street     string
	PostalCode string
	ID         string

	// Advanced search fields (each uses type=advanced with its own atype).
	Email              string
	Phone              string
	OrderNo            string
	WebLogin           string
	CustomerServiceNo  string

	// Filters (apply to all search types).
	IncludeAffiliates bool
	ConstituentGroups string

	Page     int
	PageSize int
}

// SearchConstituents performs a constituent search. Uses "basic" type for
// structured field searches and "fluent" type for free-text queries.
func (c *Client) SearchConstituents(ctx context.Context, params SearchParams) (*APIConstituentSearchResponse, error) {
	v := url.Values{}

	if atype, value := params.advancedSearch(); atype != "" {
		v.Set("type", "advanced")
		v.Set("atype", atype)
		v.Set("op", "Like")
		v.Set("value", value)
	} else if params.isBasic() {
		v.Set("type", "basic")
		if params.LastName != "" {
			v.Set("ln", params.LastName)
		}
		if params.FirstName != "" {
			v.Set("fn", params.FirstName)
		}
		if params.Street != "" {
			v.Set("street", params.Street)
		}
		if params.PostalCode != "" {
			v.Set("post", params.PostalCode)
		}
		if params.ID != "" {
			v.Set("constituentId", params.ID)
		}
	} else {
		v.Set("type", "fluent")
		v.Set("q", params.Query)
	}

	if params.IncludeAffiliates {
		v.Set("includeAffiliates", "true")
	} else {
		v.Set("includeAffiliates", "false")
	}

	if params.ConstituentGroups != "" {
		v.Set("constituentGroups", params.ConstituentGroups)
	}

	if params.Page > 0 {
		v.Set("page", fmt.Sprintf("%d", params.Page))
	}
	if params.PageSize > 0 {
		v.Set("pageSize", fmt.Sprintf("%d", params.PageSize))
	}

	data, err := c.Get(ctx, "/CRM/Constituents/Search?"+v.Encode())
	if err != nil {
		return nil, err
	}

	var resp APIConstituentSearchResponse
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (p SearchParams) isBasic() bool {
	return p.LastName != "" || p.FirstName != "" || p.Street != "" ||
		p.PostalCode != "" || p.ID != ""
}

// advancedSearch returns the atype and value if an advanced search field is set.
func (p SearchParams) advancedSearch() (atype, value string) {
	switch {
	case p.Email != "":
		return "Email", p.Email
	case p.Phone != "":
		return "Phone", p.Phone
	case p.OrderNo != "":
		return "Order No", p.OrderNo
	case p.WebLogin != "":
		return "Web Login", p.WebLogin
	case p.CustomerServiceNo != "":
		return "Customer Service No", p.CustomerServiceNo
	default:
		return "", ""
	}
}

func (p SearchParams) IsAdvanced() bool {
	atype, _ := p.advancedSearch()
	return atype != ""
}
