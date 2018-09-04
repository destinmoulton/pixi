package config

import (
	"os"
	"path"
)

var configSubPath = ".config/pixi"

var configFilename = "pixi.json"
var configFullFilePath = ""

// KeyLastOpenDirectory is the config key for the last open directory
const KeyLastOpenDirectory = "LastOpenDirectory"

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

//Init initializes the config file
func Init() {
	var config = new(store)
	configDir = path.Join(getHomeDir(), configSubPath)
	configFullFilePath = path.Join(configDir, configFilename)

	if !doesConfigFileExist() {
		createConfigFile()
	}

	loadAndMapifyConfig()
}

// Get returns the config value referred to by key
func Get(key string) interface{} {
	return configMap[key]
}

// Set a config value to a key
func Set(key string, value interface{}) {
	configMap[key] = value

	writeConfigFile(configMap)
}

// GetInitialDirectory gets the start directory
func GetInitialDirectory() string {
	dir, err := os.Getwd()
	checkErr(err)
	return dir
}
