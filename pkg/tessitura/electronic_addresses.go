package tessitura

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
)

func (c *Client) GetElectronicAddresses(ctx context.Context, constituentId int, email string) ([]APIElectronicAddress, error) {
	v := url.Values{}
	v.Set("constituentIds", strconv.Itoa(constituentId))
	v.Set("address", email)
	v.Set("electronicAddressTypeId", "1")
	v.Set("includeAffiliations", "false")
	v.Set("primaryOnly", "false")

	data, err := c.Get(ctx, "/CRM/ElectronicAddresses?"+v.Encode())
	if err != nil {
		return nil, err
	}

	var addresses []APIElectronicAddress
	if err := json.Unmarshal(data, &addresses); err != nil {
		return nil, fmt.Errorf("unable to parse electronic addresses: %w", err)
	}
	return addresses, nil
}

func (c *Client) UpdateElectronicAddress(ctx context.Context, id int, body map[string]interface{}) error {
	_, err := c.Put(ctx, fmt.Sprintf("/CRM/ElectronicAddresses/%d", id), body)
	return err
}
