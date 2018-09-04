package history

import (
	"path"

	"github.com/rivo/tview"

	"../../settings"
)

type viewedFile map[string]string
type viewedHistory []viewedFile

var uiScreen *tview.Grid
var tableWidget *tview.Table

var history viewedHistory

// StartHistory initializes the history viewer
func StartHistory() {
	loadCurrentHistory()
}

// UI initializes the history ui
func UI() *tview.Grid {
	uiScreen = tview.NewGrid().SetRows(0).SetColumns(0).SetBorders(true)
	tableWidget = tview.NewTable().SetBorders(false)

	uiScreen.AddItem(tableWidget, 0, 0, 1, 1, 0, 0, true)

	return uiScreen
}

func loadCurrentHistory() {
	opened := settings.Get(settings.SetHistory, "opened")

	for _, file := range opened.([]interface{}) {
		// Convert the returned interface (from JSON) into usable map
		tmp := make(viewedFile)
		tmp["filename"] = file.(map[string]interface{})["filename"].(string)
		tmp["path"] = file.(map[string]interface{})["path"].(string)
		history = append(history, tmp)
	}
}

func (h *viewedHistory) Add(fullPath string) {
	_, filename := path.Split(fullPath)

	file := make(viewedFile)
	file["filename"] = filename
	file["path"] = fullPath

	// unshift the new element onto the front of the history
	*h = append(viewedHistory{file}, *h...)

	settings.Set(settings.SetHistory, "opened", h)
}
