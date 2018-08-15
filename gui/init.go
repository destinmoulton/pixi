package gui

import (
	ui "github.com/gizak/termui"
)

// Init starts the gui
func Init() {
	err := ui.Init()
	if err != nil {
		panic(err)
	}
	defer ui.Close()

	setupEvents()

	InitFileExplorer()

	ui.Loop()
}
