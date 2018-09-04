package gui

import (
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"

	"./explorer"
	"./help"
	"./history"
)

var app *tview.Application
var pages *tview.Pages
var activePage string

// Init starts the gui
func Init() {
	app = tview.NewApplication()
	pages = tview.NewPages()
	redraw := func() {
		app.Draw()
	}
	pages.AddPage("explorer", explorer.UI(redraw), true, true)
	pages.AddPage("history", history.UI(redraw), true, false)
	pages.AddPage("help", help.UI(), true, false)
	activePage = "explorer"

	explorer.StartExplorer()
	history.StartHistory()

	app.SetInputCapture(exitHandler)
	if err := app.SetRoot(pages, true).SetFocus(pages).Run(); err != nil {
		panic(err)
	}
}

func switchToPage(page string) {
	activePage = page
	pages.SwitchToPage(page)
}

func exitHandler(eventKey *tcell.EventKey) *tcell.EventKey {
	if eventKey.Rune() == 'q' {
		app.Stop()
		return nil
	}

	if activePage == "help" {
		if eventKey.Key() == tcell.KeyEsc || eventKey.Key() == tcell.KeyF1 {
			switchToPage("explorer")
			return eventKey

		}
	}

	if activePage == "history" {
		return history.HandleEvents(eventKey, switchToPage)
	}

	if activePage == "explorer" {

		if eventKey.Key() == tcell.KeyF1 {
			switchToPage("help")
			return eventKey
		}

		if eventKey.Rune() == 'h' {
			switchToPage("history")
			return eventKey
		}

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
