package overrider

import (
	"encoding/json"
	"fmt"
	"os"
	"path"

	"github.com/bensadeh/despell/core"
)

func GetOverrides() map[string]core.Icon {
	configFilePath := getPathToConfigFile()

	if !exists(configFilePath) {
		return map[string]core.Icon{}
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
	despellDir := "despell"
	configFile := "overrides.json"

	return path.Join(homeDir, configDir, despellDir, configFile)
}

func unmarshal(data []byte) map[string]core.Icon {
	overrides := make(map[string]core.Icon)

	err := json.Unmarshal(data, &overrides)
	if err != nil {
		fmt.Printf("Error: %s\n", err)

		os.Exit(1)
	}

	return overrides
}
