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
	command, isColors := arguments.Parse()
	overrides := overrider.GetOverrides()
	defaults := stock.GetDefaults()

	icon := getIcon(command, overrides, defaults)
	output := format(isColors, icon)

	fmt.Println(output)
}

func format(isColors bool, icon core.Icon) string {
	if isColors {
		return "#[fg=" + icon.Color + "]" + icon.Icon
	}

	return icon.Icon
}

func getIcon(key string, overrides, defaults map[string]core.Icon) core.Icon {
	if overridesIcon, ok := overrides[key]; ok {
		return overridesIcon
	}

	if defaultsIcon, ok := defaults[key]; ok {
		return defaultsIcon
	}

	return getUnknownCommandIcon(overrides)
}

func getUnknownCommandIcon(overrides map[string]core.Icon) core.Icon {
	if unknownCommandOverride, ok := overrides[constants.UnknownCommandKey]; ok {
		return unknownCommandOverride
	}

	return constants.GetUnknownCommandIcon()
}
