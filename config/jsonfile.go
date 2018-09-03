package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func loadAndMapifyConfig() {
	configFile, err := ioutil.ReadFile(configFullFilePath)

	checkErr(err)

	errj := json.Unmarshal([]byte(configFile), &configMap)
	checkErr(errj)
}

func writeConfigFile(jsonMap tConfigMap) {
	data, err := json.Marshal(&jsonMap)
	checkErr(err)
	ioutil.WriteFile(configFullFilePath, data, 0666)
}

func createConfigFile() {
	if !doesConfigPathExist() {
		errD := os.MkdirAll(configDir, os.ModePerm)
		if errD != nil {
			panic(fmt.Errorf("Fatal error creating config directory: %s ", errD))
		}
	}

	initialJSON := make(tConfigMap)

	initialJSON[KeyLastOpenDirectory] = GetInitialDirectory()

	writeConfigFile(initialJSON)
}

func doesConfigPathExist() bool {
	if _, err := os.Stat(configDir); os.IsNotExist(err) {
		return false
	}
	return true
}

func doesConfigFileExist() bool {
	if _, err := os.Stat(configFullFilePath); os.IsNotExist(err) {
		return false
	}
	return true
}
