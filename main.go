package main

import (
	"despell/arguments"
	"despell/defaults"
	"fmt"
	"os"
)

const iconNotFound = "ﲵ"

func main() {
	arg := arguments.GetArguments()
	mappings := defaults.GetMap()

	result := mappings[arg]

	if result == "" {
		fmt.Print(iconNotFound)
		os.Exit(0)
	}

	fmt.Print(mappings[arg])
}
