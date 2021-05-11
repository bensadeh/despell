package main

import (
	"despell/arguments"
	"despell/constants"
	"despell/core"
	"despell/overrider"
	"despell/stock"
	"fmt"
)

func main() {
	arg := arguments.GetArguments()
	overrides := overrider.GetOverrides()
	defaults := stock.GetDefaults()

	icon := getIcon(arg, overrides, defaults)

	fmt.Println("#[fg=" + icon.Color + "]" + icon.Text)
}

func getIcon(key string, overrides, defaults map[string]core.Icon) core.Icon {
	unknownCommandIcon := getUnknownCommandIcon(overrides)

	if overridesIcon, ok := overrides[key]; ok {
		return overridesIcon
	}

	if defaultsIcon, ok := defaults[key]; ok {
		return defaultsIcon
	}

	return unknownCommandIcon
}

func getUnknownCommandIcon(overrides map[string]core.Icon) core.Icon {
	if unknownCommandOverride, ok := overrides[constants.UnknownCommandKey]; ok {
		return unknownCommandOverride
	}

	return core.Icon{Text: constants.UnknownCommandIcon, Color: ""}
}
