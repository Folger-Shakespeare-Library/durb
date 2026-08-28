package tessitura

import (
	"context"
	"encoding/json"
	"fmt"
)

type APIRefItem struct {
	Id          *int    `json:"Id"`
	Description *string `json:"Description"`
	Inactive    *bool   `json:"Inactive"`
}

func (c *Client) GetConstituentInactiveStatuses(ctx context.Context) ([]APIRefItem, error) {
	data, err := c.Get(ctx, "/ReferenceData/ConstituentInactives/Summary")
	if err != nil {
		return nil, fmt.Errorf("fetching inactive statuses: %w", err)
	}
	var items []APIRefItem
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("parsing inactive statuses: %w", err)
	}
	return items, nil
}

type APISeatStatus struct {
	Id             *int    `json:"Id"`
	Description    *string `json:"Description"`
	Inactive       bool    `json:"Inactive"`
	StatusCode     *string `json:"StatusCode"`
	StatusLegend   *string `json:"StatusLegend"`
	StatusPriority int     `json:"StatusPriority"`
}

func (c *Client) GetSeatStatuses(ctx context.Context) ([]APISeatStatus, error) {
	data, err := c.Get(ctx, "/ReferenceData/SeatStatuses")
	if err != nil {
		return nil, fmt.Errorf("fetching seat statuses: %w", err)
	}
	var items []APISeatStatus
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("parsing seat statuses: %w", err)
	}
	return items, nil
}

type APICardReaderTypeSummary struct {
	Id          *int    `json:"Id"`
	Description *string `json:"Description"`
}

type APIMachineSetting struct {
	Id                                    *int                     `json:"Id"`
	WorkstationName                       *string                  `json:"WorkstationName"`
	Inactive                              bool                     `json:"Inactive"`
	CardReaderHost                        *string                  `json:"CardReaderHost"`
	CardReaderPort                        *int                     `json:"CardReaderPort"`
	CardReaderType                        *APICardReaderTypeSummary `json:"CardReaderType"`
	PXStation                             *string                  `json:"PXStation"`
	MerchantId                            *string                  `json:"MerchantId"`
	PXUserName                            *string                  `json:"PXUserName"`
	PXUserKey                             *string                  `json:"PXUserKey"`
	TnspaySoftwareTerminal                *bool                    `json:"TnspaySoftwareTerminal"`
	TriposLane                            *int                     `json:"TriposLane"`
	TessituraMerchantServicesPosDevice      *string                `json:"TessituraMerchantServicesPosDevice"`
	TessituraMerchantServicesPosDeviceModel *string                `json:"TessituraMerchantServicesPosDeviceModel"`
	CreateLocation                        *string                  `json:"CreateLocation"`
	CreatedBy                             *string                  `json:"CreatedBy"`
	CreatedDateTime                       *string                  `json:"CreatedDateTime"`
	UpdatedBy                             *string                  `json:"UpdatedBy"`
	UpdatedDateTime                       *string                  `json:"UpdatedDateTime"`
}

func (c *Client) GetMachineSettings(ctx context.Context) ([]APIMachineSetting, error) {
	data, err := c.Get(ctx, "/ReferenceData/MachineSettings")
	if err != nil {
		return nil, fmt.Errorf("fetching machine settings: %w", err)
	}
	var items []APIMachineSetting
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("parsing machine settings: %w", err)
	}
	return items, nil
}

type APIConstituentType struct {
	Id                               *int         `json:"Id"`
	Description                      *string      `json:"Description"`
	Inactive                         bool         `json:"Inactive"`
	ConstituentGroup                 *APIRefItem  `json:"ConstituentGroup"`
	DefaultIndicator                 *bool        `json:"DefaultIndicator"`
	DefaultSalutationId              *int         `json:"DefaultSalutationId"`
	DefaultAffiliatedConstituentTypeId *int       `json:"DefaultAffiliatedConstituentTypeId"`
	DefaultAffiliationTypeId         *int         `json:"DefaultAffiliationTypeId"`
	AddressTypeId                    *int         `json:"AddressTypeId"`
	ElectronicAddressTypeId          *int         `json:"ElectronicAddressTypeId"`
	PhoneTypeId                      *int         `json:"PhoneTypeId"`
	LoginTypeId                      *int         `json:"LoginTypeId"`
	GiftAidIndicator                 bool         `json:"GiftAidIndicator"`
	CreatedBy                        *string      `json:"CreatedBy"`
	CreatedDateTime                  *string      `json:"CreatedDateTime"`
	CreateLocation                   *string      `json:"CreateLocation"`
	UpdatedBy                        *string      `json:"UpdatedBy"`
	UpdatedDateTime                  *string      `json:"UpdatedDateTime"`
}

func (c *Client) GetConstituentTypes(ctx context.Context) ([]APIConstituentType, error) {
	data, err := c.Get(ctx, "/ReferenceData/ConstituentTypes")
	if err != nil {
		return nil, fmt.Errorf("fetching constituent types: %w", err)
	}
	var items []APIConstituentType
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("parsing constituent types: %w", err)
	}
	return items, nil
}

type APIOriginalSource struct {
	Id              *int    `json:"Id"`
	Description     *string `json:"Description"`
	Inactive        *bool   `json:"Inactive"`
	CreatedBy       *string `json:"CreatedBy"`
	CreatedDateTime *string `json:"CreatedDateTime"`
	CreateLocation  *string `json:"CreateLocation"`
	UpdatedBy       *string `json:"UpdatedBy"`
	UpdatedDateTime *string `json:"UpdatedDateTime"`
}

func (c *Client) GetOriginalSources(ctx context.Context) ([]APIOriginalSource, error) {
	data, err := c.Get(ctx, "/ReferenceData/OriginalSources")
	if err != nil {
		return nil, fmt.Errorf("fetching original sources: %w", err)
	}
	var items []APIOriginalSource
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("parsing original sources: %w", err)
	}
	return items, nil
}

type APISpecialActivityType struct {
	Id              *int        `json:"Id"`
	Description     *string     `json:"Description"`
	Inactive        bool        `json:"Inactive"`
	ControlGroup    *APIRefItem `json:"ControlGroup"`
	CreatedBy       *string     `json:"CreatedBy"`
	CreatedDateTime *string     `json:"CreatedDateTime"`
	CreateLocation  *string     `json:"CreateLocation"`
	UpdatedBy       *string     `json:"UpdatedBy"`
	UpdatedDateTime *string     `json:"UpdatedDateTime"`
}

func (c *Client) GetSpecialActivityTypes(ctx context.Context) ([]APISpecialActivityType, error) {
	data, err := c.Get(ctx, "/ReferenceData/SpecialActivityTypes")
	if err != nil {
		return nil, fmt.Errorf("fetching special activity types: %w", err)
	}
	var items []APISpecialActivityType
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("parsing special activity types: %w", err)
	}
	return items, nil
}

type APISpecialActivityStatus struct {
	Id              *int    `json:"Id"`
	Description     *string `json:"Description"`
	Inactive        bool    `json:"Inactive"`
	CreatedBy       *string `json:"CreatedBy"`
	CreatedDateTime *string `json:"CreatedDateTime"`
	CreateLocation  *string `json:"CreateLocation"`
	UpdatedBy       *string `json:"UpdatedBy"`
	UpdatedDateTime *string `json:"UpdatedDateTime"`
}

func (c *Client) GetSpecialActivityStatuses(ctx context.Context) ([]APISpecialActivityStatus, error) {
	data, err := c.Get(ctx, "/ReferenceData/SpecialActivityStatuses")
	if err != nil {
		return nil, fmt.Errorf("fetching special activity statuses: %w", err)
	}
	var items []APISpecialActivityStatus
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("parsing special activity statuses: %w", err)
	}
	return items, nil
}

type APIKeyword struct {
	Id                         *int        `json:"Id"`
	Description                *string     `json:"Description"`
	Category                   *APIRefItem `json:"Category"`
	ControlGroup               *APIRefItem `json:"ControlGroup"`
	ConstituentType            *int        `json:"ConstituentType"`
	DataType                   *string     `json:"DataType"`
	KeywordUse                 *string     `json:"KeywordUse"`
	MultipleValue              *bool       `json:"MultipleValue"`
	CustomRequired             *bool       `json:"CustomRequired"`
	CustomDefaultValue         *string     `json:"CustomDefaultValue"`
	CustomLimit                *int        `json:"CustomLimit"`
	CustomId                   *int        `json:"CustomId"`
	EditMask                   *string     `json:"EditMask"`
	ExtendedDescription        *string     `json:"ExtendedDescription"`
	HelpText                   *string     `json:"HelpText"`
	UseForSearch               *bool       `json:"UseForSearch"`
	SortOrder                  *int        `json:"SortOrder"`
	ValuesCodedIndicator       *bool       `json:"ValuesCodedIndicator"`
	PrimaryGroupDefault        *string     `json:"PrimaryGroupDefault"`
	EditIndicator              *bool       `json:"EditIndicator"`
	FrequentUpdateDate         *string     `json:"FrequentUpdateDate"`
	DetailTable                *string     `json:"DetailTable"`
	DetailColumn               *string     `json:"DetailColumn"`
	ParentTable                *string     `json:"ParentTable"`
	ParentKeyColumn            *string     `json:"ParentKeyColumn"`
	KeyColumn                  *string     `json:"KeyColumn"`
	ReferenceTable             *string     `json:"ReferenceTable"`
	ReferenceIdColumn          *string     `json:"ReferenceIdColumn"`
	ReferenceDescriptionColumn *string     `json:"ReferenceDescriptionColumn"`
	ReferenceSort              *string     `json:"ReferenceSort"`
	ReferenceWhere             *string     `json:"ReferenceWhere"`
	CreatedBy                  *string     `json:"CreatedBy"`
	CreatedDateTime            *string     `json:"CreatedDateTime"`
	CreateLocation             *string     `json:"CreateLocation"`
	UpdatedBy                  *string     `json:"UpdatedBy"`
	UpdatedDateTime            *string     `json:"UpdatedDateTime"`
}

func (c *Client) GetKeywords(ctx context.Context) ([]APIKeyword, error) {
	data, err := c.Get(ctx, "/ReferenceData/Keywords")
	if err != nil {
		return nil, fmt.Errorf("fetching keywords: %w", err)
	}
	var items []APIKeyword
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("parsing keywords: %w", err)
	}
	return items, nil
}

type APIInterestType struct {
	Id              *int        `json:"Id"`
	Description     *string     `json:"Description"`
	Category        *APIRefItem `json:"Category"`
	ControlGroup    *APIRefItem `json:"ControlGroup"`
	UsedIn          *string     `json:"UsedIn"`
	SearchIndicator bool        `json:"SearchIndicator"`
	CreatedBy       *string     `json:"CreatedBy"`
	CreatedDateTime *string     `json:"CreatedDateTime"`
	CreateLocation  *string     `json:"CreateLocation"`
	UpdatedBy       *string     `json:"UpdatedBy"`
	UpdatedDateTime *string     `json:"UpdatedDateTime"`
}

func (c *Client) GetInterestTypes(ctx context.Context) ([]APIInterestType, error) {
	data, err := c.Get(ctx, "/ReferenceData/InterestTypes")
	if err != nil {
		return nil, fmt.Errorf("fetching interest types: %w", err)
	}
	var items []APIInterestType
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("parsing interest types: %w", err)
	}
	return items, nil
}

type APIInactiveReason struct {
	Id              *int    `json:"Id"`
	Description     *string `json:"Description"`
	Inactive        bool    `json:"Inactive"`
	CreatedBy       *string `json:"CreatedBy"`
	CreatedDateTime *string `json:"CreatedDateTime"`
	CreateLocation  *string `json:"CreateLocation"`
	UpdatedBy       *string `json:"UpdatedBy"`
	UpdatedDateTime *string `json:"UpdatedDateTime"`
}

func (c *Client) GetInactiveReasons(ctx context.Context) ([]APIInactiveReason, error) {
	data, err := c.Get(ctx, "/ReferenceData/InactiveReasons")
	if err != nil {
		return nil, fmt.Errorf("fetching inactive reasons: %w", err)
	}
	var items []APIInactiveReason
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("parsing inactive reasons: %w", err)
	}
	return items, nil
}

func (c *Client) GetConstituentInactiveReasons(ctx context.Context) ([]APIRefItem, error) {
	data, err := c.Get(ctx, "/ReferenceData/InactiveReasons")
	if err != nil {
		return nil, fmt.Errorf("fetching inactive reasons: %w", err)
	}
	var items []APIRefItem
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("parsing inactive reasons: %w", err)
	}
	return items, nil
}
