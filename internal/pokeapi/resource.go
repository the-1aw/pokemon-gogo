package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/the-1aw/pokemon-gogo/internal/pokecache"
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

var cache = pokecache.NewCache(50 * time.Second)

func GetNamedResourceList(url string) (NamedAPIResourceList, error) {
	resourceList := NamedAPIResourceList{}
	queryResult, ok := cache.Get(url)
	if !ok {
		res, err := http.Get(url)
		if err != nil {
			return resourceList, err
		}
		queryResult, err = io.ReadAll(res.Body)
		if err != nil {
			return resourceList, err
		}
		cache.Add(url, queryResult)
	}
	if err := json.Unmarshal(queryResult, &resourceList); err != nil {
		return resourceList, err
	}
	return resourceList, nil

}
