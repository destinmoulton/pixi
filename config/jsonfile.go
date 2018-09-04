package settings

import (
	"encoding/json"
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

func (s *store) initStorage() {
	if !s.doesFileExist() {
		s.createFile()
	}
	s.loadAndMapifystore()
}

func (s *store) loadAndMapifystore() {
	storeFile, err := ioutil.ReadFile(s.fullPath())

	checkErr(err)

	errj := json.Unmarshal([]byte(storeFile), &s.data)
	checkErr(errj)
}

func (s *store) writeFile() {
	json, err := json.Marshal(&s.data)
	checkErr(err)
	ioutil.WriteFile(s.fullPath(), json, 0666)
}

func (s *store) createFile() {
	if !s.doesBaseDirExist() {
		errD := os.MkdirAll(s.baseDir(), os.ModePerm)
		checkErr(errD)
	}

	initialJSON := make(s.data)

	initialJSON[KeyLastOpenDirectory] = GetInitialDirectory()

	writeFile(initialJSON)
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
