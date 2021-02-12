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
	level      = 3
	timeFormat = "2006-01-02 15:04:05.000000"
	ERROR      = logOut{os.Stderr, "ERROR"}
	WARNING    = logOut{os.Stdout, "WARNING"}
	INFO       = logOut{os.Stdout, "INFO"}
	DEBUG      = logOut{os.Stdout, "DEBUG"}
)

type logOut struct {
	Output   io.Writer
	LogLevel string
}

func SetLogLevel(lvl int) {
	if lvl >= 0 && lvl < 5 {
		level = lvl
	}
}

func SetTimeFormat(f string) {
	timeFormat = f
}

func Error(err interface{}, i ...interface{}) {
	if level > 0 {
		switch v := err.(type) {
		case error:
			fPrintLog(ERROR, v.Error())
		case string:
			fPrintLog(ERROR, fmt.Sprintf(v, i...))
		default:
			fPrintLog(ERROR, v)
		}
	}
}

func Warning(msg string, i ...interface{}) {
	if level > 1 {
		fPrintLog(WARNING, fmt.Sprintf(msg, i...))
	}
}

func Info(msg string, i ...interface{}) {
	if level > 2 {
		fPrintLog(INFO, fmt.Sprintf(msg, i...))
	}
}

func Debug(msg string, i ...interface{}) {
	if level > 3 {
		fPrintLog(DEBUG, fmt.Sprintf(msg, i...))
	}
}

func getTimeNow() string {
	return time.Now().Format(timeFormat)
}

func fPrintLog(l logOut, s ...interface{}) {
	if _, err := fmt.Fprintf(l.Output, "[%s][%s]: %s\n", getTimeNow(), l.Output, s); err != nil {
		fPrintLog(ERROR, err)
	}
}
