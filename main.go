package main

import (
	"despell/arguments"
	"despell/icons"
	"fmt"

	"github.com/spf13/cobra/cobra/cmd"
)

func main() {
	_ = cmd.Execute()

	arg := arguments.GetArguments()
	icon := icons.GetIcon(arg)

	fmt.Println(icon)
}
