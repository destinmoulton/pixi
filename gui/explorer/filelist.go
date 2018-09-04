package explorer

import (
	"io/ioutil"
	"log"
	"os"
	"path"
	"sort"
	"strings"

	"github.com/gdamore/tcell"

	"../../player"
	"../../settings"
	"../history"
)

var currentPath string

type tpretty struct {
	filename string
	fgColor  tcell.Color
	bgColor  tcell.Color
}

var filelist struct {
	fullInfo []os.FileInfo
	pretty   []tpretty
}

var shouldShowHidden = false

func initFileList() {

	initialPath := settings.Get(settings.SetConfig, settings.KeyLastOpenDirectory).(string)
	if !doesDirectoryExist(initialPath) || !isDirectoryReadable(initialPath) {
		initialPath = settings.GetInitialDirectory()
	}
	setCurrentPath(initialPath)

	setPathWidgetText(initialPath)
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
	filelist.pretty = []tpretty{}

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
}

func getPrettyList() []tpretty {
	colorifyDirList()
	return filelist.pretty
}

// NavUpDirectory navigates up to the parent directory
func NavUpDirectory() {
	path := path.Clean(currentPath + "/../")
	if isDirectoryReadable(path) {
		setCurrentPath(path)
		populateDirList()

		setPathWidgetText(path)
		renderFileList()
	}
}

// NavIntoDirectory navigates into the selected directory
func NavIntoDirectory() {
	selectedFile := filelist.fullInfo[getSelectedFileIndex()]
	path := path.Join(currentPath, selectedFile.Name())

	if selectedFile.IsDir() {

		if isDirectoryReadable(path) {
			setCurrentPath(path)
			populateDirList()
			setPathWidgetText(path)
			renderFileList()
		}
	}
}

// PerformFileAction either opens the dir or opens
// the selected file
func PerformFileAction() {
	selectedFile := filelist.fullInfo[getSelectedFileIndex()]
	path := path.Join(currentPath, selectedFile.Name())
	if !selectedFile.IsDir() && player.IsVideoFile(selectedFile.Name()) {
		history.Add(path)
		player.PlayVideo(path)
	}
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

func setCurrentPath(path string) {
	currentPath = path

	settings.Set(settings.SetConfig, settings.KeyLastOpenDirectory, path)
}

func colorifyDirList() {
	for _, file := range filelist.fullInfo {
		fgColor := tcell.ColorWhite
		bgColor := tcell.ColorBlack

		if file.IsDir() {
			fgColor = tcell.ColorYellow
		} else if player.IsVideoFile(file.Name()) {
			fgColor = tcell.ColorDarkMagenta
		}

		data := tpretty{file.Name(), fgColor, bgColor}
		filelist.pretty = append(filelist.pretty, data)
	}
}
