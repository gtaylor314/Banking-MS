package main

import (
	"github.com/gtaylor314/Banking-Lib/logger"
	"github.com/gtaylor314/Banking-MS/app"
)

func main() {
	logger.Info("Starting the application")
	app.Start()
}
