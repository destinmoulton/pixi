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
var dirListFileInfo []os.FileInfo
var dirListPrettyNames []string

var visibleList struct {
	maxNumberVisible int
	beginIndex       int
	endIndex         int
	selectedIndex    int
}

var shouldShowHidden = false
var filetypesAllowedVideoFiles = []string{".avi", ".mpeg", ".mkv", ".mp4"}

func initFileList() {

	visibleList.maxNumberVisible = filelistWidgetDims.Height - 2

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
	dirListFileInfo = []os.FileInfo{}

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

func getPrettyList() []string {
	colorifyDirList()
	return dirListPrettyNames
}

// SelectPrevFile switches to the previous file
func SelectPrevFile() {
	if visibleList.selectedIndex > 0 {
		visibleList.selectedIndex--
	} else if visibleList.selectedIndex == 0 {
		if visibleList.beginIndex > 0 {
			// Move the visible "fame" up
			visibleList.beginIndex--
			visibleList.endIndex--
		}
	}
	renderFileList()
}

// SelectNextfile switches to the next file
func SelectNextFile() {
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
	visibleFilesInfo := dirListFileInfo[visibleList.beginIndex : visibleList.endIndex+1]
	selectedFile := visibleFilesInfo[visibleList.selectedIndex]
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
	visibleFilesInfo := dirListFileInfo[visibleList.beginIndex : visibleList.endIndex+1]
	selectedFile := visibleFilesInfo[visibleList.selectedIndex]
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
