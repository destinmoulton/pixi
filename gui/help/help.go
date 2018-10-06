package help

import (
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

var helpItems = [][]string{
	{"up/down arrows", "Select directories/files."},
	{"left arrow", "Navigate to parent directory"},
	{"right arrow", "Navigate into directory"},
	{"enter/return", "Play selected file with omxplayer."},
	{"h", "Show history of played files"},
	{"c", "Clear history of played files (when history open)"},
	{"s", "Open player settings"},
	{">/.", "Toggle viewing hidden files."},
	{"q/Ctrl+c", "Quit pixi"},
	{"ESC/F1", "Close Help - Return to Explorer"},
}

var uiScreen *tview.Grid
var tableWidget *tview.Table

// UI creates the help window
func UI() *tview.Grid {
	uiScreen = tview.NewGrid().SetRows(0).SetColumns(0).SetBorders(true)
	tableWidget = tview.NewTable().SetBorders(false)

	populateTable()

	uiFrame := tview.NewFrame(tableWidget).
		AddText("Help", true, tview.AlignCenter, tcell.ColorDarkMagenta).
		AddText("ESC or F1 to leave Help", false, tview.AlignCenter, tcell.ColorDarkMagenta)

	uiScreen.AddItem(uiFrame, 0, 0, 1, 1, 0, 0, false)

	return uiScreen
}

func populateTable() {
	tableWidget.Clear()
	for i := range helpItems {
		for j, col := range helpItems[i] {
			tableWidget.SetCell(i, j, tview.NewTableCell(col).SetExpansion(1))
		}
	}
}
