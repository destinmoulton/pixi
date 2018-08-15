package gui

import (
	"../dirlist"
	ui "github.com/gizak/termui"
)

// setupEvents creates the termui event handlers
func setupEvents() {
	ui.Handle("/sys/kbd/q", func(ui.Event) {
		ui.StopLoop()
	})

	ui.Handle("/sys/kbd/C-c", func(ui.Event) {
		ui.StopLoop()
	})

	ui.Handle("/sys/kbd/.", func(ui.Event) {
		dirlist.ToggleHidden()
		updateFileList()
	})

	ui.Handle("/sys/kbd/<up>", func(ui.Event) {
		dirlist.SelectPrevElement()
		updateFileList()
	})

	ui.Handle("/sys/kbd/<down>", func(ui.Event) {
		dirlist.SelectNextElement()
		updateFileList()
	})

	ui.Handle("/sys/kbd/<left>", func(ui.Event) {
		dirlist.NavUpDirectory()
		updateFileList()
	})

	ui.Handle("/sys/kbd/<right>", func(ui.Event) {
		dirlist.PerformFileAction()
		updateFileList()
	})

	ui.Handle("/sys/kbd/<enter>", func(ui.Event) {
		dirlist.PerformFileAction()
		updateFileList()
	})
}
