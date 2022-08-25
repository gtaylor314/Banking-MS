package main

import (
	"Banking/app"
	"Banking/logger"
)

func main() {
	logger.Info("Starting the application")
	app.Start()
}
