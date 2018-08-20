package explorer

import (
	ui "github.com/gizak/termui"

	"../../types"
)

// fileListWidget termui widget List
var fileListWidget = ui.NewList()

// statusBarWidget termui widget Status Bar
var statusBarWidget = ui.NewPar("[F1](fg-green) Help | [h](fg-green) Toggle Hidden Files")
var timeWidget = ui.NewPar("Time")
var filelistWidgetDims types.WidgetDimensions

// InitRender initializes the File Explorer
func InitRender() {
	filelistWidgetDims.Width = ui.TermWidth()
	filelistWidgetDims.Height = ui.TermHeight() - 3

	setupExplorerGUI()

	initFileList()

	renderFileList()
}

// ReRender re-builds the explorer
func ReRender() {
	setupExplorerGUI()
	renderFileList()
}

func setupExplorerGUI() {
	fileListWidget.ItemFgColor = ui.ColorYellow
	fileListWidget.Height = filelistWidgetDims.Height

	statusBarWidget.Height = 3
	timeWidget.Height = 3

	ui.Clear()
	ui.Body.Rows = ui.Body.Rows[:0]
	ui.Body.AddRows(
		ui.NewRow(ui.NewCol(12, 0, fileListWidget)),
		ui.NewRow(ui.NewCol(10, 0, statusBarWidget), ui.NewCol(2, 0, timeWidget)))

	ui.Body.Align()
}

func renderStatusMessage(text string) {
	statusBarWidget.Text = text
	renderExplorer()
}

func renderPathBar(path string) {
	fileListWidget.BorderLabel = path + " "
	renderExplorer()
}

func renderFileList() {
	fileListWidget.Items = getPrettyList()
	renderExplorer()
}

func renderExplorer() {
	ui.Render(ui.Body)
}
