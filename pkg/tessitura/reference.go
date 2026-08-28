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

type APIReportCategory struct {
	Id              int     `json:"Id"`
	Description     *string `json:"Description"`
	DescriptionText *string `json:"DescriptionText"`
	Inactive        *bool   `json:"Inactive"`
	CreatedBy       *string `json:"CreatedBy"`
	CreatedDateTime *string `json:"CreatedDateTime"`
	CreateLocation  *string `json:"CreateLocation"`
	UpdatedBy       *string `json:"UpdatedBy"`
	UpdatedDateTime *string `json:"UpdatedDateTime"`
}

func (c *Client) GetReportCategories(ctx context.Context) ([]APIReportCategory, error) {
	data, err := c.Get(ctx, "/ReferenceData/ReportCategories")
	if err != nil {
		return nil, fmt.Errorf("fetching report categories: %w", err)
	}
	var items []APIReportCategory
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("parsing report categories: %w", err)
	}
	return items, nil
}

type APIReportType struct {
	Id              int     `json:"Id"`
	Description     *string `json:"Description"`
	HelpText        *string `json:"HelpText"`
	Inactive        bool    `json:"Inactive"`
	CreatedBy       *string `json:"CreatedBy"`
	CreatedDateTime *string `json:"CreatedDateTime"`
	CreateLocation  *string `json:"CreateLocation"`
	UpdatedBy       *string `json:"UpdatedBy"`
	UpdatedDateTime *string `json:"UpdatedDateTime"`
}

func (c *Client) GetReportTypes(ctx context.Context) ([]APIReportType, error) {
	data, err := c.Get(ctx, "/ReferenceData/ReportTypes")
	if err != nil {
		return nil, fmt.Errorf("fetching report types: %w", err)
	}
	var items []APIReportType
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("parsing report types: %w", err)
	}
	return items, nil
}

type APISeason struct {
	Id                       *int                `json:"Id"`
	Description              *string             `json:"Description"`
	Inactive                 *bool               `json:"Inactive"`
	DefaultIndicator         *bool               `json:"DefaultIndicator"`
	DisplayInSeasonOverview  *bool               `json:"DisplayInSeasonOverview"`
	StartDateTime            *string             `json:"StartDateTime"`
	EndDateTime              *string             `json:"EndDateTime"`
	FYear                    *int                `json:"FYear"`
	YearlySeason             *int                `json:"YearlySeason"`
	ConfirmationNoticeFormat *int                `json:"ConfirmationNoticeFormat"`
	RenewalNoticeFormat      *int                `json:"RenewalNoticeFormat"`
	SubscriptionFund1        *int                `json:"SubscriptionFund1"`
	SubscriptionFund2        *int                `json:"SubscriptionFund2"`
	Type                     *APIRefItem         `json:"Type"`
	ControlGroup             *APIControlGroupRef `json:"ControlGroup"`
	CreatedBy                *string             `json:"CreatedBy"`
	CreatedDateTime          *string             `json:"CreatedDateTime"`
	CreateLocation           *string             `json:"CreateLocation"`
	UpdatedBy                *string             `json:"UpdatedBy"`
	UpdatedDateTime          *string             `json:"UpdatedDateTime"`
}

type APIControlGroupRef struct {
	Id          int     `json:"Id"`
	Description *string `json:"Description"`
	Inactive    bool    `json:"Inactive"`
}

func (c *Client) GetSeasons(ctx context.Context) ([]APISeason, error) {
	data, err := c.Get(ctx, "/ReferenceData/Seasons")
	if err != nil {
		return nil, fmt.Errorf("fetching seasons: %w", err)
	}
	var items []APISeason
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("parsing seasons: %w", err)
	}
	return items, nil
}

type APIPerformanceStatus struct {
	Id              *int    `json:"Id"`
	Description     *string `json:"Description"`
	Inactive        *bool   `json:"Inactive"`
	CreatedBy       *string `json:"CreatedBy"`
	CreatedDateTime *string `json:"CreatedDateTime"`
	CreateLocation  *string `json:"CreateLocation"`
	UpdatedBy       *string `json:"UpdatedBy"`
	UpdatedDateTime *string `json:"UpdatedDateTime"`
}

func (c *Client) GetPerformanceStatuses(ctx context.Context) ([]APIPerformanceStatus, error) {
	data, err := c.Get(ctx, "/ReferenceData/PerformanceStatuses")
	if err != nil {
		return nil, fmt.Errorf("fetching performance statuses: %w", err)
	}
	var items []APIPerformanceStatus
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("parsing performance statuses: %w", err)
	}
	return items, nil
}

type APIPerformanceType struct {
	Id               *int    `json:"Id"`
	Description      *string `json:"Description"`
	Inactive         *bool   `json:"Inactive"`
	ValidCountryList *string `json:"ValidCountryList"`
	CreatedBy        *string `json:"CreatedBy"`
	CreatedDateTime  *string `json:"CreatedDateTime"`
	CreateLocation   *string `json:"CreateLocation"`
	UpdatedBy        *string `json:"UpdatedBy"`
	UpdatedDateTime  *string `json:"UpdatedDateTime"`
}

func (c *Client) GetPerformanceTypes(ctx context.Context) ([]APIPerformanceType, error) {
	data, err := c.Get(ctx, "/ReferenceData/PerformanceTypes")
	if err != nil {
		return nil, fmt.Errorf("fetching performance types: %w", err)
	}
	var items []APIPerformanceType
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("parsing performance types: %w", err)
	}
	return items, nil
}

type APIPriceCategory struct {
	Id              *int    `json:"Id"`
	Description     *string `json:"Description"`
	Inactive        *bool   `json:"Inactive"`
	Rank            int     `json:"Rank"`
	CreatedBy       *string `json:"CreatedBy"`
	CreatedDateTime *string `json:"CreatedDateTime"`
	CreateLocation  *string `json:"CreateLocation"`
	UpdatedBy       *string `json:"UpdatedBy"`
	UpdatedDateTime *string `json:"UpdatedDateTime"`
}

func (c *Client) GetPriceCategories(ctx context.Context) ([]APIPriceCategory, error) {
	data, err := c.Get(ctx, "/ReferenceData/PriceCategories")
	if err != nil {
		return nil, fmt.Errorf("fetching price categories: %w", err)
	}
	var items []APIPriceCategory
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("parsing price categories: %w", err)
	}
	return items, nil
}

type APIPriceTypeCategory struct {
	Id               *int    `json:"Id"`
	Description      *string `json:"Description"`
	ShortDescription *string `json:"ShortDescription"`
	CreatedBy        *string `json:"CreatedBy"`
	CreatedDateTime  *string `json:"CreatedDateTime"`
	CreateLocation   *string `json:"CreateLocation"`
	UpdatedBy        *string `json:"UpdatedBy"`
	UpdatedDateTime  *string `json:"UpdatedDateTime"`
}

func (c *Client) GetPriceTypeCategories(ctx context.Context) ([]APIPriceTypeCategory, error) {
	data, err := c.Get(ctx, "/ReferenceData/PriceTypeCategories")
	if err != nil {
		return nil, fmt.Errorf("fetching price type categories: %w", err)
	}
	var items []APIPriceTypeCategory
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("parsing price type categories: %w", err)
	}
	return items, nil
}

type APIPriceTypeGroup struct {
	Id              *int                `json:"Id"`
	Description     *string             `json:"Description"`
	Inactive        *bool               `json:"Inactive"`
	ControlGroup    *APIControlGroupRef `json:"ControlGroup"`
	CreatedBy       *string             `json:"CreatedBy"`
	CreatedDateTime *string             `json:"CreatedDateTime"`
	CreateLocation  *string             `json:"CreateLocation"`
	UpdatedBy       *string             `json:"UpdatedBy"`
	UpdatedDateTime *string             `json:"UpdatedDateTime"`
}

func (c *Client) GetPriceTypeGroups(ctx context.Context) ([]APIPriceTypeGroup, error) {
	data, err := c.Get(ctx, "/ReferenceData/PriceTypeGroups")
	if err != nil {
		return nil, fmt.Errorf("fetching price type groups: %w", err)
	}
	var items []APIPriceTypeGroup
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("parsing price type groups: %w", err)
	}
	return items, nil
}

type APIPaymentType struct {
	Id                   int     `json:"Id"`
	Description          *string `json:"Description"`
	Inactive             bool    `json:"Inactive"`
	MerchantServicesType *string `json:"MerchantServicesType"`
	IsExtendedAuth       bool    `json:"IsExtendedAuth"`
	OnlineOnly           bool    `json:"OnlineOnly"`
	HasTokens            bool    `json:"HasTokens"`
	CreatedBy            *string `json:"CreatedBy"`
	CreatedDateTime      *string `json:"CreatedDateTime"`
	CreateLocation       *string `json:"CreateLocation"`
	UpdatedBy            *string `json:"UpdatedBy"`
	UpdatedDateTime      *string `json:"UpdatedDateTime"`
}

func (c *Client) GetPaymentTypes(ctx context.Context) ([]APIPaymentType, error) {
	data, err := c.Get(ctx, "/ReferenceData/PaymentTypes")
	if err != nil {
		return nil, fmt.Errorf("fetching payment types: %w", err)
	}
	var items []APIPaymentType
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("parsing payment types: %w", err)
	}
	return items, nil
}

type APIDeliveryMethod struct {
	Id              *int    `json:"Id"`
	Description     *string `json:"Description"`
	AliasDescription *string `json:"AliasDescription"`
	Inactive        bool    `json:"Inactive"`
	PrintAtHome     bool    `json:"PrintAtHome"`
	RequireAddress  bool    `json:"RequireAddress"`
	MobileIndicator bool    `json:"MobileIndicator"`
	CreatedBy       *string `json:"CreatedBy"`
	CreatedDateTime *string `json:"CreatedDateTime"`
	CreateLocation  *string `json:"CreateLocation"`
	UpdatedBy       *string `json:"UpdatedBy"`
	UpdatedDateTime *string `json:"UpdatedDateTime"`
}

func (c *Client) GetDeliveryMethods(ctx context.Context) ([]APIDeliveryMethod, error) {
	data, err := c.Get(ctx, "/ReferenceData/DeliveryMethods")
	if err != nil {
		return nil, fmt.Errorf("fetching delivery methods: %w", err)
	}
	var items []APIDeliveryMethod
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("parsing delivery methods: %w", err)
	}
	return items, nil
}

type APIOrderCategory struct {
	Id              *int    `json:"Id"`
	Description     *string `json:"Description"`
	Inactive        bool    `json:"Inactive"`
	CreatedBy       *string `json:"CreatedBy"`
	CreatedDateTime *string `json:"CreatedDateTime"`
	CreateLocation  *string `json:"CreateLocation"`
	UpdatedBy       *string `json:"UpdatedBy"`
	UpdatedDateTime *string `json:"UpdatedDateTime"`
}

func (c *Client) GetOrderCategories(ctx context.Context) ([]APIOrderCategory, error) {
	data, err := c.Get(ctx, "/ReferenceData/OrderCategories")
	if err != nil {
		return nil, fmt.Errorf("fetching order categories: %w", err)
	}
	var items []APIOrderCategory
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("parsing order categories: %w", err)
	}
	return items, nil
}

type APISalesChannel struct {
	Id              *int                `json:"Id"`
	Description     *string             `json:"Description"`
	Inactive        *bool               `json:"Inactive"`
	ControlGroup    *APIControlGroupRef `json:"ControlGroup"`
	CreatedBy       *string             `json:"CreatedBy"`
	CreatedDateTime *string             `json:"CreatedDateTime"`
	CreateLocation  *string             `json:"CreateLocation"`
	UpdatedBy       *string             `json:"UpdatedBy"`
	UpdatedDateTime *string             `json:"UpdatedDateTime"`
}

func (c *Client) GetSalesChannels(ctx context.Context) ([]APISalesChannel, error) {
	data, err := c.Get(ctx, "/ReferenceData/SalesChannels")
	if err != nil {
		return nil, fmt.Errorf("fetching sales channels: %w", err)
	}
	var items []APISalesChannel
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("parsing sales channels: %w", err)
	}
	return items, nil
}

type APINoteType struct {
	Id              *int                `json:"Id"`
	Description     *string             `json:"Description"`
	Inactive        bool                `json:"Inactive"`
	OkToPrint       *string             `json:"OkToPrint"`
	ControlGroup    *APIControlGroupRef `json:"ControlGroup"`
	CreatedBy       *string             `json:"CreatedBy"`
	CreatedDateTime *string             `json:"CreatedDateTime"`
	CreateLocation  *string             `json:"CreateLocation"`
	UpdatedBy       *string             `json:"UpdatedBy"`
	UpdatedDateTime *string             `json:"UpdatedDateTime"`
}

func (c *Client) GetNoteTypes(ctx context.Context) ([]APINoteType, error) {
	data, err := c.Get(ctx, "/ReferenceData/NoteTypes")
	if err != nil {
		return nil, fmt.Errorf("fetching note types: %w", err)
	}
	var items []APINoteType
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("parsing note types: %w", err)
	}
	return items, nil
}

type APIElectronicAddressType struct {
	Id              *int                `json:"Id"`
	Description     *string             `json:"Description"`
	Inactive        *bool               `json:"Inactive"`
	IsEmail         bool                `json:"IsEmail"`
	ControlGroup    *APIControlGroupRef `json:"ControlGroup"`
	CreatedBy       *string             `json:"CreatedBy"`
	CreatedDateTime *string             `json:"CreatedDateTime"`
	CreateLocation  *string             `json:"CreateLocation"`
	UpdatedBy       *string             `json:"UpdatedBy"`
	UpdatedDateTime *string             `json:"UpdatedDateTime"`
}

func (c *Client) GetElectronicAddressTypes(ctx context.Context) ([]APIElectronicAddressType, error) {
	data, err := c.Get(ctx, "/ReferenceData/ElectronicAddressTypes")
	if err != nil {
		return nil, fmt.Errorf("fetching electronic address types: %w", err)
	}
	var items []APIElectronicAddressType
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("parsing electronic address types: %w", err)
	}
	return items, nil
}

type APIContactPermissionCategory struct {
	Id                 *int                `json:"Id"`
	Description        *string             `json:"Description"`
	Inactive           bool                `json:"Inactive"`
	AskFrequencyMonths *int                `json:"AskFrequencyMonths"`
	ControlGroup       *APIControlGroupRef `json:"ControlGroup"`
	CreatedBy          *string             `json:"CreatedBy"`
	CreatedDateTime    *string             `json:"CreatedDateTime"`
	CreateLocation     *string             `json:"CreateLocation"`
	UpdatedBy          *string             `json:"UpdatedBy"`
	UpdatedDateTime    *string             `json:"UpdatedDateTime"`
}

func (c *Client) GetContactPermissionCategories(ctx context.Context) ([]APIContactPermissionCategory, error) {
	data, err := c.Get(ctx, "/ReferenceData/ContactPermissionCategories")
	if err != nil {
		return nil, fmt.Errorf("fetching contact permission categories: %w", err)
	}
	var items []APIContactPermissionCategory
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("parsing contact permission categories: %w", err)
	}
	return items, nil
}

type APIContactPermissionType struct {
	Id                 *int                          `json:"Id"`
	Description        *string                       `json:"Description"`
	ShortDescription   *string                       `json:"ShortDescription"`
	Inactive           bool                          `json:"Inactive"`
	DefaultValueForAdd *string                       `json:"DefaultValueForAdd"`
	Presenter          bool                          `json:"Presenter"`
	Rank               int                           `json:"Rank"`
	EditIndicator      bool                          `json:"EditIndicator"`
	Category           *APIContactPermissionCategory `json:"Category"`
	CreatedBy          *string                       `json:"CreatedBy"`
	CreatedDateTime    *string                       `json:"CreatedDateTime"`
	CreateLocation     *string                       `json:"CreateLocation"`
	UpdatedBy          *string                       `json:"UpdatedBy"`
	UpdatedDateTime    *string                       `json:"UpdatedDateTime"`
}

func (c *Client) GetContactPermissionTypes(ctx context.Context) ([]APIContactPermissionType, error) {
	data, err := c.Get(ctx, "/ReferenceData/ContactPermissionTypes")
	if err != nil {
		return nil, fmt.Errorf("fetching contact permission types: %w", err)
	}
	var items []APIContactPermissionType
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("parsing contact permission types: %w", err)
	}
	return items, nil
}

type APIRelationshipCategoryRef struct {
	Id          *int    `json:"Id"`
	Description *string `json:"Description"`
	Inactive    *bool   `json:"Inactive"`
}

type APIAffiliationType struct {
	Id                                  *int                        `json:"Id"`
	Description                         *string                     `json:"Description"`
	Inactive                            *bool                       `json:"Inactive"`
	UseSalary                           bool                        `json:"UseSalary"`
	UseTitle                            bool                        `json:"UseTitle"`
	IsAllowedToTransactDefault          bool                        `json:"IsAllowedToTransactDefault"`
	IsIncludedInSearchResultsDefault    bool                        `json:"IsIncludedInSearchResultsDefault"`
	UsesBenefitsWhenInitiator           bool                        `json:"UsesBenefitsWhenInitiator"`
	ControlGroup                        *APIControlGroupRef         `json:"ControlGroup"`
	RelationshipCategory                *APIRelationshipCategoryRef `json:"RelationshipCategory"`
	CreatedBy                           *string                     `json:"CreatedBy"`
	CreatedDateTime                     *string                     `json:"CreatedDateTime"`
	CreateLocation                      *string                     `json:"CreateLocation"`
	UpdatedBy                           *string                     `json:"UpdatedBy"`
	UpdatedDateTime                     *string                     `json:"UpdatedDateTime"`
}

func (c *Client) GetAffiliationTypes(ctx context.Context) ([]APIAffiliationType, error) {
	data, err := c.Get(ctx, "/ReferenceData/AffiliationTypes")
	if err != nil {
		return nil, fmt.Errorf("fetching affiliation types: %w", err)
	}
	var items []APIAffiliationType
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("parsing affiliation types: %w", err)
	}
	return items, nil
}

type APIAssociationType struct {
	Id                                *int                        `json:"Id"`
	Description                       *string                     `json:"Description"`
	Inactive                          *bool                       `json:"Inactive"`
	UseBirthDate                      bool                        `json:"UseBirthDate"`
	UseGender                         bool                        `json:"UseGender"`
	IsIncludedInSearchResultsDefault  bool                        `json:"IsIncludedInSearchResultsDefault"`
	ControlGroup                      *APIControlGroupRef         `json:"ControlGroup"`
	RelationshipCategory              *APIRelationshipCategoryRef `json:"RelationshipCategory"`
	ReciprocalType                    *APIRefItem                 `json:"ReciprocalType"`
	CreatedBy                         *string                     `json:"CreatedBy"`
	CreatedDateTime                   *string                     `json:"CreatedDateTime"`
	CreateLocation                    *string                     `json:"CreateLocation"`
	UpdatedBy                         *string                     `json:"UpdatedBy"`
	UpdatedDateTime                   *string                     `json:"UpdatedDateTime"`
}

func (c *Client) GetAssociationTypes(ctx context.Context) ([]APIAssociationType, error) {
	data, err := c.Get(ctx, "/ReferenceData/AssociationTypes")
	if err != nil {
		return nil, fmt.Errorf("fetching association types: %w", err)
	}
	var items []APIAssociationType
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("parsing association types: %w", err)
	}
	return items, nil
}

type APIAliasType struct {
	Id              *int                `json:"Id"`
	Description     *string             `json:"Description"`
	Inactive        *bool               `json:"Inactive"`
	ControlGroup    *APIControlGroupRef `json:"ControlGroup"`
	CreatedBy       *string             `json:"CreatedBy"`
	CreatedDateTime *string             `json:"CreatedDateTime"`
	CreateLocation  *string             `json:"CreateLocation"`
	UpdatedBy       *string             `json:"UpdatedBy"`
	UpdatedDateTime *string             `json:"UpdatedDateTime"`
}

func (c *Client) GetAliasTypes(ctx context.Context) ([]APIAliasType, error) {
	data, err := c.Get(ctx, "/ReferenceData/AliasTypes")
	if err != nil {
		return nil, fmt.Errorf("fetching alias types: %w", err)
	}
	var items []APIAliasType
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("parsing alias types: %w", err)
	}
	return items, nil
}

type APILoginType struct {
	Id              *int                `json:"Id"`
	Description     *string             `json:"Description"`
	Inactive        *bool               `json:"Inactive"`
	LoginWithEmail  *string             `json:"LoginWithEmail"`
	ControlGroup    *APIControlGroupRef `json:"ControlGroup"`
	CreatedBy       *string             `json:"CreatedBy"`
	CreatedDateTime *string             `json:"CreatedDateTime"`
	CreateLocation  *string             `json:"CreateLocation"`
	UpdatedBy       *string             `json:"UpdatedBy"`
	UpdatedDateTime *string             `json:"UpdatedDateTime"`
}

func (c *Client) GetLoginTypes(ctx context.Context) ([]APILoginType, error) {
	data, err := c.Get(ctx, "/ReferenceData/LoginTypes")
	if err != nil {
		return nil, fmt.Errorf("fetching login types: %w", err)
	}
	var items []APILoginType
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("parsing login types: %w", err)
	}
	return items, nil
}
