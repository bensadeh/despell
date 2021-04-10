package main

import (
	"despell/arguments"
	"despell/defaults"
	"fmt"
	"os"
)

func main() {
	arg := arguments.GetArguments()
	mappings := defaults.GetMap()

	result := mappings[arg]

	if result == "" {
		os.Exit(0)
	}

	fmt.Println(mappings[arg])
}
