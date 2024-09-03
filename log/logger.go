package log

import (
	"os"
	"time"

	charmLog "github.com/charmbracelet/log"
)

const (
	level  = "debug"
	output = "console"
)

// NewLogger creates a new logger
func NewLogger() *charmLog.Logger {
	var logLevel charmLog.Level

	switch level {
	case "debug":
		logLevel = charmLog.DebugLevel
	case "info":
		logLevel = charmLog.InfoLevel
	case "warn":
		logLevel = charmLog.WarnLevel
	case "error":
		logLevel = charmLog.ErrorLevel
	case "fatal":
		logLevel = charmLog.FatalLevel
	default:
		logLevel = charmLog.InfoLevel
	}

	if output == "file" {
		file, err := os.OpenFile("LogManager.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
		if err != nil {
			panic(err)
		}

		return charmLog.NewWithOptions(file, charmLog.Options{
			Level:           logLevel,
			ReportCaller:    false,
			ReportTimestamp: true,
			TimeFormat:      time.DateTime,
		})
	}

	return charmLog.NewWithOptions(os.Stdout, charmLog.Options{
		Level:           logLevel,
		ReportCaller:    false,
		ReportTimestamp: true,
		TimeFormat:      "15:04:05.0000",
	})
}
