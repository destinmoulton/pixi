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

	"../../config"
)

var currentPath string

type tvisible struct {
	beginIndex    int
	endIndex      int
	selectedIndex int
}

var filelist struct {
	visible  tvisible
	fullInfo []os.FileInfo
	pretty   []string
}

// visibleHistory stores the visible/selected history
var visibleHistory = make(map[string]tvisible)

var maxNumberVisible int
var shouldShowHidden = false

var filetypes = map[string]string{
	".avi":  "video",
	".mpeg": "video",
	".mkv":  "video",
	".mp4":  "video",
}

func initFileList() {

	maxNumberVisible = filelistWidgetDims.Height - 2

	initialPath := config.Get(config.KeyLastOpenDirectory).(string)
	if !doesDirectoryExist(initialPath) || !isDirectoryReadable(initialPath) {
		initialPath = config.GetInitialDirectory()
	}
	setCurrentPath(initialPath)

	renderPathBar(currentPath)
	populateDirList()
}

func doesDirectoryExist(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

// populateDirList builds the list of elements
// in the selected path
func populateDirList() {
	filelist.fullInfo = []os.FileInfo{}

	dirList, err := ioutil.ReadDir(currentPath)

	if err != nil {
		log.Fatal(err)
	}

	sort.Sort(SortByLowerCaseFilename(dirList))

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
	filelist.fullInfo = append(filelist.fullInfo, dirs...)
	filelist.fullInfo = append(filelist.fullInfo, files...)

	priorVisible, hasHistory := visibleHistory[currentPath]
	if hasHistory && priorVisible.selectedIndex <= len(filelist.fullInfo)-1 {
		filelist.visible = priorVisible
	} else {
		// Setup the visible list
		filelist.visible.selectedIndex = 0
		filelist.visible.beginIndex = 0
		if len(filelist.fullInfo) > maxNumberVisible {
			filelist.visible.endIndex = maxNumberVisible - 1
		} else {
			filelist.visible.endIndex = len(filelist.fullInfo) - 1
		}
		visibleHistory[currentPath] = filelist.visible
	}
}

func getPrettyList() []string {
	colorifyDirList()
	return filelist.pretty
}

// SelectPrevFile switches to the previous file
func SelectPrevFile() {
	if filelist.visible.selectedIndex > 0 {
		filelist.visible.selectedIndex--
	} else if filelist.visible.selectedIndex == 0 {
		if filelist.visible.beginIndex > 0 {
			// Move the visible "fame" up
			filelist.visible.beginIndex--
			filelist.visible.endIndex--
		}
	}
	visibleHistory[currentPath] = filelist.visible
	renderFileList()
}

// SelectNextFile switches to the next file
func SelectNextFile() {
	frameEndIndex := filelist.visible.endIndex - filelist.visible.beginIndex
	if filelist.visible.selectedIndex < frameEndIndex {
		filelist.visible.selectedIndex++
	} else if filelist.visible.selectedIndex == frameEndIndex {
		if filelist.visible.endIndex < (len(filelist.fullInfo) - 1) {
			// Move the visible "frame" down
			filelist.visible.beginIndex++
			filelist.visible.endIndex++
		}
	}
	visibleHistory[currentPath] = filelist.visible
	renderFileList()
}

// NavUpDirectory navigates up to the parent directory
func NavUpDirectory() {
	path := path.Clean(currentPath + "/../")
	if isDirectoryReadable(path) {
		setCurrentPath(path)
		populateDirList()

		renderPathBar(path)
		renderFileList()
	}
}

// NavIntoDirectory navigates into the selected directory
func NavIntoDirectory() {
	visibleFilesInfo := filelist.fullInfo[filelist.visible.beginIndex : filelist.visible.endIndex+1]
	selectedFile := visibleFilesInfo[filelist.visible.selectedIndex]
	path := path.Join(currentPath, selectedFile.Name())

	if selectedFile.IsDir() {

		if isDirectoryReadable(path) {
			setCurrentPath(path)
			populateDirList()
			renderPathBar(currentPath)
			renderFileList()
		}
	}
}

// PerformFileAction either opens the dir or opens
// the selected file
func PerformFileAction() {
	visibleFilesInfo := filelist.fullInfo[filelist.visible.beginIndex : filelist.visible.endIndex+1]
	selectedFile := visibleFilesInfo[filelist.visible.selectedIndex]
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

func isDirectoryReadable(dir string) bool {
	if _, err := ioutil.ReadDir(dir); err != nil {
		return false
	}
	return true
}

func runVideoPlayer(selectedFilePath string) {
	cmd := exec.Command("xterm", "-e", "omxplayer", "-b", selectedFilePath)
	err := cmd.Run()

	if err != nil {

	}
}

func setCurrentPath(path string) {
	currentPath = path

	config.Set(config.KeyLastOpenDirectory, path)
}

func colorifyDirList() {
	filelist.pretty = []string{}
	visibleFilesInfo := filelist.fullInfo[filelist.visible.beginIndex : filelist.visible.endIndex+1]
	for idx, file := range visibleFilesInfo {
		fgColor := "fg-white"
		bgColor := ""

		if filelist.visible.selectedIndex == idx {
			if file.IsDir() {
				fgColor = "fg-white"
				bgColor = "bg-blue"
			} else if isVideoFile(file.Name()) {
				fgColor = "fg-black"
				bgColor = "bg-magenta"
			} else {
				fgColor = "fg-black"
				bgColor = "bg-green"
			}
		} else if file.IsDir() {
			fgColor = "fg-yellow"
		} else {
			if isVideoFile(file.Name()) {
				fgColor = "fg-magenta"
				bgColor = ""
			}
		}

		// Pad the width of the filename
		formatString := "[%-" + strconv.Itoa(filelistWidgetDims.Width-3) + "s](%s,%s)"
		prettyName := fmt.Sprintf(formatString, file.Name(), fgColor, bgColor)

		filelist.pretty = append(filelist.pretty, prettyName)
	}
}

func isVideoFile(filename string) bool {
	if val, ok := filetypes[path.Ext(filename)]; ok == true && val == "video" {
		return true
	}
	return false
}
