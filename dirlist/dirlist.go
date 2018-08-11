package dirlist

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"strconv"
)

var dirListFileInfo []os.FileInfo
var dirListPrettyNames []string
var selectedElementIndex int
var currentPath string
var termWidth int
var outputStatusMessage func(string)

// Init initializes the dirlist
func Init(width int, statusMessageHandler func(string)) {
	termWidth = width
	outputStatusMessage = statusMessageHandler

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	currentPath = dir
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
	if selectedElementIndex > 0 {
		selectedElementIndex--
	}
}

// SelectNextElement switches to the next element
// in DirList
func SelectNextElement() {
	if selectedElementIndex < (len(dirListFileInfo) - 1) {
		selectedElementIndex++
	}
}

// NavUpDirectory navigates up to the parent directory
func NavUpDirectory() {
	path := path.Clean(currentPath + "/../")

	setCurrentPath(path)
	PopulateDirList()
}

// PerformFileAction either opens the dir or opens
// the selected file
func PerformFileAction() {
	selectedFile := dirListFileInfo[selectedElementIndex]
	path := path.Join(currentPath, selectedFile.Name())
	if selectedFile.IsDir() {

		outputStatusMessage(path)
		setCurrentPath(path)
		PopulateDirList()
	} else {
		runVideoPlayer(path)
	}
	outputStatusMessage(selectedFile.Name())
}

func runVideoPlayer(selectedFilePath string) {
	cmd := exec.Command("omxplayer", "-b", selectedFilePath)
	err := cmd.Run()

	if err != nil {

	}
}

func setCurrentPath(path string) {
	currentPath = path
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

		formatString := "[%-" + strconv.Itoa(termWidth-3) + "s](%s,%s)"
		prettyName := fmt.Sprintf(formatString, file.Name(), fgColor, bgColor)

		dirListPrettyNames = append(dirListPrettyNames, prettyName)
	}
}
