package main

import (
	"./dirlist"
	"github.com/gizak/termui"
)

func main() {
	err := termui.Init()
	if err != nil {
		panic(err)
	}
	defer termui.Close()

	dirlist.Init()

	renderList(dirlist.GetPrettyList())

	setupEvents()

	termui.Loop()

}

func setupEvents() {
	termui.Handle("/sys/kbd/q", func(termui.Event) {
		termui.StopLoop()
	})

	termui.Handle("/sys/kbd/<escape>", func(termui.Event) {
		termui.StopLoop()
	})

	termui.Handle("/sys/kbd/C-c", func(termui.Event) {
		termui.StopLoop()
	})

	termui.Handle("/sys/kbd/<up>", func(termui.Event) {
		dirlist.SelectPrevElement()
		renderList(dirlist.GetPrettyList())
	})

	termui.Handle("/sys/kbd/<down>", func(termui.Event) {
		dirlist.SelectNextElement()
		renderList(dirlist.GetPrettyList())
	})
}

func renderList(dirList []string) {
	ls := termui.NewList()
	ls.Items = dirList
	ls.ItemFgColor = termui.ColorYellow
	ls.BorderLabel = "Files"
	ls.Height = 20
	ls.Width = 70
	ls.Y = 0

	termui.Render(ls)
}
