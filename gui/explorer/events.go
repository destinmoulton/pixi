package explorer

import (
	"path"

	"../../player"
	"../history"
	"github.com/gdamore/tcell"
)

// HandleEvents dispatches the events for the explorer
func HandleEvents(eventKey *tcell.EventKey, switchToPage func(string)) *tcell.EventKey {
	if eventKey.Key() == tcell.KeyF1 {
		switchToPage("help")
		return eventKey
	}

	if eventKey.Rune() == 'h' {
		switchToPage("history")
		return eventKey
	}

	if eventKey.Key() == tcell.KeyLeft {
		navUpDirectory()
		return nil
	}

	if eventKey.Key() == tcell.KeyRight {
		navIntoDirectory()
		return nil
	}

	if eventKey.Key() == tcell.KeyEnter {
		performFileAction()
		return nil
	}
	return eventKey
}

// navUpDirectory navigates up to the parent directory
func navUpDirectory() {
	path := path.Clean(currentPath + "/../")
	//oldChildPath := currentPath
	changeDirectory(path)

}

// navIntoDirectory navigates into the selected directory
func navIntoDirectory() {
	selectedFile := filelist.fullInfo[getSelectedFileIndex()]
	path := path.Join(currentPath, selectedFile.Name())

	if selectedFile.IsDir() {
		changeDirectory(path)
	}
}

// performFileAction either opens the dir or opens
// the selected file
func performFileAction() {
	selectedFile := filelist.fullInfo[getSelectedFileIndex()]
	path := path.Join(currentPath, selectedFile.Name())
	if selectedFile.IsDir() {
		navIntoDirectory()
	} else if !selectedFile.IsDir() && player.IsVideoFile(selectedFile.Name()) {
		history.Add(path)
		player.PlayVideo(path)
	}
}

// toggleHidden enables/disables showing the hidden files (.<filename>)
func toggleHidden() {
	shouldShowHidden = !shouldShowHidden
	populateDirList()
	renderFileList()
}
