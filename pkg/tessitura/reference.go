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
