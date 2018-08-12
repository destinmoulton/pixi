package main

import (
	ui "github.com/gizak/termui"

	"./dirlist"
	"./types"
)

var widgetFileList = ui.NewList()
var widgetStatusBar = ui.NewPar("")

var widgetFileListDimensions types.WidgetDimensions

func main() {
	err := ui.Init()
	if err != nil {
		panic(err)
	}
	defer ui.Close()

	widgetFileListDimensions.Width = ui.TermWidth()
	widgetFileListDimensions.Height = ui.TermHeight() - 3
	dirlist.Init(widgetFileListDimensions, updateStatusMessage)

	renderList(dirlist.GetPrettyList())

	setupEvents()

	ui.Loop()

}

func setupEvents() {
	ui.Handle("/sys/kbd/q", func(ui.Event) {
		ui.StopLoop()
	})

	ui.Handle("/sys/kbd/<escape>", func(ui.Event) {
		ui.StopLoop()
	})

	ui.Handle("/sys/kbd/C-c", func(ui.Event) {
		ui.StopLoop()
	})

	ui.Handle("/sys/kbd/<up>", func(ui.Event) {
		dirlist.SelectPrevElement()
		updateFileList(dirlist.GetPrettyList())
	})

	ui.Handle("/sys/kbd/<down>", func(ui.Event) {
		dirlist.SelectNextElement()
		updateFileList(dirlist.GetPrettyList())
	})

	ui.Handle("/sys/kbd/<left>", func(ui.Event) {
		dirlist.NavUpDirectory()
		updateFileList(dirlist.GetPrettyList())
	})

	ui.Handle("/sys/kbd/<right>", func(ui.Event) {
		dirlist.PerformFileAction()
		updateFileList(dirlist.GetPrettyList())
	})
}

func updateStatusMessage(text string) {
	widgetStatusBar.Text = text
	ui.Render(ui.Body)
}

func updateFileList(dirList []string) {

	widgetFileList.Items = dirList
	ui.Render(ui.Body)
}

func renderList(dirList []string) {

	widgetFileList.Items = dirList
	widgetFileList.ItemFgColor = ui.ColorYellow
	widgetFileList.BorderLabel = "Files"
	widgetFileList.Height = widgetFileListDimensions.Height

	widgetStatusBar.Height = 3
	widgetStatusBar.Text = ""

	ui.Body.AddRows(
		ui.NewRow(ui.NewCol(12, 0, widgetFileList)),
		ui.NewRow(ui.NewCol(12, 0, widgetStatusBar)))

	ui.Body.Align()
	ui.Render(ui.Body)
}
