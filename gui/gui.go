package gui

import (
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"

	"./explorer"
)

var app *tview.Application
var activePage = "explorer"

// Init starts the gui
func Init() {
	app = tview.NewApplication()
	pages := tview.NewPages()
	redraw := func() {
		app.Draw()
	}
	pages.AddPage("explorer", explorer.UI(redraw), true, true)
	explorer.StartExplorer()

	app.SetInputCapture(exitHandler)
	if err := app.SetRoot(pages, true).SetFocus(pages).Run(); err != nil {
		panic(err)
	}
}

func exitHandler(eventKey *tcell.EventKey) *tcell.EventKey {
	if eventKey.Rune() == 'q' {
		app.Stop()
		return nil
	}

	if activePage == "explorer" {
		if eventKey.Key() == tcell.KeyLeft {
			explorer.NavUpDirectory()
			return nil
		}

		if eventKey.Key() == tcell.KeyRight {
			explorer.NavIntoDirectory()
			return nil
		}

		if eventKey.Key() == tcell.KeyEnter {
			explorer.PerformFileAction()
			return nil
		}

	}

	return eventKey
}
