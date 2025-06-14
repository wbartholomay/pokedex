package pokeapi

import (
	"encoding/json"
)

// ListLocations -
func (c *Client) ListLocations(pageURL string) (RespShallowLocations, error) {
	url := baseURL + "/location-area"
	if pageURL != "" {
		url = pageURL
	}

	data, err := makeCachedRequest(c, url)
	if err != nil {
		return RespShallowLocations{}, err
	}

	locationsResp := RespShallowLocations{}
	err = json.Unmarshal(data, &locationsResp)
	if err != nil {
		return RespShallowLocations{}, err
	}

	return locationsResp, nil
}
