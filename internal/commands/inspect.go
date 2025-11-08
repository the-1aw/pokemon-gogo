package commands

import (
	"fmt"
)

const inspectDescription = "<pokemon-name> display info of captured pokemon"

func inspect(c *Config) error {
	if len(c.Args) == 0 {
		return fmt.Errorf("usage: explore %s", catchDescription)
	}
	name := c.Args[0]
	pokemon, ok := c.Pc[name]
	if ok {
		fmt.Printf("Name: %v #%v\n", pokemon.Name, pokemon.Order)
		fmt.Printf("XP: %v\n", pokemon.BaseExperience)
		fmt.Printf("Height: %v cm\n", pokemon.Height*10)
		fmt.Printf("Weight: %v kg\n", pokemon.Weight/10)
	} else {
		fmt.Println("you have not caught that pokemon")
	}
	return nil
}

func init() {
	cmd := Command{
		Description: inspectDescription,
		Run:         inspect,
	}
	registerCommand("inspect", cmd)
}
