package settingsform

import (
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"

	"../../settings"
)

var command string
var redrawParent func()
var uiScreen *tview.Grid
var uiForm *tview.Form

// UI creates the help window
func UI(redraw func()) *tview.Form {
	redrawParent = redraw
	command = settings.Get(settings.SetConfig, settings.KeyOmxplayerCommand).(string)

	uiForm = tview.NewForm().
		AddInputField("omxmplayer command", command, 40, nil, handleCommandChange).
		AddButton("Save", handlePressSave).
		AddButton("Cancel", handlePressCancel)

	uiForm.SetFieldBackgroundColor(tcell.ColorGold).SetFieldTextColor(tcell.ColorBlack)
	uiForm.SetBorder(true).SetTitle("Settings")

	return uiForm
}

func handleCommandChange(cmd string) {
	command = cmd
}

func handlePressSave() {
	settings.Set(settings.SetConfig, settings.KeyOmxplayerCommand, command)
	parentScreenSwitchPage("explorer")
}

func handlePressCancel() {
	parentScreenSwitchPage("explorer")
}
