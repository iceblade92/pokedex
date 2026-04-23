package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) FetchLocations(pageURL *string) (LocationAreaResp, error) {
	url := apiURL
	var resp LocationAreaResp

	if pageURL != nil {
		url = *pageURL
	}

	if data, ok := c.cache.Get(url); ok {
		if err := json.Unmarshal(data, &resp); err != nil {
			return LocationAreaResp{}, err
		}
		return resp, nil
	}

	res, err := http.Get(url)
	if err != nil {
		return resp, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return resp, err
	}

	c.cache.Add(url, body)

	if err := json.Unmarshal(body, &resp); err != nil {
		return resp, err
	}

	return resp, nil
}
