package overrider

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
)

func GetOverrides() map[string]string {
	configFilePath := getPathToConfigFile()

	if !exists(configFilePath) {
		return map[string]string{}
	}

	config, err := os.ReadFile(configFilePath)
	if err != nil {
		print("Error: could not read from file")

		os.Exit(1)
	}

	overrides := unmarshal(config)

	return overrides
}

func exists(pathToFile string) bool {
	if _, err := os.Stat(pathToFile); os.IsNotExist(err) {
		return false
	}

	return true
}

func getPathToConfigFile() string {
	homeDir, _ := os.UserHomeDir()
	configDir := ".config"
	magicaDir := "magica"
	configFile := "overrides.json"

	return path.Join(homeDir, configDir, magicaDir, configFile)
}

func unmarshal(data []byte) map[string]string {
	overrides := make(map[string]string)

	err := json.Unmarshal(data, &overrides)
	if err != nil {
		fmt.Printf("Error: %s\n", err)

		os.Exit(1)
	}

	return overrides
}
