package main

import (
	"despell/arguments"
	"despell/nerdfonts"
	"fmt"
)

func main() {
	arg := arguments.GetArguments()
	icon := nerdfonts.GetIcon(arg)

	fmt.Print(icon)
}
