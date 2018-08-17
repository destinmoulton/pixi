package config

import (
	"fmt"
	"os"
	"path"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var configSubPath = ".config/pixi"
var configDir = ""
var configFilename = "pixi.json"
var configFullFilePath = ""

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

//Init initializes the viper config library
func Init() {
	configDir = path.Join(getHomeDir(), configSubPath)
	configFullFilePath = path.Join(configDir, configFilename)
	fmt.Println(configFullFilePath)
	if !doesConfigFileExist() {
		createConfigFile()
	}

	viper.SetDefault("LastOpenDirectory", getCwd())
	viper.SetConfigType("json")
	viper.SetConfigFile(configFullFilePath)

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s ", err))
	}
}

// Get returns the config value referred to by key
func Get(key string) interface{} {
	return viper.Get(key)
}

// Set a config value to a key
func Set(key string, value interface{}) {
	viper.Set(key, value)

	viper.WriteConfig()
}

func getHomeDir() string {
	dir, err := homedir.Dir()
	checkErr(err)
	return dir
}

func getCwd() string {
	dir, err := os.Getwd()
	checkErr(err)
	return dir
}

func createConfigFile() {
	if !doesConfigPathExist() {
		errD := os.MkdirAll(configDir, os.ModePerm)
		if errD != nil {
			panic(fmt.Errorf("Fatal error creating config directory: %s ", errD))
		}
	}

	f, err := os.Create(configFullFilePath)
	checkErr(err)

	defer f.Close()

	_, err2 := f.WriteString("{}")

	checkErr(err2)
	f.Sync()
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
