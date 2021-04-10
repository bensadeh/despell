package arguments

import "os"

func GetArguments() string {
	argsWithoutProgramName := os.Args[1:]

	return argsWithoutProgramName[0]
}
