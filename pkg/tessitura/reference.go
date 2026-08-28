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

type APIPrefixRef struct {
	Id          *int    `json:"Id"`
	Description *string `json:"Description"`
	Inactive    *bool   `json:"Inactive"`
}

type APIPronounRef struct {
	Id          *int    `json:"Id"`
	Description *string `json:"Description"`
	Inactive    bool    `json:"Inactive"`
}

type APIGender struct {
	Id               *int           `json:"Id"`
	Description      *string        `json:"Description"`
	Inactive         bool           `json:"Inactive"`
	ShortDescription *string        `json:"ShortDescription"`
	DefaultPrefix    *APIPrefixRef  `json:"DefaultPrefix"`
	DefaultPronoun   *APIPronounRef `json:"DefaultPronoun"`
	CreatedBy        *string        `json:"CreatedBy"`
	CreatedDateTime  *string        `json:"CreatedDateTime"`
	CreateLocation   *string        `json:"CreateLocation"`
	UpdatedBy        *string        `json:"UpdatedBy"`
	UpdatedDateTime  *string        `json:"UpdatedDateTime"`
}

func (c *Client) GetGenders(ctx context.Context) ([]APIGender, error) {
	data, err := c.Get(ctx, "/ReferenceData/Genders")
	if err != nil {
		return nil, fmt.Errorf("fetching genders: %w", err)
	}
	var items []APIGender
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("parsing genders: %w", err)
	}
	return items, nil
}

type APIPronoun struct {
	Id              *int    `json:"Id"`
	Description     *string `json:"Description"`
	Inactive        bool    `json:"Inactive"`
	CreatedBy       *string `json:"CreatedBy"`
	CreatedDateTime *string `json:"CreatedDateTime"`
	CreateLocation  *string `json:"CreateLocation"`
	UpdatedBy       *string `json:"UpdatedBy"`
	UpdatedDateTime *string `json:"UpdatedDateTime"`
}

func (c *Client) GetPronouns(ctx context.Context) ([]APIPronoun, error) {
	data, err := c.Get(ctx, "/ReferenceData/Pronouns")
	if err != nil {
		return nil, fmt.Errorf("fetching pronouns: %w", err)
	}
	var items []APIPronoun
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("parsing pronouns: %w", err)
	}
	return items, nil
}

type APIPrefix struct {
	Id              *int    `json:"Id"`
	Description     *string `json:"Description"`
	Inactive        *bool   `json:"Inactive"`
	CreatedBy       *string `json:"CreatedBy"`
	CreatedDateTime *string `json:"CreatedDateTime"`
	CreateLocation  *string `json:"CreateLocation"`
	UpdatedBy       *string `json:"UpdatedBy"`
	UpdatedDateTime *string `json:"UpdatedDateTime"`
}

func (c *Client) GetPrefixes(ctx context.Context) ([]APIPrefix, error) {
	data, err := c.Get(ctx, "/ReferenceData/Prefixes")
	if err != nil {
		return nil, fmt.Errorf("fetching prefixes: %w", err)
	}
	var items []APIPrefix
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("parsing prefixes: %w", err)
	}
	return items, nil
}

type APISuffix struct {
	Id              *int    `json:"Id"`
	Description     *string `json:"Description"`
	Inactive        *bool   `json:"Inactive"`
	CreatedBy       *string `json:"CreatedBy"`
	CreatedDateTime *string `json:"CreatedDateTime"`
	CreateLocation  *string `json:"CreateLocation"`
	UpdatedBy       *string `json:"UpdatedBy"`
	UpdatedDateTime *string `json:"UpdatedDateTime"`
}

func (c *Client) GetSuffixes(ctx context.Context) ([]APISuffix, error) {
	data, err := c.Get(ctx, "/ReferenceData/Suffixes")
	if err != nil {
		return nil, fmt.Errorf("fetching suffixes: %w", err)
	}
	var items []APISuffix
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("parsing suffixes: %w", err)
	}
	return items, nil
}

type APICountry struct {
	Id                   *int    `json:"Id"`
	Description          *string `json:"Description"`
	Inactive             *bool   `json:"Inactive"`
	ShortDescription     *string `json:"ShortDescription"`
	IsoAlpha2Code        *string `json:"IsoAlpha2Code"`
	IsoAlpha3Code        *string `json:"IsoAlpha3Code"`
	PhoneCode            *int    `json:"PhoneCode"`
	DecimalSeparator     *string `json:"DecimalSeparator"`
	PhoneEditString      *string `json:"PhoneEditString"`
	PhoneMask            *string `json:"PhoneMask"`
	MobileEditString     *string `json:"MobileEditString"`
	MobileMask           *string `json:"MobileMask"`
	PostalCodeEditString *string `json:"PostalCodeEditString"`
	PostalCodeMask       *string `json:"PostalCodeMask"`
	PostalCodeValidLengths *string `json:"PostalCodeValidLengths"`
	RequireCity          bool    `json:"RequireCity"`
	RequirePostalCode    bool    `json:"RequirePostalCode"`
	UseAvs               bool    `json:"UseAvs"`
	UseStateField        *string `json:"UseStateField"`
	CreatedBy            *string `json:"CreatedBy"`
	CreatedDateTime      *string `json:"CreatedDateTime"`
	CreateLocation       *string `json:"CreateLocation"`
	UpdatedBy            *string `json:"UpdatedBy"`
	UpdatedDateTime      *string `json:"UpdatedDateTime"`
}

func (c *Client) GetCountries(ctx context.Context) ([]APICountry, error) {
	data, err := c.Get(ctx, "/ReferenceData/Countries")
	if err != nil {
		return nil, fmt.Errorf("fetching countries: %w", err)
	}
	var items []APICountry
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("parsing countries: %w", err)
	}
	return items, nil
}

type APICountryRef struct {
	Id          *int    `json:"Id"`
	Description *string `json:"Description"`
	Inactive    *bool   `json:"Inactive"`
}

type APIState struct {
	Id              *int           `json:"Id"`
	Description     *string        `json:"Description"`
	Inactive        *bool          `json:"Inactive"`
	StateCode       *string        `json:"StateCode"`
	Country         *APICountryRef `json:"Country"`
	CreatedBy       *string        `json:"CreatedBy"`
	CreatedDateTime *string        `json:"CreatedDateTime"`
	CreateLocation  *string        `json:"CreateLocation"`
	UpdatedBy       *string        `json:"UpdatedBy"`
	UpdatedDateTime *string        `json:"UpdatedDateTime"`
}

func (c *Client) GetStates(ctx context.Context) ([]APIState, error) {
	data, err := c.Get(ctx, "/ReferenceData/States")
	if err != nil {
		return nil, fmt.Errorf("fetching states: %w", err)
	}
	var items []APIState
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("parsing states: %w", err)
	}
	return items, nil
}

type APIAddressType struct {
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

func (c *Client) GetAddressTypes(ctx context.Context) ([]APIAddressType, error) {
	data, err := c.Get(ctx, "/ReferenceData/AddressTypes")
	if err != nil {
		return nil, fmt.Errorf("fetching address types: %w", err)
	}
	var items []APIAddressType
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("parsing address types: %w", err)
	}
	return items, nil
}

type APILanguage struct {
	Id              *int    `json:"Id"`
	Description     *string `json:"Description"`
	Inactive        *bool   `json:"Inactive"`
	CreatedBy       *string `json:"CreatedBy"`
	CreatedDateTime *string `json:"CreatedDateTime"`
	CreateLocation  *string `json:"CreateLocation"`
	UpdatedBy       *string `json:"UpdatedBy"`
	UpdatedDateTime *string `json:"UpdatedDateTime"`
}

func (c *Client) GetLanguages(ctx context.Context) ([]APILanguage, error) {
	data, err := c.Get(ctx, "/ReferenceData/Languages")
	if err != nil {
		return nil, fmt.Errorf("fetching languages: %w", err)
	}
	var items []APILanguage
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("parsing languages: %w", err)
	}
	return items, nil
}

type APIConstituencyType struct {
	Id                *int                `json:"Id"`
	Description       *string             `json:"Description"`
	Inactive          *bool               `json:"Inactive"`
	ShortDescription  *string             `json:"ShortDescription"`
	Rank              *int                `json:"Rank"`
	UsedForMemberships bool               `json:"UsedForMemberships"`
	ControlGroup      *APIControlGroupRef `json:"ControlGroup"`
	CreatedBy         *string             `json:"CreatedBy"`
	CreatedDateTime   *string             `json:"CreatedDateTime"`
	CreateLocation    *string             `json:"CreateLocation"`
	UpdatedBy         *string             `json:"UpdatedBy"`
	UpdatedDateTime   *string             `json:"UpdatedDateTime"`
}

func (c *Client) GetConstituencyTypes(ctx context.Context) ([]APIConstituencyType, error) {
	data, err := c.Get(ctx, "/ReferenceData/ConstituencyTypes")
	if err != nil {
		return nil, fmt.Errorf("fetching constituency types: %w", err)
	}
	var items []APIConstituencyType
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("parsing constituency types: %w", err)
	}
	return items, nil
}

type APIConstituentGroup struct {
	Id              *int    `json:"Id"`
	Description     *string `json:"Description"`
	Inactive        *bool   `json:"Inactive"`
	CreatedBy       *string `json:"CreatedBy"`
	CreatedDateTime *string `json:"CreatedDateTime"`
	CreateLocation  *string `json:"CreateLocation"`
	UpdatedBy       *string `json:"UpdatedBy"`
	UpdatedDateTime *string `json:"UpdatedDateTime"`
}

func (c *Client) GetConstituentGroups(ctx context.Context) ([]APIConstituentGroup, error) {
	data, err := c.Get(ctx, "/ReferenceData/ConstituentGroups")
	if err != nil {
		return nil, fmt.Errorf("fetching constituent groups: %w", err)
	}
	var items []APIConstituentGroup
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("parsing constituent groups: %w", err)
	}
	return items, nil
}

type APIKeywordCategory struct {
	Id              *int    `json:"Id"`
	Description     *string `json:"Description"`
	Inactive        *bool   `json:"Inactive"`
	CreatedBy       *string `json:"CreatedBy"`
	CreatedDateTime *string `json:"CreatedDateTime"`
	CreateLocation  *string `json:"CreateLocation"`
	UpdatedBy       *string `json:"UpdatedBy"`
	UpdatedDateTime *string `json:"UpdatedDateTime"`
}

func (c *Client) GetKeywordCategories(ctx context.Context) ([]APIKeywordCategory, error) {
	data, err := c.Get(ctx, "/ReferenceData/KeywordCategories")
	if err != nil {
		return nil, fmt.Errorf("fetching keyword categories: %w", err)
	}
	var items []APIKeywordCategory
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("parsing keyword categories: %w", err)
	}
	return items, nil
}

type APIInterestCategory struct {
	Id              *int    `json:"Id"`
	Description     *string `json:"Description"`
	CreatedBy       *string `json:"CreatedBy"`
	CreatedDateTime *string `json:"CreatedDateTime"`
	UpdatedBy       *string `json:"UpdatedBy"`
	UpdatedDateTime *string `json:"UpdatedDateTime"`
}

func (c *Client) GetInterestCategories(ctx context.Context) ([]APIInterestCategory, error) {
	data, err := c.Get(ctx, "/ReferenceData/InterestCategories")
	if err != nil {
		return nil, fmt.Errorf("fetching interest categories: %w", err)
	}
	var items []APIInterestCategory
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("parsing interest categories: %w", err)
	}
	return items, nil
}

type APIDivisionRef struct {
	Id          *string `json:"Id"`
	Name        *string `json:"Name"`
	Description *string `json:"Description"`
}

type APIOriginalSourceRef struct {
	Id          *int    `json:"Id"`
	Description *string `json:"Description"`
	Inactive    *bool   `json:"Inactive"`
}

type APIUserGroup struct {
	Id                                 *string               `json:"Id"`
	Name                               *string               `json:"Name"`
	Description                        *string               `json:"Description"`
	IsAdmin                            bool                  `json:"IsAdmin"`
	AllowApp                           bool                  `json:"AllowApp"`
	AllowTablet                        bool                  `json:"AllowTablet"`
	AllowTessituraWeb                  bool                  `json:"AllowTessituraWeb"`
	AllowAccessControl                 bool                  `json:"AllowAccessControl"`
	DefaultConstituentHeaderFormatId   *int                  `json:"DefaultConstituentHeaderFormatId"`
	Division                           *APIDivisionRef       `json:"Division"`
	DefaultOriginalSource              *APIOriginalSourceRef `json:"DefaultOriginalSource"`
	CreatedBy                          *string               `json:"CreatedBy"`
	CreatedDateTime                    *string               `json:"CreatedDateTime"`
	CreateLocation                     *string               `json:"CreateLocation"`
	UpdatedBy                          *string               `json:"UpdatedBy"`
	UpdatedDateTime                    *string               `json:"UpdatedDateTime"`
}

func (c *Client) GetUserGroups(ctx context.Context) ([]APIUserGroup, error) {
	data, err := c.Get(ctx, "/ReferenceData/UserGroups")
	if err != nil {
		return nil, fmt.Errorf("fetching user groups: %w", err)
	}
	var items []APIUserGroup
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("parsing user groups: %w", err)
	}
	return items, nil
}

type APIBusinessUnit struct {
	Id              *int    `json:"Id"`
	Description     *string `json:"Description"`
	Inactive        *bool   `json:"Inactive"`
	CreatedBy       *string `json:"CreatedBy"`
	CreatedDateTime *string `json:"CreatedDateTime"`
	CreateLocation  *string `json:"CreateLocation"`
	UpdatedBy       *string `json:"UpdatedBy"`
	UpdatedDateTime *string `json:"UpdatedDateTime"`
}

func (c *Client) GetBusinessUnits(ctx context.Context) ([]APIBusinessUnit, error) {
	data, err := c.Get(ctx, "/ReferenceData/BusinessUnits")
	if err != nil {
		return nil, fmt.Errorf("fetching business units: %w", err)
	}
	var items []APIBusinessUnit
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("parsing business units: %w", err)
	}
	return items, nil
}

type APITheater struct {
	Id                   *int    `json:"Id"`
	Description          *string `json:"Description"`
	Inactive             *bool   `json:"Inactive"`
	Street               *string `json:"Street"`
	City                 *string `json:"City"`
	State                *string `json:"State"`
	PostalCode           *string `json:"PostalCode"`
	Phone                *string `json:"Phone"`
	DrivingDirections    *string `json:"DrivingDirections"`
	DataWindowDefinition *string `json:"DataWindowDefinition"`
	MaximumNumberOfSeats *int    `json:"MaximumNumberOfSeats"`
	CreatedBy            *string `json:"CreatedBy"`
	CreatedDateTime      *string `json:"CreatedDateTime"`
	CreateLocation       *string `json:"CreateLocation"`
	UpdatedBy            *string `json:"UpdatedBy"`
	UpdatedDateTime      *string `json:"UpdatedDateTime"`
}

func (c *Client) GetTheaters(ctx context.Context) ([]APITheater, error) {
	data, err := c.Get(ctx, "/ReferenceData/Theaters")
	if err != nil {
		return nil, fmt.Errorf("fetching theaters: %w", err)
	}
	var items []APITheater
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("parsing theaters: %w", err)
	}
	return items, nil
}

type APISection struct {
	Id              *int    `json:"Id"`
	Description     *string `json:"Description"`
	ShortDesc       *string `json:"ShortDesc"`
	PrintDesc       *string `json:"PrintDesc"`
	PrintSequence   int     `json:"PrintSequence"`
	SectionLegend   *string `json:"SectionLegend"`
	AdditionalText  *string `json:"AdditionalText"`
	AdditionalText2 *string `json:"AdditionalText2"`
	CreatedBy       *string `json:"CreatedBy"`
	CreatedDateTime *string `json:"CreatedDateTime"`
	CreateLocation  *string `json:"CreateLocation"`
	UpdatedBy       *string `json:"UpdatedBy"`
	UpdatedDateTime *string `json:"UpdatedDateTime"`
}

func (c *Client) GetSections(ctx context.Context) ([]APISection, error) {
	data, err := c.Get(ctx, "/ReferenceData/Sections")
	if err != nil {
		return nil, fmt.Errorf("fetching sections: %w", err)
	}
	var items []APISection
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("parsing sections: %w", err)
	}
	return items, nil
}

type APISeatWebSymbolRef struct {
	Id          *int    `json:"Id"`
	Description *string `json:"Description"`
	Inactive    bool    `json:"Inactive"`
	IconName    *string `json:"IconName"`
}

type APISeatCode struct {
	Id               *int                 `json:"Id"`
	Description      *string              `json:"Description"`
	Inactive         bool                 `json:"Inactive"`
	DisplayLetter    *string              `json:"DisplayLetter"`
	AliasDescription *string              `json:"AliasDescription"`
	TicketText       *string              `json:"TicketText"`
	Context          *string              `json:"Context"`
	IsSeat           *int                 `json:"IsSeat"`
	BackColor        *int                 `json:"BackColor"`
	ForeColor        *int                 `json:"ForeColor"`
	SeatWebSymbol    *APISeatWebSymbolRef `json:"SeatWebSymbol"`
	CreatedBy        *string              `json:"CreatedBy"`
	CreatedDateTime  *string              `json:"CreatedDateTime"`
	CreateLocation   *string              `json:"CreateLocation"`
	UpdatedBy        *string              `json:"UpdatedBy"`
	UpdatedDateTime  *string              `json:"UpdatedDateTime"`
}

func (c *Client) GetSeatCodes(ctx context.Context) ([]APISeatCode, error) {
	data, err := c.Get(ctx, "/ReferenceData/SeatCodes")
	if err != nil {
		return nil, fmt.Errorf("fetching seat codes: %w", err)
	}
	var items []APISeatCode
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("parsing seat codes: %w", err)
	}
	return items, nil
}

type APIAppealCategory struct {
	Id              *int    `json:"Id"`
	Description     *string `json:"Description"`
	Inactive        *bool   `json:"Inactive"`
	CreatedBy       *string `json:"CreatedBy"`
	CreatedDateTime *string `json:"CreatedDateTime"`
	CreateLocation  *string `json:"CreateLocation"`
	UpdatedBy       *string `json:"UpdatedBy"`
	UpdatedDateTime *string `json:"UpdatedDateTime"`
}

func (c *Client) GetAppealCategories(ctx context.Context) ([]APIAppealCategory, error) {
	data, err := c.Get(ctx, "/ReferenceData/AppealCategories")
	if err != nil {
		return nil, fmt.Errorf("fetching appeal categories: %w", err)
	}
	var items []APIAppealCategory
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("parsing appeal categories: %w", err)
	}
	return items, nil
}

type APICampaignCategory struct {
	Id              *int    `json:"Id"`
	Description     *string `json:"Description"`
	Inactive        bool    `json:"Inactive"`
	CreatedBy       *string `json:"CreatedBy"`
	CreatedDateTime *string `json:"CreatedDateTime"`
	CreateLocation  *string `json:"CreateLocation"`
	UpdatedBy       *string `json:"UpdatedBy"`
	UpdatedDateTime *string `json:"UpdatedDateTime"`
}

func (c *Client) GetCampaignCategories(ctx context.Context) ([]APICampaignCategory, error) {
	data, err := c.Get(ctx, "/ReferenceData/CampaignCategories")
	if err != nil {
		return nil, fmt.Errorf("fetching campaign categories: %w", err)
	}
	var items []APICampaignCategory
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("parsing campaign categories: %w", err)
	}
	return items, nil
}

type APIContributionDesignation struct {
	Id               *int                `json:"Id"`
	Description      *string             `json:"Description"`
	Inactive         *bool               `json:"Inactive"`
	LetterText       *string             `json:"LetterText"`
	AliasDescription *string             `json:"AliasDescription"`
	ControlGroup     *APIControlGroupRef `json:"ControlGroup"`
	CreatedBy        *string             `json:"CreatedBy"`
	CreatedDateTime  *string             `json:"CreatedDateTime"`
	CreateLocation   *string             `json:"CreateLocation"`
	UpdatedBy        *string             `json:"UpdatedBy"`
	UpdatedDateTime  *string             `json:"UpdatedDateTime"`
}

func (c *Client) GetContributionDesignations(ctx context.Context) ([]APIContributionDesignation, error) {
	data, err := c.Get(ctx, "/ReferenceData/ContributionDesignations")
	if err != nil {
		return nil, fmt.Errorf("fetching contribution designations: %w", err)
	}
	var items []APIContributionDesignation
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("parsing contribution designations: %w", err)
	}
	return items, nil
}

type APIContributionImportSet struct {
	Id                              *int        `json:"Id"`
	Description                     *string     `json:"Description"`
	Inactive                        bool        `json:"Inactive"`
	AccountMatchKeyword             *APIRefItem `json:"AccountMatchKeyword"`
	AcknowledgementLetterMode       *int        `json:"AcknowledgementLetterMode"`
	BatchType                       *APIRefItem `json:"BatchType"`
	BillingSchedule                 *APIRefItem `json:"BillingSchedule"`
	BillingType                     *APIRefItem `json:"BillingType"`
	Campaign                        *APIRefItem `json:"Campaign"`
	SalesChannel                    *APIRefItem `json:"SalesChannel"`
	ContributionDateTime            *string     `json:"ContributionDateTime"`
	ContributionPayMode             int         `json:"ContributionPayMode"`
	CreatePotentialDuplicate        bool        `json:"CreatePotentialDuplicate"`
	CrediteeMode                    *int        `json:"CrediteeMode"`
	CrediteeType                    *APIRefItem `json:"CrediteeType"`
	DefaultCountryCode              *string     `json:"DefaultCountryCode"`
	DefaultConstituentType          *APIRefItem `json:"DefaultConstituentType"`
	DefaultHouseholdConstituentType *APIRefItem `json:"DefaultHouseholdConstituentType"`
	DefaultOriginalSource           *APIRefItem `json:"DefaultOriginalSource"`
	Designation                     *APIRefItem `json:"Designation"`
	FilePath                        *string     `json:"FilePath"`
	FormatFile                      *string     `json:"FormatFile"`
	Fund                            *APIRefItem `json:"Fund"`
	ImportRefNoLocation             *int        `json:"ImportRefNoLocation"`
	PaymentMethod                   *APIRefItem `json:"PaymentMethod"`
	Source                          *APIRefItem `json:"Source"`
	StripPhoneFormatting            bool        `json:"StripPhoneFormatting"`
	TransactAsHousehold             bool        `json:"TransactAsHousehold"`
	TransactAsHouseholdCreditee     bool        `json:"TransactAsHouseholdCreditee"`
	Worker                          *APIRefItem `json:"Worker"`
	CreatedBy                       *string     `json:"CreatedBy"`
	CreatedDateTime                 *string     `json:"CreatedDateTime"`
	CreateLocation                  *string     `json:"CreateLocation"`
	UpdatedBy                       *string     `json:"UpdatedBy"`
	UpdatedDateTime                 *string     `json:"UpdatedDateTime"`
}

func (c *Client) GetContributionImportSets(ctx context.Context) ([]APIContributionImportSet, error) {
	data, err := c.Get(ctx, "/ReferenceData/ContributionImportSets")
	if err != nil {
		return nil, fmt.Errorf("fetching contribution import sets: %w", err)
	}
	var items []APIContributionImportSet
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("parsing contribution import sets: %w", err)
	}
	return items, nil
}

type APIDesignationCode struct {
	Id              *int    `json:"Id"`
	Description     *string `json:"Description"`
	Inactive        *bool   `json:"Inactive"`
	CreatedBy       *string `json:"CreatedBy"`
	CreatedDateTime *string `json:"CreatedDateTime"`
	CreateLocation  *string `json:"CreateLocation"`
	UpdatedBy       *string `json:"UpdatedBy"`
	UpdatedDateTime *string `json:"UpdatedDateTime"`
}

func (c *Client) GetDesignationCodes(ctx context.Context) ([]APIDesignationCode, error) {
	data, err := c.Get(ctx, "/ReferenceData/DesignationCodes")
	if err != nil {
		return nil, fmt.Errorf("fetching designation codes: %w", err)
	}
	var items []APIDesignationCode
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("parsing designation codes: %w", err)
	}
	return items, nil
}

type APIRecognitionType struct {
	Id              *int                `json:"Id"`
	Description     *string             `json:"Description"`
	ControlGroup    *APIControlGroupRef `json:"ControlGroup"`
	CreatedBy       *string             `json:"CreatedBy"`
	CreatedDateTime *string             `json:"CreatedDateTime"`
	CreateLocation  *string             `json:"CreateLocation"`
	UpdatedBy       *string             `json:"UpdatedBy"`
	UpdatedDateTime *string             `json:"UpdatedDateTime"`
}

func (c *Client) GetRecognitionTypes(ctx context.Context) ([]APIRecognitionType, error) {
	data, err := c.Get(ctx, "/ReferenceData/RecognitionTypes")
	if err != nil {
		return nil, fmt.Errorf("fetching recognition types: %w", err)
	}
	var items []APIRecognitionType
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("parsing recognition types: %w", err)
	}
	return items, nil
}

type APIDonationLevel struct {
	Id              *int                `json:"Id"`
	Description     *string             `json:"Description"`
	Inactive        *bool               `json:"Inactive"`
	RecognitionType *APIRecognitionType `json:"RecognitionType"`
	CreatedBy       *string             `json:"CreatedBy"`
	CreatedDateTime *string             `json:"CreatedDateTime"`
	CreateLocation  *string             `json:"CreateLocation"`
	UpdatedBy       *string             `json:"UpdatedBy"`
	UpdatedDateTime *string             `json:"UpdatedDateTime"`
}

func (c *Client) GetDonationLevels(ctx context.Context) ([]APIDonationLevel, error) {
	data, err := c.Get(ctx, "/ReferenceData/DonationLevels")
	if err != nil {
		return nil, fmt.Errorf("fetching donation levels: %w", err)
	}
	var items []APIDonationLevel
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("parsing donation levels: %w", err)
	}
	return items, nil
}

type APIPhilanthropyType struct {
	Id              *int                `json:"Id"`
	Description     *string             `json:"Description"`
	Inactive        bool                `json:"Inactive"`
	ControlGroup    *APIControlGroupRef `json:"ControlGroup"`
	CreatedBy       *string             `json:"CreatedBy"`
	CreatedDateTime *string             `json:"CreatedDateTime"`
	CreateLocation  *string             `json:"CreateLocation"`
	UpdatedBy       *string             `json:"UpdatedBy"`
	UpdatedDateTime *string             `json:"UpdatedDateTime"`
}

func (c *Client) GetPhilanthropyTypes(ctx context.Context) ([]APIPhilanthropyType, error) {
	data, err := c.Get(ctx, "/ReferenceData/PhilanthropyTypes")
	if err != nil {
		return nil, fmt.Errorf("fetching philanthropy types: %w", err)
	}
	var items []APIPhilanthropyType
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("parsing philanthropy types: %w", err)
	}
	return items, nil
}

type APIPlanPriority struct {
	Id              *int                `json:"Id"`
	Description     *string             `json:"Description"`
	Inactive        *bool               `json:"Inactive"`
	Ranking         *int                `json:"Ranking"`
	ControlGroup    *APIControlGroupRef `json:"ControlGroup"`
	CreatedBy       *string             `json:"CreatedBy"`
	CreatedDateTime *string             `json:"CreatedDateTime"`
	CreateLocation  *string             `json:"CreateLocation"`
	UpdatedBy       *string             `json:"UpdatedBy"`
	UpdatedDateTime *string             `json:"UpdatedDateTime"`
}

func (c *Client) GetPlanPriorities(ctx context.Context) ([]APIPlanPriority, error) {
	data, err := c.Get(ctx, "/ReferenceData/PlanPriorities")
	if err != nil {
		return nil, fmt.Errorf("fetching plan priorities: %w", err)
	}
	var items []APIPlanPriority
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("parsing plan priorities: %w", err)
	}
	return items, nil
}

type APIPlanSource struct {
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

func (c *Client) GetPlanSources(ctx context.Context) ([]APIPlanSource, error) {
	data, err := c.Get(ctx, "/ReferenceData/PlanSources")
	if err != nil {
		return nil, fmt.Errorf("fetching plan sources: %w", err)
	}
	var items []APIPlanSource
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("parsing plan sources: %w", err)
	}
	return items, nil
}

type APIPlanStatus struct {
	Id              *int                `json:"Id"`
	Description     *string             `json:"Description"`
	Inactive        *bool               `json:"Inactive"`
	Rank            int                 `json:"Rank"`
	ControlGroup    *APIControlGroupRef `json:"ControlGroup"`
	CreatedBy       *string             `json:"CreatedBy"`
	CreatedDateTime *string             `json:"CreatedDateTime"`
	CreateLocation  *string             `json:"CreateLocation"`
	UpdatedBy       *string             `json:"UpdatedBy"`
	UpdatedDateTime *string             `json:"UpdatedDateTime"`
}

func (c *Client) GetPlanStatuses(ctx context.Context) ([]APIPlanStatus, error) {
	data, err := c.Get(ctx, "/ReferenceData/PlanStatuses")
	if err != nil {
		return nil, fmt.Errorf("fetching plan statuses: %w", err)
	}
	var items []APIPlanStatus
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("parsing plan statuses: %w", err)
	}
	return items, nil
}

type APIPlanType struct {
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

func (c *Client) GetPlanTypes(ctx context.Context) ([]APIPlanType, error) {
	data, err := c.Get(ctx, "/ReferenceData/PlanTypes")
	if err != nil {
		return nil, fmt.Errorf("fetching plan types: %w", err)
	}
	var items []APIPlanType
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("parsing plan types: %w", err)
	}
	return items, nil
}

type APIPlannedGivingCode struct {
	Id              *int    `json:"Id"`
	Description     *string `json:"Description"`
	Inactive        bool    `json:"Inactive"`
	CreatedBy       *string `json:"CreatedBy"`
	CreatedDateTime *string `json:"CreatedDateTime"`
	CreateLocation  *string `json:"CreateLocation"`
	UpdatedBy       *string `json:"UpdatedBy"`
	UpdatedDateTime *string `json:"UpdatedDateTime"`
}

func (c *Client) GetPlannedGivingCodes(ctx context.Context) ([]APIPlannedGivingCode, error) {
	data, err := c.Get(ctx, "/ReferenceData/PlannedGivingCodes")
	if err != nil {
		return nil, fmt.Errorf("fetching planned giving codes: %w", err)
	}
	var items []APIPlannedGivingCode
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("parsing planned giving codes: %w", err)
	}
	return items, nil
}

type APIPlannedGivingFunding struct {
	Id              *int    `json:"Id"`
	Description     *string `json:"Description"`
	Inactive        bool    `json:"Inactive"`
	CreatedBy       *string `json:"CreatedBy"`
	CreatedDateTime *string `json:"CreatedDateTime"`
	CreateLocation  *string `json:"CreateLocation"`
	UpdatedBy       *string `json:"UpdatedBy"`
	UpdatedDateTime *string `json:"UpdatedDateTime"`
}

func (c *Client) GetPlannedGivingFundings(ctx context.Context) ([]APIPlannedGivingFunding, error) {
	data, err := c.Get(ctx, "/ReferenceData/PlannedGivingFundings")
	if err != nil {
		return nil, fmt.Errorf("fetching planned giving fundings: %w", err)
	}
	var items []APIPlannedGivingFunding
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("parsing planned giving fundings: %w", err)
	}
	return items, nil
}

type APIPlannedGivingGiftType struct {
	Id              *int                `json:"Id"`
	Description     *string             `json:"Description"`
	Inactive        bool                `json:"Inactive"`
	ControlGroup    *APIControlGroupRef `json:"ControlGroup"`
	CreatedBy       *string             `json:"CreatedBy"`
	CreatedDateTime *string             `json:"CreatedDateTime"`
	CreateLocation  *string             `json:"CreateLocation"`
	UpdatedBy       *string             `json:"UpdatedBy"`
	UpdatedDateTime *string             `json:"UpdatedDateTime"`
}

func (c *Client) GetPlannedGivingGiftTypes(ctx context.Context) ([]APIPlannedGivingGiftType, error) {
	data, err := c.Get(ctx, "/ReferenceData/PlannedGivingGiftTypes")
	if err != nil {
		return nil, fmt.Errorf("fetching planned giving gift types: %w", err)
	}
	var items []APIPlannedGivingGiftType
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("parsing planned giving gift types: %w", err)
	}
	return items, nil
}

type APIPlannedGivingOnFile struct {
	Id              *int    `json:"Id"`
	Description     *string `json:"Description"`
	Inactive        bool    `json:"Inactive"`
	CreatedBy       *string `json:"CreatedBy"`
	CreatedDateTime *string `json:"CreatedDateTime"`
	CreateLocation  *string `json:"CreateLocation"`
	UpdatedBy       *string `json:"UpdatedBy"`
	UpdatedDateTime *string `json:"UpdatedDateTime"`
}

func (c *Client) GetPlannedGivingOnFiles(ctx context.Context) ([]APIPlannedGivingOnFile, error) {
	data, err := c.Get(ctx, "/ReferenceData/PlannedGivingOnFiles")
	if err != nil {
		return nil, fmt.Errorf("fetching planned giving on-file statuses: %w", err)
	}
	var items []APIPlannedGivingOnFile
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("parsing planned giving on-file statuses: %w", err)
	}
	return items, nil
}

type APIPlannedGivingPurpose struct {
	Id              *int    `json:"Id"`
	Description     *string `json:"Description"`
	Inactive        bool    `json:"Inactive"`
	CreatedBy       *string `json:"CreatedBy"`
	CreatedDateTime *string `json:"CreatedDateTime"`
	CreateLocation  *string `json:"CreateLocation"`
	UpdatedBy       *string `json:"UpdatedBy"`
	UpdatedDateTime *string `json:"UpdatedDateTime"`
}

func (c *Client) GetPlannedGivingPurposes(ctx context.Context) ([]APIPlannedGivingPurpose, error) {
	data, err := c.Get(ctx, "/ReferenceData/PlannedGivingPurposes")
	if err != nil {
		return nil, fmt.Errorf("fetching planned giving purposes: %w", err)
	}
	var items []APIPlannedGivingPurpose
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("parsing planned giving purposes: %w", err)
	}
	return items, nil
}

type APIPlannedGivingSource struct {
	Id              *int    `json:"Id"`
	Description     *string `json:"Description"`
	Inactive        bool    `json:"Inactive"`
	CreatedBy       *string `json:"CreatedBy"`
	CreatedDateTime *string `json:"CreatedDateTime"`
	CreateLocation  *string `json:"CreateLocation"`
	UpdatedBy       *string `json:"UpdatedBy"`
	UpdatedDateTime *string `json:"UpdatedDateTime"`
}

func (c *Client) GetPlannedGivingSources(ctx context.Context) ([]APIPlannedGivingSource, error) {
	data, err := c.Get(ctx, "/ReferenceData/PlannedGivingSources")
	if err != nil {
		return nil, fmt.Errorf("fetching planned giving sources: %w", err)
	}
	var items []APIPlannedGivingSource
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("parsing planned giving sources: %w", err)
	}
	return items, nil
}

type APIPlannedGivingStatus struct {
	Id              *int    `json:"Id"`
	Description     *string `json:"Description"`
	Inactive        bool    `json:"Inactive"`
	CreatedBy       *string `json:"CreatedBy"`
	CreatedDateTime *string `json:"CreatedDateTime"`
	CreateLocation  *string `json:"CreateLocation"`
	UpdatedBy       *string `json:"UpdatedBy"`
	UpdatedDateTime *string `json:"UpdatedDateTime"`
}

func (c *Client) GetPlannedGivingStatuses(ctx context.Context) ([]APIPlannedGivingStatus, error) {
	data, err := c.Get(ctx, "/ReferenceData/PlannedGivingStatuses")
	if err != nil {
		return nil, fmt.Errorf("fetching planned giving statuses: %w", err)
	}
	var items []APIPlannedGivingStatus
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("parsing planned giving statuses: %w", err)
	}
	return items, nil
}

type APIMembershipBenefitFrequency struct {
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

func (c *Client) GetMembershipBenefitFrequencies(ctx context.Context) ([]APIMembershipBenefitFrequency, error) {
	data, err := c.Get(ctx, "/ReferenceData/MembershipBenefitFrequencies")
	if err != nil {
		return nil, fmt.Errorf("fetching membership benefit frequencies: %w", err)
	}
	var items []APIMembershipBenefitFrequency
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("parsing membership benefit frequencies: %w", err)
	}
	return items, nil
}

type APIMembershipBenefitType struct {
	Id              *int    `json:"Id"`
	Description     *string `json:"Description"`
	Inactive        *bool   `json:"Inactive"`
	CreatedBy       *string `json:"CreatedBy"`
	CreatedDateTime *string `json:"CreatedDateTime"`
	CreateLocation  *string `json:"CreateLocation"`
	UpdatedBy       *string `json:"UpdatedBy"`
	UpdatedDateTime *string `json:"UpdatedDateTime"`
}

func (c *Client) GetMembershipBenefitTypes(ctx context.Context) ([]APIMembershipBenefitType, error) {
	data, err := c.Get(ctx, "/ReferenceData/MembershipBenefitTypes")
	if err != nil {
		return nil, fmt.Errorf("fetching membership benefit types: %w", err)
	}
	var items []APIMembershipBenefitType
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("parsing membership benefit types: %w", err)
	}
	return items, nil
}

type APIMembershipLevelCategory struct {
	Id              *int    `json:"Id"`
	Description     *string `json:"Description"`
	Inactive        *bool   `json:"Inactive"`
	CreatedBy       *string `json:"CreatedBy"`
	CreatedDateTime *string `json:"CreatedDateTime"`
	CreateLocation  *string `json:"CreateLocation"`
	UpdatedBy       *string `json:"UpdatedBy"`
	UpdatedDateTime *string `json:"UpdatedDateTime"`
}

func (c *Client) GetMembershipLevelCategories(ctx context.Context) ([]APIMembershipLevelCategory, error) {
	data, err := c.Get(ctx, "/ReferenceData/MembershipLevelCategories")
	if err != nil {
		return nil, fmt.Errorf("fetching membership level categories: %w", err)
	}
	var items []APIMembershipLevelCategory
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("parsing membership level categories: %w", err)
	}
	return items, nil
}

type APIMembershipLevelTrend struct {
	Id              *int    `json:"Id"`
	Description     *string `json:"Description"`
	CreatedBy       *string `json:"CreatedBy"`
	CreatedDateTime *string `json:"CreatedDateTime"`
	CreateLocation  *string `json:"CreateLocation"`
	UpdatedBy       *string `json:"UpdatedBy"`
	UpdatedDateTime *string `json:"UpdatedDateTime"`
}

func (c *Client) GetMembershipLevelTrends(ctx context.Context) ([]APIMembershipLevelTrend, error) {
	data, err := c.Get(ctx, "/ReferenceData/MembershipLevelTrends")
	if err != nil {
		return nil, fmt.Errorf("fetching membership level trends: %w", err)
	}
	var items []APIMembershipLevelTrend
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("parsing membership level trends: %w", err)
	}
	return items, nil
}

type APIMembershipPeriod struct {
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

func (c *Client) GetMembershipPeriods(ctx context.Context) ([]APIMembershipPeriod, error) {
	data, err := c.Get(ctx, "/ReferenceData/MembershipPeriods")
	if err != nil {
		return nil, fmt.Errorf("fetching membership periods: %w", err)
	}
	var items []APIMembershipPeriod
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("parsing membership periods: %w", err)
	}
	return items, nil
}

type APIMembershipStanding struct {
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

func (c *Client) GetMembershipStandings(ctx context.Context) ([]APIMembershipStanding, error) {
	data, err := c.Get(ctx, "/ReferenceData/MembershipStandings")
	if err != nil {
		return nil, fmt.Errorf("fetching membership standings: %w", err)
	}
	var items []APIMembershipStanding
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("parsing membership standings: %w", err)
	}
	return items, nil
}

type APIMembershipStatus struct {
	Id              *int    `json:"Id"`
	Description     *string `json:"Description"`
	Rank            int     `json:"Rank"`
	CreatedBy       *string `json:"CreatedBy"`
	CreatedDateTime *string `json:"CreatedDateTime"`
	CreateLocation  *string `json:"CreateLocation"`
	UpdatedBy       *string `json:"UpdatedBy"`
	UpdatedDateTime *string `json:"UpdatedDateTime"`
}

func (c *Client) GetMembershipStatuses(ctx context.Context) ([]APIMembershipStatus, error) {
	data, err := c.Get(ctx, "/ReferenceData/MembershipStatuses")
	if err != nil {
		return nil, fmt.Errorf("fetching membership statuses: %w", err)
	}
	var items []APIMembershipStatus
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("parsing membership statuses: %w", err)
	}
	return items, nil
}

type APIConstituentProtectionType struct {
	Id              *int                `json:"Id"`
	Description     *string             `json:"Description"`
	Inactive        bool                `json:"Inactive"`
	ControlGroup    *APIControlGroupRef `json:"ControlGroup"`
	CreatedBy       *string             `json:"CreatedBy"`
	CreatedDateTime *string             `json:"CreatedDateTime"`
	CreateLocation  *string             `json:"CreateLocation"`
	UpdatedBy       *string             `json:"UpdatedBy"`
	UpdatedDateTime *string             `json:"UpdatedDateTime"`
}

func (c *Client) GetConstituentProtectionTypes(ctx context.Context) ([]APIConstituentProtectionType, error) {
	data, err := c.Get(ctx, "/ReferenceData/ConstituentProtectionTypes")
	if err != nil {
		return nil, fmt.Errorf("fetching constituent protection types: %w", err)
	}
	var items []APIConstituentProtectionType
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("parsing constituent protection types: %w", err)
	}
	return items, nil
}

type APIConstituentTypeAffiliate struct {
	Id               *int                      `json:"Id"`
	ConstituentType  *APIConstituentTypeSummary `json:"ConstituentType"`
	AffiliationType  *APIAffiliationTypeSummary `json:"AffiliationType"`
	HouseholdPrimary *bool                     `json:"HouseholdPrimary"`
	ShowWithGroup    *bool                     `json:"ShowWithGroup"`
	Rank             *int                      `json:"Rank"`
	CreatedBy        *string                   `json:"CreatedBy"`
	CreatedDateTime  *string                   `json:"CreatedDateTime"`
	CreateLocation   *string                   `json:"CreateLocation"`
	UpdatedBy        *string                   `json:"UpdatedBy"`
	UpdatedDateTime  *string                   `json:"UpdatedDateTime"`
}

func (c *Client) GetConstituentTypeAffiliates(ctx context.Context) ([]APIConstituentTypeAffiliate, error) {
	data, err := c.Get(ctx, "/ReferenceData/ConstituentTypeAffiliates")
	if err != nil {
		return nil, fmt.Errorf("fetching constituent type affiliates: %w", err)
	}
	var items []APIConstituentTypeAffiliate
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("parsing constituent type affiliates: %w", err)
	}
	return items, nil
}

type APIContactLogActivityType struct {
	Id              *int    `json:"Id"`
	Description     *string `json:"Description"`
	Inactive        *bool   `json:"Inactive"`
	CreatedBy       *string `json:"CreatedBy"`
	CreatedDateTime *string `json:"CreatedDateTime"`
	CreateLocation  *string `json:"CreateLocation"`
	UpdatedBy       *string `json:"UpdatedBy"`
	UpdatedDateTime *string `json:"UpdatedDateTime"`
}

func (c *Client) GetContactLogActivityTypes(ctx context.Context) ([]APIContactLogActivityType, error) {
	data, err := c.Get(ctx, "/ReferenceData/ContactLogActivityTypes")
	if err != nil {
		return nil, fmt.Errorf("fetching contact log activity types: %w", err)
	}
	var items []APIContactLogActivityType
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("parsing contact log activity types: %w", err)
	}
	return items, nil
}

type APIContactPointCategory struct {
	Id                *int    `json:"Id"`
	Description       *string `json:"Description"`
	ContactPointKey   *string `json:"ContactPointKey"`
	ContactPointTable *string `json:"ContactPointTable"`
	CreatedBy         *string `json:"CreatedBy"`
	CreatedDateTime   *string `json:"CreatedDateTime"`
	CreateLocation    *string `json:"CreateLocation"`
	UpdatedBy         *string `json:"UpdatedBy"`
	UpdatedDateTime   *string `json:"UpdatedDateTime"`
}

func (c *Client) GetContactPointCategories(ctx context.Context) ([]APIContactPointCategory, error) {
	data, err := c.Get(ctx, "/ReferenceData/ContactPointCategories")
	if err != nil {
		return nil, fmt.Errorf("fetching contact point categories: %w", err)
	}
	var items []APIContactPointCategory
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("parsing contact point categories: %w", err)
	}
	return items, nil
}

type APIContactPointCategorySummary struct {
	Id          *int    `json:"Id"`
	Description *string `json:"Description"`
}

type APIPurposeCategorySummary struct {
	Id          *int    `json:"Id"`
	Description *string `json:"Description"`
	Inactive    bool    `json:"Inactive"`
}

type APIContactPointPurposeSummary struct {
	Id              *int                      `json:"Id"`
	Description     *string                   `json:"Description"`
	Inactive        bool                      `json:"Inactive"`
	PurposeCategory *APIPurposeCategorySummary `json:"PurposeCategory"`
}

type APIContactPointCategoryPurpose struct {
	Id                   *int                            `json:"Id"`
	ContactPointCategory *APIContactPointCategorySummary `json:"ContactPointCategory"`
	Purpose              *APIContactPointPurposeSummary  `json:"Purpose"`
	CreatedBy            *string                         `json:"CreatedBy"`
	CreatedDateTime      *string                         `json:"CreatedDateTime"`
	CreateLocation       *string                         `json:"CreateLocation"`
	UpdatedBy            *string                         `json:"UpdatedBy"`
	UpdatedDateTime      *string                         `json:"UpdatedDateTime"`
}

func (c *Client) GetContactPointCategoryPurposes(ctx context.Context) ([]APIContactPointCategoryPurpose, error) {
	data, err := c.Get(ctx, "/ReferenceData/ContactPointCategoryPurposes")
	if err != nil {
		return nil, fmt.Errorf("fetching contact point category purposes: %w", err)
	}
	var items []APIContactPointCategoryPurpose
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("parsing contact point category purposes: %w", err)
	}
	return items, nil
}

type APIContactPointPurposeCategory struct {
	Id              *int                `json:"Id"`
	Description     *string             `json:"Description"`
	Inactive        bool                `json:"Inactive"`
	ControlGroup    *APIControlGroupRef `json:"ControlGroup"`
	CreatedBy       *string             `json:"CreatedBy"`
	CreatedDateTime *string             `json:"CreatedDateTime"`
	CreateLocation  *string             `json:"CreateLocation"`
	UpdatedBy       *string             `json:"UpdatedBy"`
	UpdatedDateTime *string             `json:"UpdatedDateTime"`
}

func (c *Client) GetContactPointPurposeCategories(ctx context.Context) ([]APIContactPointPurposeCategory, error) {
	data, err := c.Get(ctx, "/ReferenceData/ContactPointPurposeCategories")
	if err != nil {
		return nil, fmt.Errorf("fetching contact point purpose categories: %w", err)
	}
	var items []APIContactPointPurposeCategory
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("parsing contact point purpose categories: %w", err)
	}
	return items, nil
}

type APIContactPointPurpose struct {
	Id              *int                      `json:"Id"`
	Description     *string                   `json:"Description"`
	Inactive        bool                      `json:"Inactive"`
	PurposeCategory *APIPurposeCategorySummary `json:"PurposeCategory"`
	CreatedBy       *string                   `json:"CreatedBy"`
	CreatedDateTime *string                   `json:"CreatedDateTime"`
	CreateLocation  *string                   `json:"CreateLocation"`
	UpdatedBy       *string                   `json:"UpdatedBy"`
	UpdatedDateTime *string                   `json:"UpdatedDateTime"`
}

func (c *Client) GetContactPointPurposes(ctx context.Context) ([]APIContactPointPurpose, error) {
	data, err := c.Get(ctx, "/ReferenceData/ContactPointPurposes")
	if err != nil {
		return nil, fmt.Errorf("fetching contact point purposes: %w", err)
	}
	var items []APIContactPointPurpose
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("parsing contact point purposes: %w", err)
	}
	return items, nil
}

type APIContactType struct {
	Id              *int    `json:"Id"`
	Description     *string `json:"Description"`
	Inactive        *bool   `json:"Inactive"`
	CreatedBy       *string `json:"CreatedBy"`
	CreatedDateTime *string `json:"CreatedDateTime"`
	CreateLocation  *string `json:"CreateLocation"`
	UpdatedBy       *string `json:"UpdatedBy"`
	UpdatedDateTime *string `json:"UpdatedDateTime"`
}

func (c *Client) GetContactTypes(ctx context.Context) ([]APIContactType, error) {
	data, err := c.Get(ctx, "/ReferenceData/ContactTypes")
	if err != nil {
		return nil, fmt.Errorf("fetching contact types: %w", err)
	}
	var items []APIContactType
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("parsing contact types: %w", err)
	}
	return items, nil
}

type APIKeywordSummary struct {
	Id          int     `json:"Id"`
	Description *string `json:"Description"`
}

type APIKeywordConstituentType struct {
	Id              *int                      `json:"Id"`
	Inactive        *bool                     `json:"Inactive"`
	Rank            *int                      `json:"Rank"`
	Keyword         *APIKeywordSummary        `json:"Keyword"`
	ConstituentType *APIConstituentTypeSummary `json:"ConstituentType"`
	ControlGroup    *APIControlGroupRef       `json:"ControlGroup"`
}

func (c *Client) GetKeywordConstituentTypes(ctx context.Context) ([]APIKeywordConstituentType, error) {
	data, err := c.Get(ctx, "/ReferenceData/KeywordConstituentTypes")
	if err != nil {
		return nil, fmt.Errorf("fetching keyword constituent types: %w", err)
	}
	var items []APIKeywordConstituentType
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("parsing keyword constituent types: %w", err)
	}
	return items, nil
}

type APIMailIndicator struct {
	Id              *int    `json:"Id"`
	Description     *string `json:"Description"`
	Inactive        *bool   `json:"Inactive"`
	CreatedBy       *string `json:"CreatedBy"`
	CreatedDateTime *string `json:"CreatedDateTime"`
	CreateLocation  *string `json:"CreateLocation"`
	UpdatedBy       *string `json:"UpdatedBy"`
	UpdatedDateTime *string `json:"UpdatedDateTime"`
}

func (c *Client) GetMailIndicators(ctx context.Context) ([]APIMailIndicator, error) {
	data, err := c.Get(ctx, "/ReferenceData/MailIndicators")
	if err != nil {
		return nil, fmt.Errorf("fetching mail indicators: %w", err)
	}
	var items []APIMailIndicator
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("parsing mail indicators: %w", err)
	}
	return items, nil
}

type APINameStatus struct {
	Id              *int    `json:"Id"`
	Description     *string `json:"Description"`
	Inactive        *bool   `json:"Inactive"`
	CreatedBy       *string `json:"CreatedBy"`
	CreatedDateTime *string `json:"CreatedDateTime"`
	CreateLocation  *string `json:"CreateLocation"`
	UpdatedBy       *string `json:"UpdatedBy"`
	UpdatedDateTime *string `json:"UpdatedDateTime"`
}

func (c *Client) GetNameStatuses(ctx context.Context) ([]APINameStatus, error) {
	data, err := c.Get(ctx, "/ReferenceData/NameStatuses")
	if err != nil {
		return nil, fmt.Errorf("fetching name statuses: %w", err)
	}
	var items []APINameStatus
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("parsing name statuses: %w", err)
	}
	return items, nil
}

type APIPhoneIndicator struct {
	Id              *int    `json:"Id"`
	Description     *string `json:"Description"`
	Inactive        *bool   `json:"Inactive"`
	CreatedBy       *string `json:"CreatedBy"`
	CreatedDateTime *string `json:"CreatedDateTime"`
	CreateLocation  *string `json:"CreateLocation"`
	UpdatedBy       *string `json:"UpdatedBy"`
	UpdatedDateTime *string `json:"UpdatedDateTime"`
}

func (c *Client) GetPhoneIndicators(ctx context.Context) ([]APIPhoneIndicator, error) {
	data, err := c.Get(ctx, "/ReferenceData/PhoneIndicators")
	if err != nil {
		return nil, fmt.Errorf("fetching phone indicators: %w", err)
	}
	var items []APIPhoneIndicator
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("parsing phone indicators: %w", err)
	}
	return items, nil
}

type APIPhoneType struct {
	Id                 *int                `json:"Id"`
	Description        *string             `json:"Description"`
	Inactive           *bool               `json:"Inactive"`
	AllowTelemarketing *bool               `json:"AllowTelemarketing"`
	ControlGroup       *APIControlGroupRef `json:"ControlGroup"`
	CreatedBy          *string             `json:"CreatedBy"`
	CreatedDateTime    *string             `json:"CreatedDateTime"`
	CreateLocation     *string             `json:"CreateLocation"`
	UpdatedBy          *string             `json:"UpdatedBy"`
	UpdatedDateTime    *string             `json:"UpdatedDateTime"`
}

func (c *Client) GetPhoneTypes(ctx context.Context) ([]APIPhoneType, error) {
	data, err := c.Get(ctx, "/ReferenceData/PhoneTypes")
	if err != nil {
		return nil, fmt.Errorf("fetching phone types: %w", err)
	}
	var items []APIPhoneType
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("parsing phone types: %w", err)
	}
	return items, nil
}

type APIRelationshipCategory struct {
	Id              *int    `json:"Id"`
	Description     *string `json:"Description"`
	Inactive        *bool   `json:"Inactive"`
	CreatedBy       *string `json:"CreatedBy"`
	CreatedDateTime *string `json:"CreatedDateTime"`
	UpdatedBy       *string `json:"UpdatedBy"`
	UpdatedDateTime *string `json:"UpdatedDateTime"`
}

func (c *Client) GetRelationshipCategories(ctx context.Context) ([]APIRelationshipCategory, error) {
	data, err := c.Get(ctx, "/ReferenceData/RelationshipCategories")
	if err != nil {
		return nil, fmt.Errorf("fetching relationship categories: %w", err)
	}
	var items []APIRelationshipCategory
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("parsing relationship categories: %w", err)
	}
	return items, nil
}

type APIDiscountType struct {
	Id                  *int    `json:"Id"`
	Description         *string `json:"Description"`
	Inactive            bool    `json:"Inactive"`
	Amount              float64 `json:"Amount"`
	PercentIndicator    bool    `json:"PercentIndicator"`
	RespectMinimumPrice bool    `json:"RespectMinimumPrice"`
	ShortDescription    *string `json:"ShortDescription"`
	CreatedBy           *string `json:"CreatedBy"`
	CreatedDateTime     *string `json:"CreatedDateTime"`
	CreateLocation      *string `json:"CreateLocation"`
	UpdatedBy           *string `json:"UpdatedBy"`
	UpdatedDateTime     *string `json:"UpdatedDateTime"`
}

func (c *Client) GetDiscountTypes(ctx context.Context) ([]APIDiscountType, error) {
	data, err := c.Get(ctx, "/ReferenceData/DiscountTypes")
	if err != nil {
		return nil, fmt.Errorf("fetching discount types: %w", err)
	}
	var items []APIDiscountType
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("parsing discount types: %w", err)
	}
	return items, nil
}

type APIEventLevel struct {
	Id              *int    `json:"Id"`
	Description     *string `json:"Description"`
	Inactive        bool    `json:"Inactive"`
	CreatedBy       *string `json:"CreatedBy"`
	CreatedDateTime *string `json:"CreatedDateTime"`
	CreateLocation  *string `json:"CreateLocation"`
	UpdatedBy       *string `json:"UpdatedBy"`
	UpdatedDateTime *string `json:"UpdatedDateTime"`
}

func (c *Client) GetEventLevels(ctx context.Context) ([]APIEventLevel, error) {
	data, err := c.Get(ctx, "/ReferenceData/EventLevels")
	if err != nil {
		return nil, fmt.Errorf("fetching event levels: %w", err)
	}
	var items []APIEventLevel
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("parsing event levels: %w", err)
	}
	return items, nil
}

type APIFeeCategory struct {
	Id              *int    `json:"Id"`
	Description     *string `json:"Description"`
	Inactive        bool    `json:"Inactive"`
	CreatedBy       *string `json:"CreatedBy"`
	CreatedDateTime *string `json:"CreatedDateTime"`
	CreateLocation  *string `json:"CreateLocation"`
	UpdatedBy       *string `json:"UpdatedBy"`
	UpdatedDateTime *string `json:"UpdatedDateTime"`
}

func (c *Client) GetFeeCategories(ctx context.Context) ([]APIFeeCategory, error) {
	data, err := c.Get(ctx, "/ReferenceData/FeeCategories")
	if err != nil {
		return nil, fmt.Errorf("fetching fee categories: %w", err)
	}
	var items []APIFeeCategory
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("parsing fee categories: %w", err)
	}
	return items, nil
}

type APIHoldCodeCategory struct {
	Id              *int                `json:"Id"`
	Description     *string             `json:"Description"`
	Inactive        bool                `json:"Inactive"`
	ControlGroup    *APIControlGroupRef `json:"ControlGroup"`
	CreatedBy       *string             `json:"CreatedBy"`
	CreatedDateTime *string             `json:"CreatedDateTime"`
	CreateLocation  *string             `json:"CreateLocation"`
	UpdatedBy       *string             `json:"UpdatedBy"`
	UpdatedDateTime *string             `json:"UpdatedDateTime"`
}

func (c *Client) GetHoldCodeCategories(ctx context.Context) ([]APIHoldCodeCategory, error) {
	data, err := c.Get(ctx, "/ReferenceData/HoldCodeCategories")
	if err != nil {
		return nil, fmt.Errorf("fetching hold code categories: %w", err)
	}
	var items []APIHoldCodeCategory
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("parsing hold code categories: %w", err)
	}
	return items, nil
}

type APIPackageType struct {
	Id              *int    `json:"Id"`
	Description     *string `json:"Description"`
	Inactive        *bool   `json:"Inactive"`
	CreatedBy       *string `json:"CreatedBy"`
	CreatedDateTime *string `json:"CreatedDateTime"`
	CreateLocation  *string `json:"CreateLocation"`
	UpdatedBy       *string `json:"UpdatedBy"`
	UpdatedDateTime *string `json:"UpdatedDateTime"`
}

func (c *Client) GetPackageTypes(ctx context.Context) ([]APIPackageType, error) {
	data, err := c.Get(ctx, "/ReferenceData/PackageTypes")
	if err != nil {
		return nil, fmt.Errorf("fetching package types: %w", err)
	}
	var items []APIPackageType
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("parsing package types: %w", err)
	}
	return items, nil
}

type APIPerformanceSegmentType struct {
	Id              *int    `json:"Id"`
	Description     *string `json:"Description"`
	Inactive        bool    `json:"Inactive"`
	CreatedBy       *string `json:"CreatedBy"`
	CreatedDateTime *string `json:"CreatedDateTime"`
	CreateLocation  *string `json:"CreateLocation"`
	UpdatedBy       *string `json:"UpdatedBy"`
	UpdatedDateTime *string `json:"UpdatedDateTime"`
}

func (c *Client) GetPerformanceSegmentTypes(ctx context.Context) ([]APIPerformanceSegmentType, error) {
	data, err := c.Get(ctx, "/ReferenceData/PerformanceSegmentTypes")
	if err != nil {
		return nil, fmt.Errorf("fetching performance segment types: %w", err)
	}
	var items []APIPerformanceSegmentType
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("parsing performance segment types: %w", err)
	}
	return items, nil
}

type APIPriceCategorySummary struct {
	Id          int     `json:"Id"`
	Description *string `json:"Description"`
	Inactive    bool    `json:"Inactive"`
	Rank        int     `json:"Rank"`
}

type APIPriceLayerType struct {
	Id               *int                    `json:"Id"`
	Description      *string                 `json:"Description"`
	Inactive         *bool                   `json:"Inactive"`
	ExcludeFromRules bool                    `json:"ExcludeFromRules"`
	Rank             int                     `json:"Rank"`
	PriceCategory    *APIPriceCategorySummary `json:"PriceCategory"`
	CreatedBy        *string                 `json:"CreatedBy"`
	CreatedDateTime  *string                 `json:"CreatedDateTime"`
	CreateLocation   *string                 `json:"CreateLocation"`
	UpdatedBy        *string                 `json:"UpdatedBy"`
	UpdatedDateTime  *string                 `json:"UpdatedDateTime"`
}

func (c *Client) GetPriceLayerTypes(ctx context.Context) ([]APIPriceLayerType, error) {
	data, err := c.Get(ctx, "/ReferenceData/PriceLayerTypes")
	if err != nil {
		return nil, fmt.Errorf("fetching price layer types: %w", err)
	}
	var items []APIPriceLayerType
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("parsing price layer types: %w", err)
	}
	return items, nil
}

type APIPriceTypeReason struct {
	Id               *int    `json:"Id"`
	Description      *string `json:"Description"`
	Inactive         *bool   `json:"Inactive"`
	ShortDescription *string `json:"ShortDescription"`
	TicketText       *string `json:"TicketText"`
	CreatedBy        *string `json:"CreatedBy"`
	CreatedDateTime  *string `json:"CreatedDateTime"`
	CreateLocation   *string `json:"CreateLocation"`
	UpdatedBy        *string `json:"UpdatedBy"`
	UpdatedDateTime  *string `json:"UpdatedDateTime"`
}

func (c *Client) GetPriceTypeReasons(ctx context.Context) ([]APIPriceTypeReason, error) {
	data, err := c.Get(ctx, "/ReferenceData/PriceTypeReasons")
	if err != nil {
		return nil, fmt.Errorf("fetching price type reasons: %w", err)
	}
	var items []APIPriceTypeReason
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("parsing price type reasons: %w", err)
	}
	return items, nil
}

type APIPricingRuleCategory struct {
	Id              *int                `json:"Id"`
	Description     *string             `json:"Description"`
	Inactive        bool                `json:"Inactive"`
	ControlGroup    *APIControlGroupRef `json:"ControlGroup"`
	CreatedBy       *string             `json:"CreatedBy"`
	CreatedDateTime *string             `json:"CreatedDateTime"`
	CreateLocation  *string             `json:"CreateLocation"`
	UpdatedBy       *string             `json:"UpdatedBy"`
	UpdatedDateTime *string             `json:"UpdatedDateTime"`
}

func (c *Client) GetPricingRuleCategories(ctx context.Context) ([]APIPricingRuleCategory, error) {
	data, err := c.Get(ctx, "/ReferenceData/PricingRuleCategories")
	if err != nil {
		return nil, fmt.Errorf("fetching pricing rule categories: %w", err)
	}
	var items []APIPricingRuleCategory
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("parsing pricing rule categories: %w", err)
	}
	return items, nil
}

type APIPricingRuleMessageType struct {
	Id              *int    `json:"Id"`
	Description     *string `json:"Description"`
	Inactive        bool    `json:"Inactive"`
	CreatedBy       *string `json:"CreatedBy"`
	CreatedDateTime *string `json:"CreatedDateTime"`
	CreateLocation  *string `json:"CreateLocation"`
	UpdatedBy       *string `json:"UpdatedBy"`
	UpdatedDateTime *string `json:"UpdatedDateTime"`
}

func (c *Client) GetPricingRuleMessageTypes(ctx context.Context) ([]APIPricingRuleMessageType, error) {
	data, err := c.Get(ctx, "/ReferenceData/PricingRuleMessageTypes")
	if err != nil {
		return nil, fmt.Errorf("fetching pricing rule message types: %w", err)
	}
	var items []APIPricingRuleMessageType
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("parsing pricing rule message types: %w", err)
	}
	return items, nil
}

type APIPricingRuleType struct {
	Id              *int    `json:"Id"`
	Description     *string `json:"Description"`
	Inactive        bool    `json:"Inactive"`
	CreatedBy       *string `json:"CreatedBy"`
	CreatedDateTime *string `json:"CreatedDateTime"`
	CreateLocation  *string `json:"CreateLocation"`
	UpdatedBy       *string `json:"UpdatedBy"`
	UpdatedDateTime *string `json:"UpdatedDateTime"`
}

func (c *Client) GetPricingRuleTypes(ctx context.Context) ([]APIPricingRuleType, error) {
	data, err := c.Get(ctx, "/ReferenceData/PricingRuleTypes")
	if err != nil {
		return nil, fmt.Errorf("fetching pricing rule types: %w", err)
	}
	var items []APIPricingRuleType
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("parsing pricing rule types: %w", err)
	}
	return items, nil
}

type APISeasonType struct {
	Id              *int                `json:"Id"`
	Description     *string             `json:"Description"`
	Inactive        *bool               `json:"Inactive"`
	BusinessUnit    *APIControlGroupRef `json:"BusinessUnit"`
	CreatedBy       *string             `json:"CreatedBy"`
	CreatedDateTime *string             `json:"CreatedDateTime"`
	CreateLocation  *string             `json:"CreateLocation"`
	UpdatedBy       *string             `json:"UpdatedBy"`
	UpdatedDateTime *string             `json:"UpdatedDateTime"`
}

func (c *Client) GetSeasonTypes(ctx context.Context) ([]APISeasonType, error) {
	data, err := c.Get(ctx, "/ReferenceData/SeasonTypes")
	if err != nil {
		return nil, fmt.Errorf("fetching season types: %w", err)
	}
	var items []APISeasonType
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("parsing season types: %w", err)
	}
	return items, nil
}

type APISubLineItemStatus struct {
	Id              *int    `json:"Id"`
	Description     *string `json:"Description"`
	Inactive        bool    `json:"Inactive"`
	Priority        int     `json:"Priority"`
	StatusCode      *string `json:"StatusCode"`
	CreatedBy       *string `json:"CreatedBy"`
	CreatedDateTime *string `json:"CreatedDateTime"`
	CreateLocation  *string `json:"CreateLocation"`
	UpdatedBy       *string `json:"UpdatedBy"`
	UpdatedDateTime *string `json:"UpdatedDateTime"`
}

func (c *Client) GetSubLineItemStatuses(ctx context.Context) ([]APISubLineItemStatus, error) {
	data, err := c.Get(ctx, "/ReferenceData/SubLineItemStatuses")
	if err != nil {
		return nil, fmt.Errorf("fetching sub-line item statuses: %w", err)
	}
	var items []APISubLineItemStatus
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("parsing sub-line item statuses: %w", err)
	}
	return items, nil
}

type APIUpgradeCategory struct {
	Id              *int    `json:"Id"`
	Description     *string `json:"Description"`
	Inactive        bool    `json:"Inactive"`
	CreatedBy       *string `json:"CreatedBy"`
	CreatedDateTime *string `json:"CreatedDateTime"`
	CreateLocation  *string `json:"CreateLocation"`
	UpdatedBy       *string `json:"UpdatedBy"`
	UpdatedDateTime *string `json:"UpdatedDateTime"`
}

func (c *Client) GetUpgradeCategories(ctx context.Context) ([]APIUpgradeCategory, error) {
	data, err := c.Get(ctx, "/ReferenceData/UpgradeCategories")
	if err != nil {
		return nil, fmt.Errorf("fetching upgrade categories: %w", err)
	}
	var items []APIUpgradeCategory
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("parsing upgrade categories: %w", err)
	}
	return items, nil
}
