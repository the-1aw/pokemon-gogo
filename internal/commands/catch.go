package commands

import (
	"fmt"
	"math/rand"

	"github.com/the-1aw/pokemon-gogo/internal/pokeapi"
)

const catchDescription = "<pokemon-name> throw a pokeball at <pokemon-name>"

func catch(c *Config) error {
	if len(c.Args) == 0 {
		return fmt.Errorf("usage: explore %s", catchDescription)
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", c.Args[0])
	pokemon, err := pokeapi.GetPokemon(c.Args[0])
	if err != nil {
		return err
	}
	catchRate := 0.85
	if (float64(rand.Intn(100)) / 100.) < catchRate {
		fmt.Printf("%s was caught!\n", pokemon.Name)
	} else {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	}
	return nil
}

func init() {
	cmd := Command{
		Description: catchDescription,
		Run:         catch,
	}
	registerCommand("catch", cmd)
}
