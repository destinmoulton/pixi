package settingsform

import (
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

var uiScreen *tview.Grid
var uiForm *tview.Form

// UI creates the help window
func UI() *tview.Grid {
	uiScreen = tview.NewGrid().SetRows(0).SetColumns(0).SetBorders(true)
	uiForm = tview.NewForm().
		AddInputField("omxmplayer command", "", 40, nil, nil).
		AddButton("Save", nil)

	uiFrame := tview.NewFrame(uiForm).
		AddText("Settings", true, tview.AlignCenter, tcell.ColorDarkMagenta).
		AddText("ESC or s to leave Settings", false, tview.AlignCenter, tcell.ColorDarkMagenta)

	uiScreen.AddItem(uiFrame, 0, 0, 1, 1, 0, 0, false)

	return uiScreen
}
