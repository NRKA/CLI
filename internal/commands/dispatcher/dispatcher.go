package dispatcher

import (
	"fmt"
	"github.com/NRKA/CLI/internal/commands"
	"github.com/NRKA/CLI/internal/commands/help"
	"github.com/NRKA/CLI/internal/commands/reformat"
	"github.com/NRKA/CLI/internal/commands/spell"
)

type commandProcessor struct {
	commands map[string]commands.CommandInterface
}

func NewCommandProcessor() *commandProcessor {
	processor := &commandProcessor{
		commands: make(map[string]commands.CommandInterface),
	}
	processor.registerCommands()
	return processor
}

func (cmdProcessor *commandProcessor) registerCommands() {
	spellCmd := spell.NewSpell()
	reformatCmd := reformat.NewReformat()

	cmdProcessor.commands[spellCmd.Name()] = spellCmd
	cmdProcessor.commands[reformatCmd.Name()] = reformatCmd

	helpCmd := help.NewHelp(cmdProcessor.commands)
	cmdProcessor.commands[helpCmd.Name()] = helpCmd
}

func (cmdProcessor *commandProcessor) ExecuteCommand(args []string) (string, error) {
	cmd, ok := cmdProcessor.commands[args[1]]
	if !ok {
		return "", fmt.Errorf("unknown command: %s", args[1])
	}
	return cmd.Execute(args)
}

func Handler(args []string) (string, error) {
	cmdProcessor := NewCommandProcessor()

	if len(args) < 2 {
		return "", fmt.Errorf("failed to run application: not enough arguments")
	}

	result, err := cmdProcessor.ExecuteCommand(args)
	if err != nil {
		return "", fmt.Errorf("failed to execute command: %v", err)
	}
	return result, nil
}
