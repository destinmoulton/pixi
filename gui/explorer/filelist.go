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

func changeDirectory(path string) {
	if isDirectoryReadable(path) {
		setCurrentPath(path)
		populateDirList()

		setPathWidgetText(path)
		renderFileList()
	}
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

func colorifyDirList() {
	history.LoadCurrentHistory()
	for _, file := range filelist.fullInfo {
		fgColor := tcell.ColorWhite
		bgColor := tcell.ColorBlack
		path := path.Join(currentPath, file.Name())
		if file.IsDir() {
			fgColor = tcell.ColorYellow
		} else if history.IsFileInHistory(path) {
			fgColor = tcell.ColorDarkMagenta
		} else if player.IsVideoFile(file.Name()) {
			fgColor = tcell.ColorGreenYellow
		}

		data := tpretty{file.Name(), fgColor, bgColor}
		filelist.pretty = append(filelist.pretty, data)
	}
}

func setCurrentPath(path string) {
	currentPath = path

	settings.Set(settings.SetConfig, settings.KeyLastOpenDirectory, path)
}

func isDirectoryReadable(dir string) bool {
	if _, err := ioutil.ReadDir(dir); err != nil {
		return false
	}
	return true
}

func doesDirectoryExist(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

func getFilelistIndexOf(pathToFind string) int {
	for i, file := range filelist.fullInfo {
		if path.Join(currentPath, file.Name()) == pathToFind {
			return i
		}
	}
	return -1
}
