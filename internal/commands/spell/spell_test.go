package spell

import (
	"testing"
)

func TestSpell_Execute(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		want    string
		wantErr bool
	}{{
		name:    "Test 1",
		args:    []string{"cmd/main.go", "spell", "hello"},
		want:    "h e l l o",
		wantErr: false,
	}, {
		name:    "Test 2",
		args:    []string{"cmd/main.go", "spell"},
		want:    "",
		wantErr: true,
	}, {
		name:    "Test 3",
		args:    []string{"cmd/main.go", "spell", "hello", "world"},
		want:    "",
		wantErr: true,
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewSpell().Execute(tt.args)
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
