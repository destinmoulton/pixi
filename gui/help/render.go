package help

import (
	"bytes"
	"fmt"
	"text/tabwriter"

	ui "github.com/gizak/termui"
)

var helpLines = []string{
	".\tToggle hidden files and folders.",
	"q or Ctrl+c   \tQuit/Exit Pixi",
	"ESC\tClose Help - Return to Explorer",
}

var helpList = ui.NewPar("")

// Render the help window
func Render() {
	helpList.Text = tabLines(helpLines)
	helpList.BorderLabel = "Pixi Help"
	helpList.Height = ui.TermHeight()

	ui.Clear()
	ui.Body.Rows = ui.Body.Rows[:0]
	ui.Body.AddRows(
		ui.NewRow(ui.NewCol(12, 0, helpList)))

	ui.Body.Align()
	ui.Render(ui.Body)
}

func tabLines(lines []string) string {
	w := new(tabwriter.Writer)

	buf := new(bytes.Buffer)
	w.Init(buf, 0, 8, 0, ' ', 0)
	for _, line := range lines {
		fmt.Fprintln(w, line)
	}
	fmt.Fprintln(w)
	w.Flush()

	return buf.String()
}
