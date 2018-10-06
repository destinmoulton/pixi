package gui

import (
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"

	"./explorer"
	"./help"
	"./history"
	"./settingsform"
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
	pages.AddPage("settingsform", settingsform.UI(redraw), true, false)
	activePage = "explorer"

	explorer.StartExplorer()
	history.StartHistory()

	app.SetInputCapture(eventHandler)
	if err := app.SetRoot(pages, true).SetFocus(pages).Run(); err != nil {
		panic(err)
	}
}

func switchToPage(page string) {
	activePage = page
	pages.SwitchToPage(page)
}

func eventHandler(eventKey *tcell.EventKey) *tcell.EventKey {

	if eventKey.Rune() == 'q' {
		app.Stop()
		return nil
	}

	if activePage == "help" {
		return help.HandleEvents(eventKey, switchToPage)
	}

	if activePage == "history" {
		return history.HandleEvents(eventKey, switchToPage)
	}

	if activePage == "explorer" {
		return explorer.HandleEvents(eventKey, switchToPage)
	}

	if activePage == "settingsform" {
		return settingsform.HandleEvents(eventKey, switchToPage)
	}

	return eventKey
}
