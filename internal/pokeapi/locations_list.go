package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

type resources struct {
	Count    int
	Next     *string
	Previous *string
	Results  []struct {
		Name string
		Url  string
	}
}

func (c *Client) ListLocations(locationsUrl *string) (resources, error) {
	locations := resources{}

	url := baseURL + "/location-area"
	if locationsUrl != nil {
		url = *locationsUrl
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return resources{}, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return resources{}, err
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return resources{}, err
	}

	if err := json.Unmarshal(data, &locations); err != nil {
		return resources{}, err
	}

	return locations, nil
}
