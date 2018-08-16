package gui

import (
	ui "github.com/gizak/termui"

	"./explorer"
	"./help"
)

var isHelpActive = false

// setupEvents creates the termui event handlers
func setupEvents() {
	ui.Handle("/sys/kbd/q", func(ui.Event) {
		ui.StopLoop()
	})

	ui.Handle("/sys/kbd/C-c", func(ui.Event) {
		ui.StopLoop()
	})

	ui.Handle("/sys/kbd/h", func(ui.Event) {
		if isHelpActive {
			explorer.ReRender()
		} else {
			help.Render()
		}

		isHelpActive = !isHelpActive
	})

	ui.Handle("/sys/kbd/<esc>", func(ui.Event) {
		if isHelpActive {
			explorer.ReRender()
			isHelpActive = false
		}
	})

	ui.Handle("/sys/kbd/.", func(ui.Event) {
		explorer.ToggleHidden()
	})

	ui.Handle("/sys/kbd/<up>", func(ui.Event) {
		explorer.SelectPrevFile()
	})

	ui.Handle("/sys/kbd/<down>", func(ui.Event) {
		explorer.SelectNextFile()
	})

	ui.Handle("/sys/kbd/<left>", func(ui.Event) {
		explorer.NavUpDirectory()
	})

	ui.Handle("/sys/kbd/<right>", func(ui.Event) {
		explorer.PerformFileAction()
	})

	ui.Handle("/sys/kbd/<enter>", func(ui.Event) {
		explorer.PerformFileAction()
	})
}
