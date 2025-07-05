package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) PokemonRetrieve(pokemonName string) (RespPokemonInfo, error) {
	url := baseURL + "/pokemon/" + pokemonName + "/"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespPokemonInfo{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return RespPokemonInfo{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return RespPokemonInfo{}, err
	}

	pokemonRes := RespPokemonInfo{}
	err = json.Unmarshal(data, &pokemonRes)
	if err != nil {
		return RespPokemonInfo{}, err
	}

	return pokemonRes, nil
}
