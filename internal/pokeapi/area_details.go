package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

type AreaEncounters struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

func (c *Client) AreaDetails(areaName string) (AreaEncounters, error) {
	encounters := AreaEncounters{}

	url := baseURL + "/location-area/" + areaName

	data, ok := c.cache.Get(url)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return AreaEncounters{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return AreaEncounters{}, err
	}

	defer resp.Body.Close()

	details, err := io.ReadAll(resp.Body)
	if err != nil {
		return AreaEncounters{}, err
	}

	if err := json.Unmarshal(details, &encounters); err != nil {
		return AreaEncounters{}, err
	}

	if !ok {
		c.cache.Add(url, data)
	}

	return encounters, nil
}
