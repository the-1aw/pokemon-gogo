package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

type NamedAPIResource struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type NamedAPIResourceList struct {
	Count    int                `json:"count"`
	Next     *string            `json:"next"`
	Previous *string            `json:"previous"`
	Results  []NamedAPIResource `json:"results"`
}

const BaseApiUrl = "https://pokeapi.co/api/v2/"

func GetNamedResourceList(url string) (NamedAPIResourceList, error) {
	resourceList := NamedAPIResourceList{}
	queryResult, ok := PokeCache.Get(url)
	if !ok {
		res, err := http.Get(url)
		if err != nil {
			return resourceList, err
		}
		queryResult, err = io.ReadAll(res.Body)
		if err != nil {
			return resourceList, err
		}
		PokeCache.Add(url, queryResult)
	}
	err := json.Unmarshal(queryResult, &resourceList)
	return resourceList, err
}
