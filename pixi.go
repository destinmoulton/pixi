package main

import (
	"flag"

	"./gui"
	"./logger"
	_ "./settings"
)

var shouldLog bool

func init() {
	flag.BoolVar(&shouldLog, "log", false, "Enable logging to file (pixi.log) in current directory")
	flag.Parse()
}

func main() {

	if shouldLog {
		fileHandler := logger.StartLogger()
		defer fileHandler.Close()
	}

	gui.Init()
}
