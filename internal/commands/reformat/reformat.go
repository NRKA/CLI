package reformat

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode"
)

type Reformat struct {
	name        string
	description string
	arg         string
}

func (cmd *Reformat) Name() string {
	return cmd.name
}

func (cmd *Reformat) Description() string {
	return cmd.description
}

func (cmd *Reformat) Args() string {
	return cmd.arg
}

func NewReformat() *Reformat {
	return &Reformat{
		name:        "reformat",
		description: "receives one .txt format argument and formats the data",
		arg:         "<command> <argument>",
	}
}

func (cmd *Reformat) Execute(args []string) (string, error) {
	lengthArgs := len(args)
	if lengthArgs != 3 {
		return "", fmt.Errorf("reformat command receives just 1 argument, but given %d", lengthArgs-2)
	}
	filename := args[2]
	inputFile, err := os.Open(filename)
	if err != nil {
		return "", fmt.Errorf("failed to %v", err)
	}

	err = formatTextProcess(inputFile, os.Stdout)
	if err != nil {
		return "", fmt.Errorf("failed to format text: %v", err)
	}
	err = inputFile.Close()
	if err != nil {
		return "", fmt.Errorf("failed to close file: %v", err)
	}
	return "", nil
}

func formatTextProcess(input io.Reader, output io.Writer) error {
	inputReader := bufio.NewReader(input)
	writer := bufio.NewWriter(output)

	for {
		line, err := inputReader.ReadString('\n')
		isEOF := errors.Is(err, io.EOF)
		if err != nil && !isEOF {
			return fmt.Errorf("failed to read line: %v", err)
		}
		line = strings.TrimSpace(line)
		formattedLine := formatLine(line, isEOF)
		if _, err = writer.WriteString(formattedLine); err != nil {
			return fmt.Errorf("failed to write formatted line: %v", err)
		}
		if isEOF {
			break
		}
	}
	if err := writer.Flush(); err != nil {
		return fmt.Errorf("failed to write buffered data: %v", err)
	}
	return nil
}

func formatLine(line string, isEOF bool) string {
	lineWords := strings.Split(line, " ")
	outputLine := make([]string, len(lineWords))

	for i, word := range lineWords {
		isUpper := unicode.IsUpper(rune(word[0]))
		switch {
		case isUpper:
			if i == 0 {
				outputLine = append(outputLine, word)
				continue
			}
			outputLine = append(outputLine, ". ", word)
		case i == len(lineWords)-1:
			if isEOF {
				outputLine = append(outputLine, " ", word, ".")
				continue
			}
			outputLine = append(outputLine, " ", word, ".\n")
		default:
			outputLine = append(outputLine, " ", word)
		}
	}
	return "\t" + strings.Join(outputLine, "")
}
