package help

import (
	"bytes"
	"fmt"
	"text/tabwriter"

	"github.com/rivo/tview"
)

var helpLines = []string{
	"  \t",
	"  Up/Down Arrows     \tSelect directories/files.",
	"  Left Arrow\tNavigate up to parent directory.",
	"  Right Arrow\tNavigate into selected directory.",
	"  Enter/Return\tPlay selected file.",
	"  h\tToggle hidden files and folders.",
	"  q or Ctrl+c   \tQuit/Exit Pixi",
	"  \t",
	"  ESC/F1\tClose Help - Return to Explorer",
}

var uiScreen *tview.Grid
var helpWidget *tview.TextView

// Render the help window
func UI() *tview.Grid {
	uiScreen = tview.NewGrid().SetRows(0).SetColumns(0).SetBorders(true)
	helpWidget = tview.NewTextView().SetText(tabLines(helpLines))

	uiScreen.AddItem(helpWidget, 0, 0, 1, 1, 0, 0, false)

	return uiScreen
}

func tabLines(lines []string) string {
	w := new(tabwriter.Writer)

	buf := new(bytes.Buffer)
	w.Init(buf, 0, 10, 0, ' ', 0)
	for _, line := range lines {
		fmt.Fprintln(w, line)
	}
	fmt.Fprintln(w)
	w.Flush()

	return buf.String()
}
