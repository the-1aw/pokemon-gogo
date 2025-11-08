package commands

type Config struct {
	Next     *string
	Previous *string
}

type Command struct {
	Description string
	Run         func(c *Config) error
}

var Registry = map[string]Command{}

func registerCommand(name string, cmd Command) {
	Registry[name] = cmd
}
