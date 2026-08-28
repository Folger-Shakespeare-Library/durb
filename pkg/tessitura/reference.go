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
