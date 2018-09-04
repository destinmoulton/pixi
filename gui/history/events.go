package history

import (
	"github.com/gdamore/tcell"

	"../../player"
)

// HandleEvents dispatches events for the history widget
func HandleEvents(eventKey *tcell.EventKey, switchToPage func(string)) *tcell.EventKey {
	if eventKey.Key() == tcell.KeyEsc || eventKey.Rune() == 'h' || eventKey.Key() == tcell.KeyLeft {
		switchToPage("explorer")
		return eventKey
	}

	if eventKey.Key() == tcell.KeyEnter {
		sel := getSelectedFile()
		if player.IsVideoFile(sel["filename"]) {
			player.PlayVideo(sel["path"])
		}
		return nil
	}

	return eventKey
}
