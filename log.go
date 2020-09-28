package log

import (
	"fmt"
	"os"
	"regexp"
	"time"
)

//# Logs
//# 0 - off
//# 1 - errors
//# 2 - info
//# 3 - warnings
//# 4 - debug (all log levels)

var (
	level      = 4
	timeFormat = "2006-01-02 15:04:05.000000Z07:00"
)

func SetLogLevel(lvl int) {
	if lvl >= 0 && lvl <= 4 {
		level = lvl
	}
}

func SetTimeFormat(f string) {
	timeFormat = f
}

func Debug(msg string, i ...interface{}) {
	if level >= 4 {
		fmt.Fprintf(os.Stdout, "[%s][DEBUG]: %s\n", getTimeNow(), fmt.Sprintf(msg, i...))
	}
}

func Info(msg string, i ...interface{}) {
	if level >= 2 {
		fmt.Fprintf(os.Stdout, "[%s][INFO]: %s\n", getTimeNow(), fmt.Sprintf(msg, i...))
	}
}

func Warning(msg string, i ...interface{}) {
	if level >= 3 {
		fmt.Fprintf(os.Stdout, "[%s][WARNING]: %s\n", getTimeNow(), fmt.Sprintf(msg, i...))
	}
}

func Error(err interface{}, i ...interface{}) {
	if level >= 1 {
		switch v := err.(type) {
		case error:
			fmt.Fprintf(os.Stderr, "[%s][ERROR]: %s\n", getTimeNow(), v.Error())
		case string:
			fmt.Fprintf(os.Stderr, "[%s][ERROR]: %s\n", getTimeNow(), fmt.Sprintf(v, i...))
		case struct{}:
			fmt.Fprintf(os.Stderr, "[%s][ERROR]: %s\n", getTimeNow(), v)
		default:
			fmt.Fprintf(os.Stderr, "[%s][ERROR]: %s\n", getTimeNow(), v)
		}
	}
}

func getTimeNow() string {
	return time.Now().Format(timeFormat)
}

func removeNextRowSymbols(st string) string {
	return regexp.MustCompile(`\r?\n`).ReplaceAllString(st, "")
}