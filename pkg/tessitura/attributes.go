package tessitura

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
)

type APIAttribute struct {
	Id              *int       `json:"Id"`
	Keyword         *APIRef    `json:"Keyword"`
	Constituent     *APIEntity `json:"Constituent"`
	Value           *string    `json:"Value"`
	EditIndicator   bool       `json:"EditIndicator"`
	CreatedDateTime *string    `json:"CreatedDateTime"`
	CreatedBy       *string    `json:"CreatedBy"`
	UpdatedDateTime *string    `json:"UpdatedDateTime"`
	UpdatedBy       *string    `json:"UpdatedBy"`
}

func (c *Client) GetAttributes(ctx context.Context, constituentId int, keywordId int) ([]APIAttribute, error) {
	v := url.Values{}
	v.Set("constituentId", strconv.Itoa(constituentId))
	if keywordId != 0 {
		v.Set("keywordIds", strconv.Itoa(keywordId))
	}

	data, err := c.Get(ctx, "/CRM/Attributes?"+v.Encode())
	if err != nil {
		return nil, err
	}

	var attrs []APIAttribute
	if err := json.Unmarshal(data, &attrs); err != nil {
		return nil, fmt.Errorf("unable to parse attributes: %w", err)
	}
	return attrs, nil
}

func (c *Client) CreateAttribute(ctx context.Context, body map[string]interface{}) error {
	_, err := c.Post(ctx, "/CRM/Attributes", body)
	return err
}

func (c *Client) UpdateAttribute(ctx context.Context, id int, body map[string]interface{}) error {
	_, err := c.Put(ctx, fmt.Sprintf("/CRM/Attributes/%d", id), body)
	return err
}

func (c *Client) DeleteAttribute(ctx context.Context, id int) error {
	return c.Delete(ctx, fmt.Sprintf("/CRM/Attributes/%d", id))
}
