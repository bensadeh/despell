package arguments

import "os"

func Parse() (string, bool) {
	argsWithoutProgramName := os.Args[1:]

	if len(argsWithoutProgramName) == 0 {
		return "", false
	}

	if len(argsWithoutProgramName) == 1 {
		return argsWithoutProgramName[0], false
	}

	if len(argsWithoutProgramName) == 2 {
		firstArgument := argsWithoutProgramName[0]
		secondArgument := argsWithoutProgramName[1]

		if firstArgument == "-c" {
			return secondArgument, true
		}
	}

	return "", false
}
