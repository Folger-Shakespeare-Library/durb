package tessitura

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
)

type APIInterest struct {
	Id              *int              `json:"Id"`
	InterestType    *APIRef           `json:"InterestType"`
	Constituent     *APIEntity        `json:"Constituent"`
	Selected        *bool             `json:"Selected"`
	Weight          *int              `json:"Weight"`
	EditIndicator   bool              `json:"EditIndicator"`
	CreatedDateTime *string           `json:"CreatedDateTime"`
	CreatedBy       *string           `json:"CreatedBy"`
	UpdatedDateTime *string           `json:"UpdatedDateTime"`
	UpdatedBy       *string           `json:"UpdatedBy"`
}

func (c *Client) GetInterests(ctx context.Context, constituentId int) ([]APIInterest, error) {
	v := url.Values{}
	v.Set("constituentId", strconv.Itoa(constituentId))

	data, err := c.Get(ctx, "/CRM/Interests?"+v.Encode())
	if err != nil {
		return nil, err
	}

	var interests []APIInterest
	if err := json.Unmarshal(data, &interests); err != nil {
		return nil, fmt.Errorf("unable to parse interests: %w", err)
	}
	return interests, nil
}

func (c *Client) CreateInterest(ctx context.Context, body map[string]interface{}) error {
	_, err := c.Post(ctx, "/CRM/Interests", body)
	return err
}

func (c *Client) UpdateInterest(ctx context.Context, id int, body map[string]interface{}) error {
	_, err := c.Put(ctx, fmt.Sprintf("/CRM/Interests/%d", id), body)
	return err
}
