package help

import (
	ui "github.com/gizak/termui"
)

var helpRows = [][]string{
	[]string{". or h", "Toggle hidden files and folders."},
	[]string{"q or Ctrl+c", "Close pixi"},
}

var helpTable = ui.NewTable()

// Render the help window
func Render() {
	helpTable.Rows = helpRows
	helpTable.FgColor = ui.ColorWhite
	helpTable.BgColor = ui.ColorDefault
	helpTable.Y = 0
	helpTable.X = 0
	helpTable.Width = 70
	helpTable.Height = 10
	ui.Clear()

	ui.Render(helpTable)
}
