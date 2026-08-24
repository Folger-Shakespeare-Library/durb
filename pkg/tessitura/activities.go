package tessitura

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
)

type APISpecialActivity struct {
	Id                      *int       `json:"Id"`
	Constituent             *APIEntity `json:"Constituent"`
	Type                    *APIRef    `json:"Type"`
	Status                  *APIRef    `json:"Status"`
	SpecialActivityDateTime *string    `json:"SpecialActivityDateTime"`
	Notes                   *string    `json:"Notes"`
	CreatedDateTime         *string    `json:"CreatedDateTime"`
	CreatedBy               *string    `json:"CreatedBy"`
	UpdatedDateTime         *string    `json:"UpdatedDateTime"`
	UpdatedBy               *string    `json:"UpdatedBy"`
}

func (c *Client) GetActivities(ctx context.Context, constituentId int, activityTypeId int) ([]APISpecialActivity, error) {
	v := url.Values{}
	v.Set("constituentId", strconv.Itoa(constituentId))
	if activityTypeId > 0 {
		v.Set("activityTypeIds", strconv.Itoa(activityTypeId))
	}

	data, err := c.Get(ctx, "/CRM/SpecialActivities?"+v.Encode())
	if err != nil {
		return nil, err
	}

	var response struct {
		Items []APISpecialActivity `json:"Items"`
	}
	if err := json.Unmarshal(data, &response); err != nil {
		return nil, fmt.Errorf("unable to parse activities: %w", err)
	}
	return response.Items, nil
}

func (c *Client) DeleteActivity(ctx context.Context, id int) error {
	return c.Delete(ctx, fmt.Sprintf("/CRM/SpecialActivities/%d", id))
}

type CreateActivityParams struct {
	ConstituentId  int
	ActivityTypeId int
	StatusId       int
	DateTime       string
	Notes          string
}

func (c *Client) CreateActivity(ctx context.Context, params CreateActivityParams) (*APISpecialActivity, error) {
	body := map[string]interface{}{
		"Constituent":             map[string]int{"Id": params.ConstituentId},
		"Type":                    map[string]int{"Id": params.ActivityTypeId},
		"Status":                  map[string]int{"Id": params.StatusId},
		"SpecialActivityDateTime": params.DateTime,
	}

	if params.Notes != "" {
		body["Notes"] = params.Notes
	}

	data, err := c.Post(ctx, "/CRM/SpecialActivities", body)
	if err != nil {
		return nil, err
	}

	var activity APISpecialActivity
	if err := json.Unmarshal(data, &activity); err != nil {
		return nil, fmt.Errorf("unable to parse activity response: %w", err)
	}
	return &activity, nil
}
