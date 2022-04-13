package main

import (
	"github.com/kenethrrizzo/banking/app"
	"github.com/kenethrrizzo/banking/logger"
)

func main() {
	logger.Info("Starting the application...")
	app.Start()
}
