package main

import (
	"starscanner/command"
	"starscanner/log"
)

var (
	version = "dev"
	logger  = log.NewLogger(true)
)

func main() {
	logger.Debugf("Starting StarScanner (Version: %s)", version)
	command.Main(version)
}
