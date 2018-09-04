package settings

import (
	"os"
)

var settings map[string]*store

// SetConfig is the name of the config setting
const SetConfig = "config"

// SetHistory is the name of the history setting
const SetHistory = "history"

const settingSubPath = ".config/pixi"

// KeyLastOpenDirectory is the config key for the last open directory
const KeyLastOpenDirectory = "LastOpenDirectory"

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

// init initializes the config map and file(s)
func init() {
	settings = make(map[string]*store)

	configStoreMap := make(storeMap)
	configStoreMap[KeyLastOpenDirectory] = GetInitialDirectory()
	settings[SetConfig] = new(store)
	settings[SetConfig].filename = "config.json"
	settings[SetConfig].data = configStoreMap
	settings[SetConfig].initStorage()

	settings[SetHistory] = new(store)
	settings[SetHistory].filename = "history.json"
	settings[SetHistory].data = make(storeMap)
	settings[SetHistory].initStorage()
}

// Get returns the config value referred to by key
func Get(set string, key string) interface{} {
	checkIfSetExists(set)
	return settings[set].data[key]
}

// Set a config value to a key
func Set(set string, key string, value interface{}) {
	checkIfSetExists(set)
	settings[set].data[key] = value
	settings[set].writeDataToFile()
}

func checkIfSetExists(set string) {
	_, ok := settings[set]
	if !ok {
		panic("That config set does not exist. (" + set + ")")
	}
}

// GetInitialDirectory gets the start directory
func GetInitialDirectory() string {
	dir, err := os.Getwd()
	checkErr(err)
	return dir
}
