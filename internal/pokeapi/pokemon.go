package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

const PokemonUrl = BaseApiUrl + "pokemon/"

type Pokemon struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	IsDefault      bool   `json:"is_default"`
	Order          int    `json:"order"`
	Weight         int    `json:"weight"`
}

func GetPokemon(name string) (Pokemon, error) {
	pokemon := Pokemon{}
	url := PokemonUrl + name
	queryResult, ok := PokeCache.Get(url)
	if !ok {
		res, err := http.Get(url)
		if err != nil {
			return pokemon, err
		}
		queryResult, err = io.ReadAll(res.Body)
		if err != nil {
			return pokemon, err
		}
		PokeCache.Add(url, queryResult)
	}
	err := json.Unmarshal(queryResult, &pokemon)
	return pokemon, err
}
