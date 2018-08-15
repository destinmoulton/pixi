package gui

import (
	ui "github.com/gizak/termui"

	"./explorer"
)

// Init starts the gui
func Init() {
	err := ui.Init()
	if err != nil {
		panic(err)
	}
	defer ui.Close()

	setupEvents()

	explorer.InitExplorer()

	ui.Loop()
}
