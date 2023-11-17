package help

import (
	"github.com/NRKA/CLI/internal/commands"
	"github.com/NRKA/CLI/internal/commands/reformat"
	"github.com/NRKA/CLI/internal/commands/spell"
	"testing"
)

func TestHelp_Execute(t *testing.T) {
	cmd := map[string]commands.CommandInterface{}
	cmd["spell"] = spell.NewSpell()
	cmd["reformat"] = reformat.NewReformat()
	cmd["help"] = NewHelp(cmd)

	tests := []struct {
		name     string
		commands map[string]commands.CommandInterface
		args     []string
		want     string
		wantErr  bool
	}{{
		name:     "Test 1",
		commands: cmd,
		args:     []string{"cmd/main.go", "help"},
		want: "Available commands:\nhelp - shows all available commands. Usage: <command>\n" +
			"reformat - receives one .txt format argument and formats the data. Usage: " +
			"<command> <argument>\n" + "spell - receives one or more arguments and converts " +
			"them into letters, separated by spaces. Usage: <command> <argument>\n",
		wantErr: false,
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := cmd["help"].Execute(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("Execute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Execute() got = %v, want %v", got, tt.want)
			}
		})
	}
}
