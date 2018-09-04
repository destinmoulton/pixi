package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"

	homedir "github.com/mitchellh/go-homedir"
)

type storeMap map[string]interface{}

type store struct {
	filename string
	data     storeMap
}

func (s *store) initStorage(){
	if(!s.doesFileExist
}

func loadAndMapifystore() {
	storeFile, err := ioutil.ReadFile(storeFullFilePath)

	checkErr(err)

	errj := json.Unmarshal([]byte(storeFile), &storeMap)
	checkErr(errj)
}

func (s *store) writeStore() {
	data, err := json.Marshal(&s.values)
	checkErr(err)
	ioutil.WriteFile(storeFullFilePath, data, 0666)
}

func (s *store) createFile() {
	if !s.doesBaseDirExist() {
		errD := os.MkdirAll(s.baseDir(), os.ModePerm)
		checkErr(errD)
	}

	initialJSON := make(s.data)

	initialJSON[KeyLastOpenDirectory] = GetInitialDirectory()

	writestoreFile(initialJSON)
}

func (s *store) baseDir() string {
	return path.Join(getHomeDir(), storeSubPath)
}

func (s *store) fullPath() string {
	return path.Join(s.getStoreDir(), s.filename)
}

func (s *store) doesBaseDirExist() bool {
	if _, err := os.Stat(s.baseDir()); os.IsNotExist(err) {
		return false
	}
	return true
}

func (s *store) doesFileExist() bool {
	if _, err := os.Stat(s.fullPath()); os.IsNotExist(err) {
		return false
	}
	return true
}

func getHomeDir() string {
	dir, err := homedir.Dir()
	checkErr(err)
	return dir
}
