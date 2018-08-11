package main

import (
	"./dirlist"
	ui "github.com/gizak/termui"
)

var uiFileList = ui.NewList()
var uiStatusBar = ui.NewPar("")

func main() {
	err := ui.Init()
	if err != nil {
		panic(err)
	}
	defer ui.Close()

	dirlist.Init(ui.TermWidth(), updateStatusMessage)

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
	uiStatusBar.Text = text
	ui.Render(ui.Body)
}

func updateFileList(dirList []string) {

	uiFileList.Items = dirList
	ui.Render(ui.Body)
}

func renderList(dirList []string) {

	uiFileList.Items = dirList
	uiFileList.ItemFgColor = ui.ColorYellow
	uiFileList.BorderLabel = "Files"
	uiFileList.Height = ui.TermHeight() - 3

	uiStatusBar.Height = 3
	uiStatusBar.Text = ""

	ui.Body.AddRows(
		ui.NewRow(ui.NewCol(12, 0, uiFileList)),
		ui.NewRow(ui.NewCol(12, 0, uiStatusBar)))

	ui.Body.Align()
	ui.Render(ui.Body)
}
