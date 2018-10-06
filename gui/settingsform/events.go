package settingsform

import "github.com/gdamore/tcell"

var parentScreenSwitchPage func(string)

// HandleEvents dispatches key events for the help package
func HandleEvents(eventKey *tcell.EventKey, switchToPage func(string)) *tcell.EventKey {
	parentScreenSwitchPage = switchToPage

	if eventKey.Key() == tcell.KeyEsc {
		switchToPage("explorer")
		return eventKey
	}

	return eventKey
}
