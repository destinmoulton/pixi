package settings

import (
	"os"
)

// Settings is a storage for the main settings
var Settings store

// KeyLastOpenDirectory is the config key for the last open directory
const KeyLastOpenDirectory = "LastOpenDirectory"

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

//Init initializes the config file
func Init() {
	Settings.filename = "pixi.json"

	Settings.initStorage()
}

// Get returns the config value referred to by key
func Get(key string) interface{} {
	return Settings.data[key]
}

// Set a config value to a key
func Set(key string, value interface{}) {
	Settings.data[key] = value

	Settings.writeFile()
}

// GetInitialDirectory gets the start directory
func GetInitialDirectory() string {
	dir, err := os.Getwd()
	checkErr(err)
	return dir
}
