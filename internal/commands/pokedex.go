package commands

import (
	"fmt"
)

func pokedex(c *Config) error {
	fmt.Println("You pokedex")
	for name, pokemon := range c.Pc {
		fmt.Printf("- %v #%v\n", name, pokemon.Order)
	}
	return nil
}

func init() {
	cmd := Command{
		Description: "displays all the pokemon you caught",
		Run:         pokedex,
	}
	registerCommand("pokedex", cmd)
}
