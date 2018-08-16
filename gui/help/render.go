package help

import (
	ui "github.com/gizak/termui"
)

var helpLines = []string{
	". or h         Toggle hidden files and folders.",
	"q or Ctrl+c    Close pixi",
}

var helpList = ui.NewList()

// Render the help window
func Render() {
	helpList.Items = helpLines
	helpList.BorderLabel = "Help"

	ui.Clear()
	ui.Body.Rows = ui.Body.Rows[:0]
	ui.Body.AddRows(
		ui.NewRow(ui.NewCol(12, 0, helpList)))

	ui.Body.Align()
	ui.Render(ui.Body)
}
