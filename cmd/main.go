package main

import (
	"fmt"
	"github.com/NRKA/CLI/internal/commands/dispatcher"
	"log"
	"os"
)

func main() {
	args := os.Args
	result, err := dispatcher.Handler(args)
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}
	fmt.Println(result)
}
