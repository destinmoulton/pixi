package explorer

import (
	"time"

	"github.com/rivo/tview"

	"../../types"
)

var redrawParent func()
var uiScreen *tview.Grid

// listWidget termui widget List
var listWidget *tview.Table

//var listBox = tview.NewFrame(listWidget)

// infoWidget termui widget Status Bar
var infoWidget *tview.TextView

var clockWidget *tview.TextView

var listWidgetDims types.WidgetDimensions
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
	uiScreen = tview.NewGrid().SetRows(0, 1).SetColumns(0, 10).SetBorders(true)
	listWidget = tview.NewTable().SetBorders(false)
	infoWidget = tview.NewTextView().SetTextAlign(tview.AlignLeft).SetText("Test")
	clockWidget = tview.NewTextView().SetTextAlign(tview.AlignCenter)

	uiScreen.AddItem(listWidget, 0, 0, 1, 2, 0, 0, true)
	uiScreen.AddItem(infoWidget, 1, 0, 1, 1, 0, 0, false)
	uiScreen.AddItem(clockWidget, 1, 1, 1, 1, 0, 0, false)
	return uiScreen
}

func getSelectedFileIndex() int {
	sel, _ := listWidget.GetSelection()
	return sel
}

func renderStatusMessage(text string) {
	infoWidget.SetText(text)
}

func renderPathBar(path string) {
	listWidget.SetTitle(path)
}

func renderFileList() {
	listWidget.Clear()
	redrawParent()
	items := getPrettyList()

	for i, item := range items {
		listWidget.SetCell(i, 0, tview.NewTableCell(item.filename).SetTextColor(item.fgColor).SetBackgroundColor(item.bgColor))
	}

	listWidget.Select(0, 0).SetSelectable(true, false)
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
