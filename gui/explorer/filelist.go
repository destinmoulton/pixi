package explorer

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"sort"
	"strings"

	"github.com/gdamore/tcell"

	"../../config"
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

var filetypes = map[string]string{
	".avi":  "video",
	".mpeg": "video",
	".mkv":  "video",
	".mp4":  "video",
}

func initFileList() {

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

		renderPathBar(path)
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
			renderPathBar(currentPath)
			renderFileList()
		}
	}
}

// PerformFileAction either opens the dir or opens
// the selected file
func PerformFileAction(r, c int) {
	selectedFile := filelist.fullInfo[r]
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
	for _, file := range filelist.fullInfo {
		fgColor := tcell.ColorWhite
		bgColor := tcell.ColorBlack

		// if filelist.visible.selectedIndex == idx {
		// 	if file.IsDir() {
		// 		fgColor = tview.
		// 		bgColor = "bg-blue"
		// 	} else if isVideoFile(file.Name()) {
		// 		fgColor = "fg-black"
		// 		bgColor = "bg-magenta"
		// 	} else {
		// 		fgColor = "fg-black"
		// 		bgColor = "bg-green"
		// 	}
		if file.IsDir() {
			fgColor = tcell.ColorYellow
		} else if isVideoFile(file.Name()) {
			fgColor = tcell.ColorDarkMagenta
		}

		data := tpretty{file.Name(), fgColor, bgColor}
		filelist.pretty = append(filelist.pretty, data)
	}
}

func isVideoFile(filename string) bool {
	if val, ok := filetypes[path.Ext(filename)]; ok == true && val == "video" {
		return true
	}
	return false
}
