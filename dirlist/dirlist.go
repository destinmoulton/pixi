package dirlist

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"strconv"
	"strings"

	"github.com/novalagung/gubrak"

	"../types"
)

var currentPath string
var dirListFileInfo []os.FileInfo
var dirListPrettyNames []string
var pathBar func(string)
var outputStatusMessage func(string)
var visibleList struct {
	maxNumberVisible int
	beginIndex       int
	endIndex         int
	selectedIndex    int
}
var widgetDimensions types.WidgetDimensions
var shouldShowHidden = false

// Init initializes the dirlist
func Init(widgetDim types.WidgetDimensions, pathBarHandler func(string), statusMessageHandler func(string)) {
	widgetDimensions = widgetDim
	pathBar = pathBarHandler
	outputStatusMessage = statusMessageHandler

	visibleList.maxNumberVisible = widgetDim.Height - 2

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	currentPath = dir
	pathBar(currentPath)
	PopulateDirList()
}

// PopulateDirList builds the list of elements
// in the selected path
func PopulateDirList() {
	dirListFileInfo = []os.FileInfo{}

	dirList, err := ioutil.ReadDir(currentPath)
	if err != nil {
		log.Fatal(err)
	}

	dirs := []os.FileInfo{}
	files := []os.FileInfo{}
	for _, file := range dirList {
		filename := file.Name()

		isHiddenFile := strings.HasPrefix(filename, ".")
		if (isHiddenFile && shouldShowHidden) || (!isHiddenFile) {

			if file.IsDir() {
				dirs = append(dirs, file)
			} else {
				files = append(files, file)
			}
		}
	}

	// Directories first, files after
	dirListFileInfo = append(dirListFileInfo, dirs...)
	dirListFileInfo = append(dirListFileInfo, files...)

	// Setup the visible list
	visibleList.selectedIndex = 0
	visibleList.beginIndex = 0
	if len(dirListFileInfo) > visibleList.maxNumberVisible {
		visibleList.endIndex = visibleList.maxNumberVisible - 1
	} else {
		visibleList.endIndex = len(dirListFileInfo) - 1
	}
}

// GetPrettyList gets the current
func GetPrettyList() []string {
	colorifyDirList()
	return dirListPrettyNames
}

// SelectPrevElement switches to the previous element
// in DirList
func SelectPrevElement() {
	if visibleList.selectedIndex > 0 {
		visibleList.selectedIndex--
	} else if visibleList.selectedIndex == 0 {
		if visibleList.beginIndex > 0 {
			// Move the visible "fame" up
			visibleList.beginIndex--
			visibleList.endIndex--
		}
	}

}

// SelectNextElement switches to the next element
// in DirList
func SelectNextElement() {
	frameEndIndex := visibleList.endIndex - visibleList.beginIndex
	if visibleList.selectedIndex < frameEndIndex {
		visibleList.selectedIndex++
	} else if visibleList.selectedIndex == frameEndIndex {
		if visibleList.endIndex < (len(dirListFileInfo) - 1) {
			// Move the visible "frame" down
			visibleList.beginIndex++
			visibleList.endIndex++
		}
	}
}

// NavUpDirectory navigates up to the parent directory
func NavUpDirectory() {
	path := path.Clean(currentPath + "/../")

	setCurrentPath(path)
	pathBar(path)
	PopulateDirList()
}

// PerformFileAction either opens the dir or opens
// the selected file
func PerformFileAction() {
	visibleFilesInfo := dirListFileInfo[visibleList.beginIndex : visibleList.endIndex+1]
	selectedFile := visibleFilesInfo[visibleList.selectedIndex]
	path := path.Join(currentPath, selectedFile.Name())
	if selectedFile.IsDir() {
		setCurrentPath(path)
		PopulateDirList()
		pathBar(currentPath)
	} else {
		runVideoPlayer(path)
	}
}

// ToggleHidden enables/disables showing the hidden files (.<filename>)
func ToggleHidden() {
	shouldShowHidden = !shouldShowHidden
	PopulateDirList()
}

func runVideoPlayer(selectedFilePath string) {
	cmd := exec.Command("xterm", "-e", "omxplayer", "-b", selectedFilePath)
	err := cmd.Run()

	if err != nil {

	}
}

func setCurrentPath(path string) {
	currentPath = path
}

func colorifyDirList() {
	dirListPrettyNames = []string{}
	visibleFilesInfo := dirListFileInfo[visibleList.beginIndex : visibleList.endIndex+1]
	for idx, file := range visibleFilesInfo {
		fgColor := "fg-white"
		bgColor := ""
		prefix := ""

		if visibleList.selectedIndex == idx {
			bgColor = "bg-green"
		}

		if file.IsDir() {
			fgColor = "fg-yellow"
			if visibleList.selectedIndex == idx {
				bgColor = "bg-blue"
				fgColor = "fg-white"
			}
		}
		//outputStatusMessage(path.Ext(file.Name()))
		if res, _ := gubrak.Includes(filetypesAllowedVideoFiles, path.Ext(file.Name())); res == true {
			prefix = "[>] "
		}
		// Pad the width of the filename
		filename := prefix + file.Name()
		formatString := "[%-" + strconv.Itoa(widgetDimensions.Width-3) + "s](%s,%s)"
		prettyName := fmt.Sprintf(formatString, filename, fgColor, bgColor)

		dirListPrettyNames = append(dirListPrettyNames, prettyName)
	}
}
