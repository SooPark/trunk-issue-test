package main

import (
	"fmt"

	"github.com/ROKT/trunk-private-module-test/logger"
	"github.com/SooPark/trunk-issue-test/clients"
)

func main() {
	zapLogger := logger.GetLogger("test")
	log := zapLogger.Sugar()
	log.Debug("test")

	c := clients.NewClient("BaseURL", "APIKey")
	fmt.Print(c.Client.BaseURL)
}