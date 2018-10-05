package help

import "github.com/gdamore/tcell"

// HandleEvents dispatches key events for the help package
func HandleEvents(eventKey *tcell.EventKey, switchToPage func(string)) *tcell.EventKey {
	if eventKey.Key() == tcell.KeyEsc || eventKey.Rune() == 's' {
		switchToPage("explorer")
		return eventKey
	}

	return eventKey
}
