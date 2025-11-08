package commands

import (
	"fmt"

	"github.com/the-1aw/pokemon-gogo/internal/pokeapi"
)

const description = "<location-area> display pokemon encounted in the location-area"

func explore(c *Config) error {
	if len(c.Args) == 0 {
		return fmt.Errorf("usage: explore %s", description)
	}
	fmt.Printf("Exploring %s ...\n", c.Args[0])
	la, err := pokeapi.GetLocationArea(c.Args[0])
	if err != nil {
		return err
	}
	if len(la.PokemonEncounters) > 0 {
		fmt.Println("Found Pokemon:")
		for _, pokemon := range la.PokemonEncounters {
			fmt.Println(pokemon.Pokemon.Name)
		}
	}
	return nil
}

func init() {
	cmd := Command{
		Description: description,
		Run:         explore,
	}
	registerCommand("explore", cmd)
}
