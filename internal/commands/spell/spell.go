package spell

import (
	"fmt"
	"strings"
)

type Spell struct {
	name        string
	description string
	arg         string
}

func (cmdSpell *Spell) Name() string {
	return cmdSpell.name
}

func (cmdSpell *Spell) Description() string {
	return cmdSpell.description
}

func (cmdSpell *Spell) Args() string {
	return cmdSpell.arg
}

func NewSpell() *Spell {
	return &Spell{
		name:        "spell",
		description: "receives one or more arguments and converts them into letters, separated by spaces",
		arg:         "<command> <argument>",
	}
}

func (cmdSpell *Spell) Execute(args []string) (string, error) {
	lengthArgs := len(args)
	if lengthArgs != 3 {
		return "", fmt.Errorf("spell command receives 1 argument, but given %d", lengthArgs-2)
	}
	word := args[2]
	spelledWord := make([]string, len(word))
	for i, letter := range word {
		spelledWord[i] = string(letter)
	}
	return strings.Join(spelledWord, " "), nil
}
