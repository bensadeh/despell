package main

import (
	"despell/arguments"
	"despell/constants"
	"despell/overrider"
	"despell/stock"
	"fmt"
)

func main() {
	arg := arguments.GetArguments()
	overrides := overrider.GetOverrides()
	defaults := stock.GetDefaults()

	icon := getIcon(arg, overrides, defaults)

	fmt.Println(icon)
}

func getIcon(key string, overrides, defaults map[string]string) string {
	overridesIcon := overrides[key]
	defaultsIcon := defaults[key]
	unknownCommandIcon := getUnknownCommandIcon(overrides)

	if overridesIcon != "" {
		return overridesIcon
	}

	if defaultsIcon != "" {
		return defaultsIcon
	}

	return unknownCommandIcon
}

func getUnknownCommandIcon(overrides map[string]string) string {
	unknownCommandOverride := overrides[constants.UnknownCommandKey]

	if unknownCommandOverride == "" {
		return constants.UnknownCommandIcon
	}

	return unknownCommandOverride
}
