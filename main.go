package main

import (
	"starscanner/command"
	"starscanner/log"
)

var (
	version = "dev"
	logger  = log.NewLogger()
)

func main() {
	logger.Infof("Starting StarScanner (Version: %s)", version)
	command.Main(version)
}
