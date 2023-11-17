package reformat

import (
	"bytes"
	"io"
	"strings"
	"testing"
)

func Test_formatTextProcess(t *testing.T) {
	tests := []struct {
		name       string
		text       io.Reader
		wantOutput string
		wantErr    bool
	}{{
		name:       "Test 1",
		text:       strings.NewReader("Again it is not my fault You crashed the car\nSo you have to pay the fine"),
		wantOutput: "\tAgain it is not my fault. You crashed the car.\n\tSo you have to pay the fine.",
		wantErr:    false,
	}, {
		name:       "Test 2",
		text:       strings.NewReader("And yet one of the lucky ones People are suffering\nPeople are dying"),
		wantOutput: "\tAnd yet one of the lucky ones. People are suffering.\n\tPeople are dying.",
		wantErr:    false,
	}, {
		name:       "Test 3",
		text:       strings.NewReader("We will not let you get away Right here right now \nThe world is waking up"),
		wantOutput: "\tWe will not let you get away. Right here right now.\n\tThe world is waking up.",
		wantErr:    false,
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := &bytes.Buffer{}
			err := formatTextProcess(tt.text, output)
			if (err != nil) != tt.wantErr {
				t.Errorf("formatTextProcess() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotOutput := output.String(); gotOutput != tt.wantOutput {
				t.Errorf("formatTextProcess() gotOutput = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}

func Test_formatLine(t *testing.T) {
	type args struct {
		line  string
		isEOF bool
	}
	tests := []struct {
		name string
		args args
		want string
	}{{
		name: "Test 1",
		args: args{
			line:  "What did you do all summer when we were working to collect our food Said one of the ants",
			isEOF: false,
		},
		want: "\tWhat did you do all summer when we were working to collect our food. Said one of the ants.\n",
	}, {
		name: "Test 2",
		args: args{
			line:  "And she went helping the ants to carry food to the store They carried on hopping and singing",
			isEOF: true,
		},
		want: "\tAnd she went helping the ants to carry food to the store. They carried on hopping and singing.",
	}, {
		name: "Test 3",
		args: args{
			line:  "I live in a house near the mountains I have two brothers and one sister I was born last",
			isEOF: true,
		},
		want: "\tI live in a house near the mountains. I have two brothers and one sister. I was born last.",
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := formatLine(tt.args.line, tt.args.isEOF); got != tt.want {
				t.Errorf("formatLine() = %v, want %v", got, tt.want)
			}
		})
	}
}
