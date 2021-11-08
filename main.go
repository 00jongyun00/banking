package main

import (
	"bangking/app"

	"github.com/nothingprogram/banking-lib/logger"
)

func main() {
	logger.Info("Starting the application...")
	app.Start()
}
