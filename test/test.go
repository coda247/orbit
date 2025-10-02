package main

import "github.com/coda247/orbit/services"

func main() {
	logger := services.NewLoggerService("Finex")

	logger.Info("This is an info message")
}
