package dirlist

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

var dirListFileInfo []os.FileInfo
var dirListPrettyNames []string
var selectedElementIndex int
var currentPath = "."

// Init initializes the dirlist
func Init() {
	PopulateDirList()
}

// PopulateDirList builds the list of elements
// in the selected path
func PopulateDirList() {
	dirListFileInfo = []os.FileInfo{}
	selectedElementIndex = 0
	dirList, err := ioutil.ReadDir(currentPath)

	if err != nil {
		log.Fatal(err)
	}

	dirs := []os.FileInfo{}
	files := []os.FileInfo{}
	for _, file := range dirList {
		if file.IsDir() {
			dirs = append(dirs, file)
		} else {
			files = append(files, file)
		}
	}

	// Directories first, files after
	dirListFileInfo = append(dirListFileInfo, dirs...)
	dirListFileInfo = append(dirListFileInfo, files...)
}

// GetPrettyList gets the current
func GetPrettyList() []string {
	colorifyDirList()
	return dirListPrettyNames
}

// SelectPrevElement switches to the previous element
// in DirList
func SelectPrevElement() {
	if selectedElementIndex >= 0 {
		selectedElementIndex--
	}
}

// SelectNextElement switches to the next element
// in DirList
func SelectNextElement() {
	if selectedElementIndex <= (len(dirListFileInfo) - 1) {
		selectedElementIndex++
	}
}

func colorifyDirList() {
	dirListPrettyNames = []string{}
	for idx, file := range dirListFileInfo {
		fgColor := "fg-white"
		bgColor := ""

		if selectedElementIndex == idx {
			bgColor = "bg-green"
		}

		if file.IsDir() {

			fgColor = "fg-blue"
			if selectedElementIndex == idx {
				bgColor = "bg-blue"
				fgColor = "fg-white"
			}

		}

		prettyName := fmt.Sprintf("[%-65s](%s,%s)", file.Name(), fgColor, bgColor)

		dirListPrettyNames = append(dirListPrettyNames, prettyName)
	}
}
