package history

import "github.com/gdamore/tcell"

func HandleEvents(eventKey *tcell.EventKey, switchToPage func(string)) *tcell.EventKey {
	if eventKey.Key() == tcell.KeyEsc || eventKey.Rune() == 'h' {
		switchToPage("explorer")
		return eventKey
	}
	return eventKey
}
