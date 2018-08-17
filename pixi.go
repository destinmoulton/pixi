package main

import (
	"./config"
	"./gui"
)

func main() {
	config.Init()
	gui.Init()
}
