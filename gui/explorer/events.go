package explorer

import "github.com/gdamore/tcell"

func HandleEvents(eventKey *tcell.EventKey, switchToPage func(string)) *tcell.EventKey {
	if eventKey.Key() == tcell.KeyF1 {
		switchToPage("help")
		return eventKey
	}

	if eventKey.Rune() == 'h' {
		switchToPage("history")
		return eventKey
	}

	if eventKey.Key() == tcell.KeyLeft {
		NavUpDirectory()
		return nil
	}

	if eventKey.Key() == tcell.KeyRight {
		NavIntoDirectory()
		return nil
	}

	if eventKey.Key() == tcell.KeyEnter {
		PerformFileAction()
		return nil
	}
	return eventKey
}
