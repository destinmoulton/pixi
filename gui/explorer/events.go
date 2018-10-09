package explorer

import (
	"os/exec"
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

	if eventKey.Key() == tcell.KeyF5 || eventKey.Key() == tcell.KeyCtrlR {
		ReRenderExplorer(false)
		return eventKey
	}

	if eventKey.Rune() == 'h' {
		switchToPage("history")
		return eventKey
	}

	if eventKey.Rune() == 's' {
		switchToPage("settingsform")
		return eventKey
	}

	if eventKey.Rune() == '>' || eventKey.Rune() == '.' {
		toggleHidden()
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
// highlights
func navUpDirectory() {
	path := path.Clean(currentPath + "/../")
	oldPath := currentPath

	changeDirectory(path)
	index := getFilelistIndexOf(oldPath)

	// Scroll to top before going to selected
	// so it doesn't appear at top
	uiScrollToTop()
	uiSetSelectedFileIndex(index)
	redrawParent()
}

// navIntoDirectory navigates into the selected directory
func navIntoDirectory() {
	selectedFile := filelist.fullInfo[uiGetSelectedFileIndex()]
	path := path.Join(currentPath, selectedFile.Name())

	if selectedFile.IsDir() {
		changeDirectory(path)
		redrawParent()
	}
}

// performFileAction either opens the dir or opens
// the selected file
func performFileAction() {
	selectedFile := filelist.fullInfo[uiGetSelectedFileIndex()]
	path := path.Join(currentPath, selectedFile.Name())
	if selectedFile.IsDir() {
		navIntoDirectory()
	} else if !selectedFile.IsDir() && player.IsVideoFile(selectedFile.Name()) {
		history.Add(path)
		player.PlayVideo(path)
		ReRenderExplorer(false)
	} else {
		cmd := exec.Command("xdg-open", path)
		err := cmd.Run()
		if err != nil {

		}
	}
}

// toggleHidden enables/disables showing the hidden files (.<filename>)
func toggleHidden() {
	shouldShowHidden = !shouldShowHidden
	ReRenderExplorer(false)
}
