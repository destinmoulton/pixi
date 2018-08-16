package explorer

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"sort"
	"strconv"
	"strings"

	"github.com/novalagung/gubrak"
)

var currentPath string
var filelistFullInfo []os.FileInfo
var dirListPrettyNames []string

var filelistVisible struct {
	maxNumberVisible int
	beginIndex       int
	endIndex         int
	selectedIndex    int
}

var shouldShowHidden = false
var filetypesAllowedVideoFiles = []string{".avi", ".mpeg", ".mkv", ".mp4"}

func initFileList() {

	filelistVisible.maxNumberVisible = filelistWidgetDims.Height - 2

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	currentPath = dir
	renderPathBar(currentPath)
	populateDirList()
}

// populateDirList builds the list of elements
// in the selected path
func populateDirList() {
	filelistFullInfo = []os.FileInfo{}

	dirList, err := ioutil.ReadDir(currentPath)

	sort.Sort(SortByLowerCaseFilename(dirList))

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
	filelistFullInfo = append(filelistFullInfo, dirs...)
	filelistFullInfo = append(filelistFullInfo, files...)

	// Setup the visible list
	filelistVisible.selectedIndex = 0
	filelistVisible.beginIndex = 0
	if len(filelistFullInfo) > filelistVisible.maxNumberVisible {
		filelistVisible.endIndex = filelistVisible.maxNumberVisible - 1
	} else {
		filelistVisible.endIndex = len(filelistFullInfo) - 1
	}
}

func getPrettyList() []string {
	colorifyDirList()
	return dirListPrettyNames
}

// SelectPrevFile switches to the previous file
func SelectPrevFile() {
	if filelistVisible.selectedIndex > 0 {
		filelistVisible.selectedIndex--
	} else if filelistVisible.selectedIndex == 0 {
		if filelistVisible.beginIndex > 0 {
			// Move the visible "fame" up
			filelistVisible.beginIndex--
			filelistVisible.endIndex--
		}
	}
	renderFileList()
}

// SelectNextFile switches to the next file
func SelectNextFile() {
	frameEndIndex := filelistVisible.endIndex - filelistVisible.beginIndex
	if filelistVisible.selectedIndex < frameEndIndex {
		filelistVisible.selectedIndex++
	} else if filelistVisible.selectedIndex == frameEndIndex {
		if filelistVisible.endIndex < (len(filelistFullInfo) - 1) {
			// Move the visible "frame" down
			filelistVisible.beginIndex++
			filelistVisible.endIndex++
		}
	}
	renderFileList()
}

// NavUpDirectory navigates up to the parent directory
func NavUpDirectory() {
	path := path.Clean(currentPath + "/../")

	setCurrentPath(path)
	populateDirList()

	renderPathBar(path)
	renderFileList()
}

// NavIntoDirectory navigates into the selected directory
func NavIntoDirectory() {
	visibleFilesInfo := filelistFullInfo[filelistVisible.beginIndex : filelistVisible.endIndex+1]
	selectedFile := visibleFilesInfo[filelistVisible.selectedIndex]
	path := path.Join(currentPath, selectedFile.Name())

	if selectedFile.IsDir() {
		setCurrentPath(path)
		populateDirList()
		renderPathBar(currentPath)
		renderFileList()
	}
}

// PerformFileAction either opens the dir or opens
// the selected file
func PerformFileAction() {
	visibleFilesInfo := filelistFullInfo[filelistVisible.beginIndex : filelistVisible.endIndex+1]
	selectedFile := visibleFilesInfo[filelistVisible.selectedIndex]
	path := path.Join(currentPath, selectedFile.Name())
	if !selectedFile.IsDir() && isVideoFile(selectedFile.Name()) {
		runVideoPlayer(path)
	}
	renderFileList()
}

// ToggleHidden enables/disables showing the hidden files (.<filename>)
func ToggleHidden() {
	shouldShowHidden = !shouldShowHidden
	populateDirList()
	renderFileList()
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
	visibleFilesInfo := filelistFullInfo[filelistVisible.beginIndex : filelistVisible.endIndex+1]
	for idx, file := range visibleFilesInfo {
		fgColor := "fg-white"
		bgColor := ""
		prefix := ""

		if filelistVisible.selectedIndex == idx {
			bgColor = "bg-green"
		}

		if file.IsDir() {
			fgColor = "fg-yellow"
			if filelistVisible.selectedIndex == idx {
				bgColor = "bg-blue"
				fgColor = "fg-white"
			}
		}

		if isVideoFile(file.Name()) {
			prefix = "[>] "
		}

		// Pad the width of the filename
		filename := prefix + file.Name()
		formatString := "[%-" + strconv.Itoa(filelistWidgetDims.Width-3) + "s](%s,%s)"
		prettyName := fmt.Sprintf(formatString, filename, fgColor, bgColor)

		dirListPrettyNames = append(dirListPrettyNames, prettyName)
	}
}

func isVideoFile(filename string) bool {
	res, _ := gubrak.Includes(filetypesAllowedVideoFiles, path.Ext(filename))
	return res
}
