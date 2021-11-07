package main

import (
	"bangking/app"
	"bangking/logger"
)

func main() {
	logger.Info("Starting the application...")
	app.Start()
}
