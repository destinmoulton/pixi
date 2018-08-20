package explorer

import (
	"log"
	"time"

	ui "github.com/gizak/termui"

	"../../types"
)

// fileListWidget termui widget List
var fileListWidget = ui.NewList()

// statusBarWidget termui widget Status Bar
var statusBarWidget = ui.NewPar("[F1](fg-green) Help | [h](fg-green) Toggle Hidden Files")
var clockWidget = ui.NewPar("Time")
var filelistWidgetDims types.WidgetDimensions
var clockTicker *time.Ticker

// StartExplorer initializes the File Explorer
func StartExplorer() {
	log.Println("StartExplorer() running")
	filelistWidgetDims.Width = ui.TermWidth()
	filelistWidgetDims.Height = ui.TermHeight() - 3

	setupExplorerGUI()

	initFileList()

	renderFileList()
	startClock()
	renderClock(time.Now())
}

// ReRender re-builds the explorer
func ReRender() {
	setupExplorerGUI()
	renderFileList()
}

func setupExplorerGUI() {
	fileListWidget.ItemFgColor = ui.ColorYellow
	fileListWidget.Height = filelistWidgetDims.Height

	statusBarWidget.Height = 3
	clockWidget.Height = 3

	ui.Clear()
	ui.Body.Rows = ui.Body.Rows[:0]
	ui.Body.AddRows(
		ui.NewRow(ui.NewCol(12, 0, fileListWidget)),
		ui.NewRow(ui.NewCol(10, 0, statusBarWidget), ui.NewCol(2, 0, clockWidget)))

	ui.Body.Align()
}

func renderStatusMessage(text string) {
	statusBarWidget.Text = text
	renderExplorer()
}

func renderPathBar(path string) {
	fileListWidget.BorderLabel = path + " "
	renderExplorer()
}

func renderFileList() {
	fileListWidget.Items = getPrettyList()
	renderExplorer()
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
	clockWidget.Text = t.Format("   3:04 pm")
	renderExplorer()
}

func renderExplorer() {
	ui.Render(ui.Body)
}
