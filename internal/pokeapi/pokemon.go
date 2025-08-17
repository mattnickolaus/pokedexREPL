package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemonDetails(pokemonName string) (PokemonDetails, error) {
	url := baseURL + "/pokemon/" + pokemonName

	if val, ok := c.cache.Get(url); ok {
		// Use the cache
		var pokemon PokemonDetails
		if err := json.Unmarshal(val, &pokemon); err != nil {
			return PokemonDetails{}, err
		}
		return pokemon, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokemonDetails{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonDetails{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return PokemonDetails{}, err
	}

	if res.StatusCode > 299 {
		fmt.Printf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
		return PokemonDetails{}, err
	}

	c.cache.Add(url, body)

	var pokemon PokemonDetails
	if err := json.Unmarshal(body, &pokemon); err != nil {
		fmt.Println(err)
		return PokemonDetails{}, err
	}

	return pokemon, nil
}
