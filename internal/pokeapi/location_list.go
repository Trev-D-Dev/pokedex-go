package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/Trev-D-Dev/pokedex-go/internal/types"
)

func (c *Client) ListLocations(pageURL *string) (types.RespShallowLocations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	data, ok := c.cache.Get(url)
	if !ok {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return types.RespShallowLocations{}, err
		}

		res, err := c.httpClient.Do(req)
		if err != nil {
			return types.RespShallowLocations{}, err
		}
		defer res.Body.Close()

		data, err = io.ReadAll(res.Body)
		if err != nil {
			return types.RespShallowLocations{}, err
		}

		c.cache.Add(url, data)
	}

	locationsRes := types.RespShallowLocations{}
	err := json.Unmarshal(data, &locationsRes)
	if err != nil {
		return types.RespShallowLocations{}, err
	}

	return locationsRes, nil
}
