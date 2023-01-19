package log

import (
	"fmt"
	"io"
	"os"
	"time"
)

const (
	levelDebug uint8 = iota + 1
	levelInfo
	levelWarning
	levelError

	DebugLevel   = "debug"
	InfoLevel    = "info"
	WarningLevel = "warning"
	ErrorLevel   = "error"

	_minLvl = levelDebug
	_maxLvl = levelError
)

const (
	JsonOutput = iota
	TextOutput
)

type Log struct{}

func New() *Log {
	settings = logSettings{
		os.Stdout,
		levelDebug,
		time.StampMicro,
		TextOutput}
	return &Log{}
}

type logSettings struct {
	Output       io.Writer
	Level        uint8
	TimeFormat   string
	OutputFormat int
}

type jsonOutputFormat struct {
	Writer io.Writer   `json:"writer"`
	Time   string      `json:"time"`
	Level  string      `json:"level"`
	Data   interface{} `json:"data"`
}

var settings logSettings

// levelOut return standard file descriptor and text level
func levelOut(l uint8) (io.Writer, string) {
	switch l {
	case levelDebug:
		return os.Stdout, DebugLevel
	case levelInfo:
		return os.Stdout, InfoLevel
	case levelWarning:
		return os.Stdout, WarningLevel
	case levelError:
		return os.Stderr, ErrorLevel
	default:
		return os.Stdout, DebugLevel
	}
}

// SetLogLevel set global log level
func (lo *Log) SetLogLevel(lvl uint8) *Log {
	if lvl >= _minLvl && lvl <= _maxLvl {
		settings.Level = lvl
	} else {
		settings.Level = levelDebug
	}

	return lo
}

func (lo *Log) SetTimeFormat(f string) *Log {
	settings.TimeFormat = f
	return lo
}

func (lo *Log) Error(err interface{}, i ...interface{}) {
	switch v := err.(type) {
	case error:
		lo.fPrintLog(levelError, v.Error())
	case string:
		lo.fPrintLog(levelError, fmt.Sprintf(v, i...))
	default:
		lo.fPrintLog(levelError, v)
	}
}

func (lo *Log) Warning(msg string, i ...interface{}) {
	lo.fPrintLog(levelWarning, fmt.Sprintf(msg, i...))
}

func (lo *Log) Info(msg string, i ...interface{}) {
	lo.fPrintLog(levelInfo, fmt.Sprintf(msg, i...))
}

func (lo *Log) Debug(msg string, i ...interface{}) {
	lo.fPrintLog(levelDebug, fmt.Sprintf(msg, i...))
}

func getTimeNow() string {
	return time.Now().Format(settings.TimeFormat)
}

func (lo *Log) fPrintLog(cl uint8, s interface{}) {
	if settings.Level < cl {
		return
	}

	switch settings.OutputFormat {
	case TextOutput:
		printText(cl, s)
	case JsonOutput:
		printJson(cl, s)
	default:
		printText(cl, s)
	}
}

func printText(cl uint8, s interface{}) {
	o, c := levelOut(cl)
	_, _ = fmt.Fprintf(o, "[%s][%s]: %s\n", getTimeNow(), c, s)
}

func printJson(cl uint8, s interface{}) {
	o, c := levelOut(cl)
	_, _ = fmt.Fprintf(o, fmt.Sprintf(
		`{"writer": "%s", "time": "%s", "level": "%s", "data": "%s"}`, o, getTimeNow(), c, s))
}
