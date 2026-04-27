package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetLocation(name string) (LocationArea, error) {
	url := baseURL + "/location-area/" + name
	var resp LocationArea
	if val, ok := c.cache.Get(url); ok {
		err := json.Unmarshal(val, &resp)
		if err != nil {
			return LocationArea{}, err
		}
		return resp, err

	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationArea{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationArea{}, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return resp, err
	}

	if err := json.Unmarshal(body, &resp); err != nil {
		return resp, err
	}

	c.cache.Add(url, body)

	return resp, nil

}
