package commands

import "github.com/the-1aw/pokemon-gogo/internal/pokeapi"

type Config struct {
	Next     *string
	Previous *string
	Pc       map[string]pokeapi.Pokemon
	Args     []string
}

type Command struct {
	Description string
	Run         func(c *Config) error
}

var Registry = map[string]Command{}

func NewConfig() Config {
	return Config{
		Next:     nil,
		Previous: nil,
		Pc:       map[string]pokeapi.Pokemon{},
		Args:     nil,
	}
}

func registerCommand(name string, cmd Command) {
	Registry[name] = cmd
}
