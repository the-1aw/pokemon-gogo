package commands

import (
	"fmt"
)

func help(_ *Config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for name, cmd := range Registry {
		fmt.Printf("%v: %v\n", name, cmd.Description)
	}
	return nil
}

func init() {
	cmd := Command{Description: "Displays a help message", Run: help}
	registerCommand("help", cmd)
}
