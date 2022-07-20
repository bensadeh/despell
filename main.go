package main

import (
	"fmt"

	"github.com/bensadeh/despell/arguments"
	"github.com/bensadeh/despell/core"
	"github.com/bensadeh/despell/overrider"
	"github.com/bensadeh/despell/stock"
	"github.com/bensadeh/despell/unknown"
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
	if unknownCommandOverride, ok := overrides[unknown.UnknownCommandKey]; ok {
		return unknownCommandOverride
	}

	return unknown.GetUnknownCommandIcon()
}
