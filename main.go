package main

import (
	"github.com/gtaylor314/Banking-MS/app"
	"github.com/gtaylor314/Banking-MS/logger"
)

func main() {
	logger.Info("Starting the application")
	app.Start()
}
