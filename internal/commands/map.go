package commands

import (
	"fmt"

	"github.com/the-1aw/pokemon-gogo/internal/pokeapi"
)

func map_(c *Config) error {
	url := pokeapi.LocationAreaUrl + "?offset=0&limit=20"
	if c.Next != nil {
		url = *c.Next
	}
	rl, err := pokeapi.GetNamedResourceList(url)
	if err != nil {
		return err
	}
	c.Next = rl.Next
	c.Previous = rl.Previous
	for _, locArea := range rl.Results {
		fmt.Println(locArea.Name)
	}
	return nil
}

func mapb(c *Config) error {
	if c.Previous == nil {
		fmt.Println("you're on the first page")
		return nil
	}
	rl, err := pokeapi.GetNamedResourceList(*c.Previous)
	if err != nil {
		return err
	}
	c.Next = rl.Next
	c.Previous = rl.Previous
	for _, locArea := range rl.Results {
		fmt.Println(locArea.Name)
	}
	return nil
}

func init() {
	mapCmd := Command{
		Description: "Display the next 20 locations",
		Run:         map_,
	}
	mapbCmd := Command{
		Description: "Display the previous 20 locations",
		Run:         mapb,
	}
	registerCommand("map", mapCmd)
	registerCommand("mapb", mapbCmd)
}
