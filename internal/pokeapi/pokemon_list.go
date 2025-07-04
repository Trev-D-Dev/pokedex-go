package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListPokemon(areaName string) (RespLocationPokemon, error) {
	url := baseURL + "/location-area/" + areaName + "/"

	data, ok := c.cache.Get(url)
	if !ok {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return RespLocationPokemon{}, err
		}

		res, err := c.httpClient.Do(req)
		if err != nil {
			return RespLocationPokemon{}, err
		}
		defer res.Body.Close()

		data, err = io.ReadAll(res.Body)
		if err != nil {
			return RespLocationPokemon{}, err
		}

		c.cache.Add(url, data)
	}

	pokemonRes := RespLocationPokemon{}
	err := json.Unmarshal(data, &pokemonRes)
	if err != nil {
		return RespLocationPokemon{}, err
	}

	return pokemonRes, nil
}
