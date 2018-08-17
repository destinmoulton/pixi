package config

import (
	"fmt"
	"os"
	"path"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var configSubPath = ".config/pixi"
var configUserPath = ""
var configFilename = "pixi"
var configFullPath = ""

//Init initializes the viper config library
func Init() {
	configUserPath = path.Join(getHomeDir(), configSubPath)
	configFullPath = path.Join(configUserPath, configFilename)

	if !doesConfigFileExist() {
		createConfigFile()
	}

	viper.SetDefault("LastDirectory", getCwd())

	viper.SetConfigName(configFilename)
	viper.AddConfigPath(configUserPath)

	err := viper.ReadInConfig()

	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s ", err))
	}
}

func getHomeDir() string {
	dir, err := homedir.Dir()
	if err != nil {
		panic(fmt.Errorf("Fatal error getting home directory: %s ", err))
	}
	return dir
}

func getCwd() string {
	dir, errCwd := os.Getwd()
	if errCwd != nil {
		panic(fmt.Errorf("Fatal error getting current working directory: %s ", errCwd))
	}
	return dir
}

func createConfigFile() {
	if !doesConfigPathExist() {
		errD := os.MkdirAll(configUserPath, os.ModePerm)
		if errD != nil {
			panic(fmt.Errorf("Fatal error creating config directory: %s ", errD))
		}
	}

	os.OpenFile(configFullPath, os.O_RDONLY|os.O_CREATE, 0666)
}

func doesConfigPathExist() bool {
	if _, err := os.Stat(configUserPath); os.IsNotExist(err) {
		return false
	}
	return true
}

func doesConfigFileExist() bool {
	if _, err := os.Stat(configFullPath); os.IsNotExist(err) {
		return false
	}
	return true
}
