package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

type resources struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func (c *Client) ListLocations(locationsUrl *string) (resources, error) {
	locations := resources{}

	url := baseURL + "/location-area"
	if locationsUrl != nil {
		url = *locationsUrl
	}

	data, ok := c.cache.Get((url))

	if ok {
		if err := json.Unmarshal(data, &locations); err != nil {
			return resources{}, err
		}
		return locations, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return resources{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return resources{}, err
	}

	defer resp.Body.Close()

	loc, err := io.ReadAll(resp.Body)
	if err != nil {
		return resources{}, err
	}

	if err := json.Unmarshal(loc, &locations); err != nil {
		return resources{}, err
	}

	if !ok {
		c.cache.Add(url, data)
	}

	return locations, nil
}
