package log

import (
	"fmt"
	"log"
	"os"

	"github.com/fatih/color"
)

type LoggerImpl struct {
	*log.Logger
}

type LogLevel int

const (
	PREFIX    = "-->"
	LOG_PRINT = iota
	LOG_FATAL
	LOG_ERROR
	LOG_WARN
	LOG_INFO
	LOG_DEBUG
)

var levelNames = map[LogLevel]string{
	LOG_PRINT: color.WhiteString(PREFIX),
	LOG_FATAL: color.HiRedString(PREFIX),
	LOG_ERROR: color.RedString(PREFIX),
	LOG_WARN:  color.YellowString(PREFIX),
	LOG_INFO:  color.GreenString(PREFIX),
	LOG_DEBUG: color.BlueString(PREFIX),
}

func NewLogger() *LoggerImpl {
	logger := &LoggerImpl{Logger: log.New(os.Stdout, "", log.Lshortfile)}

	return logger
}

func (l *LoggerImpl) Printf(msg string, v ...interface{}) {
	l.writeLog(levelNames[LOG_PRINT], msg, v...)
}

func (l *LoggerImpl) Infof(msg string, v ...interface{}) {
	l.writeLog(levelNames[LOG_INFO], msg, v...)
}

func (l *LoggerImpl) Errorf(v ...interface{}) {
	l.Logger.Printf("%v", v...)
}

func (l *LoggerImpl) Error(v ...interface{}) {
	l.Logger.Print(v...)
}

func (l *LoggerImpl) writeLog(level string, msg string, v ...interface{}) {
	outStr := fmt.Sprintf(msg, v...)
	out := fmt.Sprintf("%s %v", level, outStr)

	fmt.Println(out)
}
