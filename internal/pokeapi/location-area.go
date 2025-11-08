package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

const LocationAreaUrl = BaseApiUrl + "location-area/"

type LocationArea struct {
	Id                int    `json:"id"`
	Name              string `json:"name"`
	GameIndex         int    `json:"game_index"`
	PokemonEncounters []struct {
		Pokemon NamedAPIResource `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

func GetLocationArea(name string) (LocationArea, error) {
	locationArea := LocationArea{}
	url := LocationAreaUrl + name
	queryResult, ok := PokeCache.Get(url)
	if !ok {
		res, err := http.Get(url)
		if err != nil {
			return locationArea, err
		}
		queryResult, err = io.ReadAll(res.Body)
		if err != nil {
			return locationArea, err
		}
		PokeCache.Add(url, queryResult)
	}
	err := json.Unmarshal(queryResult, &locationArea)
	return locationArea, err
}
