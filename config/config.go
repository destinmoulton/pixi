package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
