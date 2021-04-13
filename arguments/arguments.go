package arguments

import "os"

func GetArguments() string {
	argsWithoutProgramName := os.Args[1:]

	if len(argsWithoutProgramName) == 1 {
		return argsWithoutProgramName[0]
	}

	return ""
}
