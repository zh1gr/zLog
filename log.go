package log

import (
	"fmt"
	"io"
	"os"
	"time"
)

//# Logs
//# 0 - off
//# 1 - errors
//# 2 - warnings
//# 3 - info
//# 4 - debug (all log levels)

var (
	mainLogLevel = 3
	timeFormat   = "2006-01-02 15:04:05.000000"
	lERROR       = logOut{os.Stderr, "ERROR", 1}
	lWARNING     = logOut{os.Stdout, "WARNING", 2}
	lINFO        = logOut{os.Stdout, "INFO", 3}
	lDEBUG       = logOut{os.Stdout, "DEBUG", 4}
)

type logOut struct {
	Output  io.Writer
	Caption string
	Level   int
}

func SetLogLevel(lvl int) {
	if lvl >= 0 && lvl < 5 {
		mainLogLevel = lvl
	}
}

func SetTimeFormat(f string) {
	timeFormat = f
}

func Error(err interface{}, i ...interface{}) {
	switch v := err.(type) {
	case error:
		fPrintLog(lERROR, v.Error())
	case string:
		fPrintLog(lERROR, fmt.Sprintf(v, i...))
	default:
		fPrintLog(lERROR, v)
	}
}

func Warning(msg string, i ...interface{}) {
	fPrintLog(lWARNING, fmt.Sprintf(msg, i...))
}

func Info(msg string, i ...interface{}) {
	fPrintLog(lINFO, fmt.Sprintf(msg, i...))
}

func Debug(msg string, i ...interface{}) {
	fPrintLog(lDEBUG, fmt.Sprintf(msg, i...))
}

func getTimeNow() string {
	return time.Now().Format(timeFormat)
}

func fPrintLog(l logOut, s ...interface{}) {
	if mainLogLevel <= l.Level {
		_, err := fmt.Fprintf(l.Output, "[%s][%s]: %s\n", getTimeNow(), l.Output, s)
		if err != nil {
			fmt.Fprintf(os.Stderr, "[%s][ERROR][LOG]: %s",getTimeNow(), err)
		}
	}
}
