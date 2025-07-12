package pokeapi

import (
	"encoding/json"
	//"fmt"
	"io"
	"net/http"

	"github.com/Trev-D-Dev/pokedex-go/internal/types"
)

func (c *Client) PokemonRetrieve(pokemonName string) (types.RespPokemonInfo, error) {
	url := baseURL + "/pokemon/" + pokemonName + "/"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return types.RespPokemonInfo{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return types.RespPokemonInfo{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return types.RespPokemonInfo{}, err
	}

	pokemonRes := types.RespPokemonInfo{}
	err = json.Unmarshal(data, &pokemonRes)
	if err != nil {
		return types.RespPokemonInfo{}, err
	}
	/*
		fmt.Printf("Pokemon Info: %v\n", pokemonRes)

		fmt.Printf("Name: %s\n", pokemonRes.Name)
		fmt.Printf("Base Exp: %v\n", pokemonRes.BaseExp)

		fmt.Println()

		// Prints out all abilities
		fmt.Println("Abilities:")
		for _, ability := range pokemonRes.Abilities {
			fmt.Printf("- %s\n", ability.Ability.Name)
		}

		fmt.Println()

		// Prints all forms
		fmt.Println("Forms:")
		for _, form := range pokemonRes.Forms {
			fmt.Printf("- %s\n", form.Name)
		}

		fmt.Println()

		// Prints all moves
		fmt.Println("Moves:")
		for _, move := range pokemonRes.Moves {
			fmt.Printf("- %s\n", move.Move.Name)
		}

		fmt.Println()

		// Prints all stats
		fmt.Println("Stats:")
		for _, stat := range pokemonRes.Stats {
			fmt.Printf("- %s: %v\n", stat.Stat.Name, stat.BaseStat)
		}

		fmt.Println()

		// Prints all types
		fmt.Println("Types:")
		for _, pType := range pokemonRes.Types {
			fmt.Printf("- %s\n", pType.Type.Name)
		}*/

	return pokemonRes, nil
}
