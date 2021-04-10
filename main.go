package main

import (
	"despell/arguments"
	"despell/defaults"
	"fmt"
)

func main() {
	arg := arguments.GetArguments()
	icon := defaults.GetIcon(arg)

	fmt.Print(icon)
}
