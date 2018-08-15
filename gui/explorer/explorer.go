package gui

import (
	ui "github.com/gizak/termui"

	"../dirlist"
	"../types"
)

// FileExplorerWidget termui widget List
var FileExplorerWidget = ui.NewList()

// StatusBarWidget termui widget Status Bar
var StatusBarWidget = ui.NewPar("")

var fileExplorerWidgetDimensions types.WidgetDimensions

// InitFileExplorer initializes the File Explorer
func InitFileExplorer() {
	setupEvents()

	fileExplorerWidgetDimensions.Width = ui.TermWidth()
	fileExplorerWidgetDimensions.Height = ui.TermHeight() - 3
	dirlist.Init(fileExplorerWidgetDimensions, updatePathBar, updateStatusMessage)
	renderList(dirlist.GetPrettyList())
}

func updateStatusMessage(text string) {
	StatusBarWidget.Text = text
	ui.Render(ui.Body)
}

func updatePathBar(path string) {
	FileExplorerWidget.BorderLabel = path + " "
	ui.Render(ui.Body)
}

func updateFileList() {
	FileExplorerWidget.Items = dirlist.GetPrettyList()
	ui.Render(ui.Body)
}

func renderList(dirList []string) {

	FileExplorerWidget.Items = dirList
	FileExplorerWidget.ItemFgColor = ui.ColorYellow
	FileExplorerWidget.Height = fileExplorerWidgetDimensions.Height

	StatusBarWidget.Height = 3
	StatusBarWidget.Text = ""

	ui.Body.AddRows(
		ui.NewRow(ui.NewCol(12, 0, FileExplorerWidget)),
		ui.NewRow(ui.NewCol(12, 0, StatusBarWidget)))

	ui.Body.Align()
	ui.Render(ui.Body)
}
