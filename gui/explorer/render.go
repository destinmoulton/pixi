package explorer

import (
	ui "github.com/gizak/termui"

	"../../types"
)

// fileListWidget termui widget List
var fileListWidget = ui.NewList()

// statusBarWidget termui widget Status Bar
var statusBarWidget = ui.NewPar("")

var filelistWidgetDims types.WidgetDimensions

// InitExplorer initializes the File Explorer
func InitExplorer() {
	filelistWidgetDims.Width = ui.TermWidth()
	filelistWidgetDims.Height = ui.TermHeight() - 3

	setupExplorerGUI()

	initFileList()

	renderFileList()
}

func setupExplorerGUI() {

	fileListWidget.ItemFgColor = ui.ColorYellow
	fileListWidget.Height = filelistWidgetDims.Height

	statusBarWidget.Height = 3
	statusBarWidget.Text = ""

	ui.Body.AddRows(
		ui.NewRow(ui.NewCol(12, 0, fileListWidget)),
		ui.NewRow(ui.NewCol(12, 0, statusBarWidget)))

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
