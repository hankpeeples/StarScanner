package log

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/fatih/color"
)

type LoggerImpl struct {
	*log.Logger
	debug bool
}

type LogLevel int

const (
	DATE_FORMAT = "2020-11-03"
	TIME_FORMAT = "15:04:00 AM"
	PREFIX      = "●"
	LOG_PRINT   = iota
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

func NewLogger(isDebug bool) *LoggerImpl {
	logger := &LoggerImpl{Logger: log.New(os.Stdout, "", log.Lshortfile), debug: isDebug}

	return logger
}

// ------------------- Print -----------------------

func (l *LoggerImpl) Printf(msg string, v ...interface{}) {
	l.writeLog(levelNames[LOG_PRINT], msg, v...)
}

func (l *LoggerImpl) Println(msg string) {
	l.writeLog(levelNames[LOG_PRINT], msg)
}

// ------------------- Print -----------------------
// ------------------- Info ------------------------

func (l *LoggerImpl) Infof(msg string, v ...interface{}) {
	l.writeLog(levelNames[LOG_INFO], msg, v...)
}

func (l *LoggerImpl) Infoln(msg string) {
	l.writeLog(levelNames[LOG_INFO], msg)
}

// ------------------- Info ------------------------
// ------------------- Warn ------------------------

func (l *LoggerImpl) Warnf(msg string, v ...interface{}) {
	l.writeLog(levelNames[LOG_WARN], msg, v...)
}

func (l *LoggerImpl) Warnln(msg string) {
	l.writeLog(levelNames[LOG_INFO], msg)
}

// ------------------- Warn -------------------------
// ------------------- Error ------------------------

func (l *LoggerImpl) Errorf(msg string, v ...interface{}) {
	l.writeLog(levelNames[LOG_ERROR], msg, v...)
}

func (l *LoggerImpl) Errorln(msg string) {
	l.writeLog(levelNames[LOG_INFO], msg)
}

// ------------------- Error -------------------------
// ------------------- Debug ------------------------

func (l *LoggerImpl) Debugf(msg string, v ...interface{}) {
	l.writeLog(levelNames[LOG_DEBUG], msg, v...)
}

func (l *LoggerImpl) Debugln(msg string) {
	l.writeLog(levelNames[LOG_INFO], msg)
}

// ------------------- Debug -------------------------

func (l *LoggerImpl) writeLog(level string, msg string, v ...interface{}) {
	outStr := fmt.Sprintf(msg, v...)
	var out string
	var curTime time.Time

	if l.debug {
		curTime = time.Now()
		out = fmt.Sprintf("%s %s %s %v", curTime.Format(DATE_FORMAT), curTime.Format(TIME_FORMAT), level, outStr)
	} else {
		out = fmt.Sprintf("%s %v", level, outStr)
	}

	fmt.Println(out)
}
