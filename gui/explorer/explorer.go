package explorer

import (
	"time"

	"github.com/rivo/tview"
)

var redrawParent func()
var uiScreen *tview.Grid

// tableWidget termui widget List
var tableWidget *tview.Table

// pathWidget termui widget Status Bar
var pathWidget *tview.TextView

var clockWidget *tview.TextView
var clockTicker *time.Ticker

// StartExplorer initializes the File Explorer
func StartExplorer() {
	initFileList()

	renderFileList()
	startClock()
	renderClock(time.Now())
}

// ReRender re-builds the explorer
func ReRender() {
	renderFileList()
}

// UI builds the gui for the explorer list of files
func UI(redraw func()) *tview.Grid {
	redrawParent = redraw
	uiScreen = tview.NewGrid().SetRows(1, 0).SetColumns(0, 10).SetBorders(true)
	tableWidget = tview.NewTable().SetBorders(false)
	pathWidget = tview.NewTextView().SetTextAlign(tview.AlignLeft).SetText("")
	clockWidget = tview.NewTextView().SetTextAlign(tview.AlignCenter)

	uiScreen.AddItem(pathWidget, 0, 0, 1, 1, 0, 0, false)
	uiScreen.AddItem(clockWidget, 0, 1, 1, 1, 0, 0, false)
	uiScreen.AddItem(tableWidget, 1, 0, 1, 2, 0, 0, true)

	return uiScreen
}

func uiGetSelectedFileIndex() int {
	sel, _ := tableWidget.GetSelection()
	return sel
}

func uiSetSelectedFileIndex(i int) {
	tableWidget.Select(i, 0)
}

func uiScrollToTop() {
	tableWidget.ScrollToBeginning()
}

func setPathWidgetText(text string) {
	pathWidget.SetText(text)
}

func renderFileList() {
	tableWidget.Clear()
	redrawParent()
	items := getPrettyList()

	for i, item := range items {
		cell := tview.NewTableCell(item.filename).
			SetTextColor(item.fgColor).
			SetBackgroundColor(item.bgColor).
			SetExpansion(2)

		tableWidget.SetCell(i, 0, cell)
	}

	tableWidget.Select(0, 0).SetSelectable(true, false)
	redrawParent()
}

func startClock() {
	clockTicker = time.NewTicker(time.Second * 10)
	go func() {
		for t := range clockTicker.C {
			renderClock(t)
		}
	}()

}

func renderClock(t time.Time) {
	clockWidget.SetText(t.Format("3:04 pm"))
	redrawParent()
}
