package config

import (
	"os"
	"path"

	"github.com/mitchellh/go-homedir"
)

type tConfigMap map[string]interface{}

var configMap tConfigMap

var configSubPath = ".config/pixi"
var configDir = ""
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

func getHomeDir() string {
	dir, err := homedir.Dir()
	checkErr(err)
	return dir
}

// GetInitialDirectory gets the start directory
func GetInitialDirectory() string {
	dir, err := os.Getwd()
	checkErr(err)
	return dir
}
