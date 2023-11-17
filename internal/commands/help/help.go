package help

import (
	"fmt"
	"github.com/NRKA/CLI/internal/commands"
	"sort"
	"strings"
)

type Help struct {
	name           string
	description    string
	arg            string
	commands       map[string]commands.CommandInterface
	sortedCommands []string
}

func (cmd *Help) Name() string {
	return cmd.name
}

func (cmd *Help) Description() string {
	return cmd.description
}

func (cmd *Help) Args() string {
	return cmd.arg
}

func (cmd *Help) Execute(args []string) (string, error) {
	if len(args) > 2 {
		return "", fmt.Errorf("'help' command doesn't receive argument, but given: %s", args[2:])
	}

	output := strings.Builder{}
	output.WriteString("Available commands:\n")
	for _, command := range cmd.sortedCommands {
		description := cmd.commands[command].Description()
		commandArgs := cmd.commands[command].Args()
		commandInfo := fmt.Sprintf("%s - %s. Usage: %s\n", command, description, commandArgs)
		output.WriteString(commandInfo)
	}
	return output.String(), nil
}

func NewHelp(commands map[string]commands.CommandInterface) *Help {
	help := &Help{
		name:           "help",
		description:    "shows all available commands",
		arg:            "<command>",
		commands:       commands,
		sortedCommands: make([]string, 0, len(commands)+1),
	}

	for _, command := range commands {
		help.sortedCommands = append(help.sortedCommands, command.Name())
	}
	help.sortedCommands = append(help.sortedCommands, help.Name())
	sort.Strings(help.sortedCommands)

	return help
}
