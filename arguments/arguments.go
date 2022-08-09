package arguments

import (
	"github.com/bensadeh/despell/config"
	"os"
)

func GetInputConfig() *config.Settings {
	argsWithoutProgramName := os.Args[1:]
	settings := new(config.Settings)

	if len(argsWithoutProgramName) == 0 {
		settings.Command = ""

		return settings
	}

	if len(argsWithoutProgramName) == 1 {
		settings.Command = argsWithoutProgramName[0]

		return settings
	}

	if len(argsWithoutProgramName) == 2 {
		firstArgument := argsWithoutProgramName[0]
		secondArgument := argsWithoutProgramName[1]

		if firstArgument == "-c" {
			settings.Command = secondArgument
			settings.UseColor = true

			return settings
		}

		if firstArgument == "-e" {
			settings.Command = secondArgument
			settings.UseEmoji = true

			return settings
		}
	}

	return settings
}
