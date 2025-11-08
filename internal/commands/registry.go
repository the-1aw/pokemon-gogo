package commands

type Command struct {
	Description string
	Run         func() error
}

var Registry = map[string]Command{}

func registerCommand(name string, cmd Command) {
	Registry[name] = cmd
}
