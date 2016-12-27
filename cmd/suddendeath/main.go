package main

import (
	"fmt"
	"os"

	"github.com/kechako/suddendeath"
)

func printHelp() {
	fmt.Printf("Usage: %s text\n", os.Args[0])
}

func main() {
	args := os.Args[1:]
	var text string
	if len(args) != 1 {
		printHelp()
		os.Exit(1)
	}
	text = args[0]

	fmt.Println(suddendeath.Generate(text))
}
