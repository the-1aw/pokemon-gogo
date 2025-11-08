package commands

import (
	"fmt"
	"os"
)

func exit(_ *Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func init() {
	cmd := Command{
		Description: "Exit the Pokedex",
		Run:         exit,
	}
	registerCommand("exit", cmd)
}
