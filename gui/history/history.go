package history

import (
	"path"

	"github.com/gdamore/tcell"

	"github.com/rivo/tview"

	"../../settings"
)

type viewedFile map[string]string
type viewedHistory []viewedFile

var history viewedHistory
var redrawParent func()
var uiFrame *tview.Frame
var uiScreen *tview.Grid
var tableWidget *tview.Table

// StartHistory initializes the history viewer
func StartHistory() {
	loadCurrentHistory()
	renderHistory()
}

// UI initializes the history ui and returns the grid
func UI(redraw func()) *tview.Grid {
	redrawParent = redraw
	uiScreen = tview.NewGrid().SetRows(0).SetColumns(0).SetBorders(true)

	tableWidget = tview.NewTable().SetBorders(false)

	uiFrame := tview.NewFrame(tableWidget).AddText("History", true, tview.AlignCenter, tcell.ColorGreen)

	uiScreen.AddItem(uiFrame, 0, 0, 1, 1, 0, 0, true)

	return uiScreen
}

func renderHistory() {
	tableWidget.Clear()
	redrawParent()

	for i, item := range history {
		tableWidget.SetCell(i, 0, tview.NewTableCell(item["filename"]))
	}

	tableWidget.Select(0, 0).SetSelectable(true, false)
	redrawParent()
}

func loadCurrentHistory() {
	opened := settings.Get(settings.SetHistory, "opened")

	if opened != nil {
		for _, file := range opened.([]interface{}) {
			// Convert the returned interface (from JSON) into usable map
			tmp := make(viewedFile)
			tmp["filename"] = file.(map[string]interface{})["filename"].(string)
			tmp["path"] = file.(map[string]interface{})["path"].(string)
			history = append(history, tmp)
		}
	}
}

func clearHistory() {
	history := make(viewedHistory, 0)
	renderHistory()
	settings.Set(settings.SetHistory, "opened", history)
}

func getSelectedFile() viewedFile {
	row, _ := tableWidget.GetSelection()
	return history[row]
}

// Add unshifts(prepends) a file and path onto the history
func Add(fullPath string) {
	_, filename := path.Split(fullPath)

	file := make(viewedFile)
	file["filename"] = filename
	file["path"] = fullPath

	// unshift the new element onto the front of the history
	history = append(viewedHistory{file}, history...)

	renderHistory()

	settings.Set(settings.SetHistory, "opened", history)
}
