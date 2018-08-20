package logger

import (
	"log"
	"os"
)

// StartLogger initializes logging to file
func StartLogger() *os.File {
	cwd, err := os.Getwd()
	logpath := cwd + "/pixi.log"

	file, err := os.OpenFile(logpath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		panic("Error opening log file")
	}

	log.SetOutput(file)
	log.SetFlags(log.LstdFlags | log.Llongfile)

	return file
}
