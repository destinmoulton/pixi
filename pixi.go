package main

import (
	"./dirlist"
	ui "github.com/gizak/termui"
)

func main() {
	err := ui.Init()
	if err != nil {
		panic(err)
	}
	defer ui.Close()

	dirlist.Init()

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
		renderList(dirlist.GetPrettyList())
	})

	ui.Handle("/sys/kbd/<down>", func(ui.Event) {
		dirlist.SelectNextElement()
		renderList(dirlist.GetPrettyList())
	})
}

func renderList(dirList []string) {
	ls := ui.NewList()
	ls.Items = dirList
	ls.ItemFgColor = ui.ColorYellow
	ls.BorderLabel = "Files"
	ls.Height = 20
	ls.Width = 70
	ls.Y = 0

	ui.Render(ls)
}
