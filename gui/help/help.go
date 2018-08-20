package help

import (
	"bytes"
	"fmt"
	"text/tabwriter"

	ui "github.com/gizak/termui"
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

var helpList = ui.NewPar("")

// Render the help window
func Render() {
	helpList.Text = tabLines(helpLines)
	helpList.BorderLabel = "Pixi Help "
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
	w.Init(buf, 0, 10, 0, ' ', 0)
	for _, line := range lines {
		fmt.Fprintln(w, line)
	}
	fmt.Fprintln(w)
	w.Flush()

	return buf.String()
}
