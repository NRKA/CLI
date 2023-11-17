package commands

type CommandInterface interface {
	Name() string
	Description() string
	Args() string
	Execute(args []string) (string, error)
}
