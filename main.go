package main

import (
	"despell/arguments"
	"despell/icons"
	"fmt"
)

func main() {
	arg := arguments.GetArguments()
	icon := icons.GetIcon(arg)

	fmt.Println(icon)
}
