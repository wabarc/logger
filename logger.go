// Copyright 2021 Wayback Archiver. All rights reserved.
// Use of this source code is governed by the GNU GPL v3
// license that can be found in the LICENSE file.

package logger // import "github.com/wabarc/logger"

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/fatih/color"
)

var logLevel = LevelInfo
var showTime = true

// LogLevel type.
type LogLevel uint32

const (
	// LevelFatal should be used in fatal situations, the app will exit.
	LevelFatal LogLevel = iota

	// LevelError should be used when someone should really look at the error.
	LevelError

	// LevelWarn should be used when some logic on failure.
	LevelWarn

	// LevelInfo should be used during normal operations.
	LevelInfo

	// LevelDebug should be used only during development.
	LevelDebug
)

var colorable = map[LogLevel]string{
	LevelFatal: color.RedString("%s",LevelFatal),
	LevelError: color.HiRedString("%s",LevelError),
	LevelWarn:  color.YellowString("%s", LevelWarn),
	LevelInfo:  color.BlueString("%s", LevelInfo),
	LevelDebug: color.WhiteString("%s", LevelDebug),
}

func (l LogLevel) String() string {
	switch l {
	case LevelDebug:
		return "DEBUG"
	case LevelInfo:
		return "INFO"
	case LevelWarn:
		return "WARN"
	case LevelError:
		return "ERROR"
	case LevelFatal:
		return "FATAL"
	default:
		return "UNKNOWN"
	}
}

// DisableTime hides time in log messages.
func DisableTime() {
	showTime = false
}

// EnableDebug increases logging, more verbose (debug)
func EnableDebug() {
	logLevel = LevelDebug
	logging(LevelInfo, "Debug mode enabled")
}

// SetLogLevel set the log level
func SetLogLevel(l LogLevel) {
	logLevel = l
}

// Debug sends a debug log message.
func Debug(format string, v ...interface{}) {
	if logLevel >= LevelDebug {
		logging(LevelDebug, format, v...)
	}
}

// Info sends an info log message.
func Info(format string, v ...interface{}) {
	if logLevel >= LevelInfo {
		logging(LevelInfo, format, v...)
	}
}

// Warn sends a warn log message.
func Warn(format string, v ...interface{}) {
	if logLevel >= LevelWarn {
		logging(LevelWarn, format, v...)
	}
}

// Error sends an error log message.
func Error(format string, v ...interface{}) {
	if logLevel >= LevelError {
		logging(LevelError, format, v...)
	}
}

// Fatal sends a fatal log message and stop the execution of the program.
func Fatal(format string, v ...interface{}) {
	if logLevel >= LevelFatal {
		logging(LevelFatal, format, v...)
		os.Exit(1)
	}
}

func logging(l LogLevel, format string, v ...interface{}) {
	var prefix string

	if showTime {
		prefix = fmt.Sprintf("[%s] [%s] ", color.CyanString(time.Now().Format("2006-01-02T15:04:05")), colorable[l])
	} else {
		prefix = fmt.Sprintf("[%s] ", colorable[l])
	}

	pc, file, line, _ := runtime.Caller(2)
	files := strings.Split(file, "/")
	file = files[len(files)-1]
	name := runtime.FuncForPC(pc).Name()
	fns := strings.Split(name, ".")
	name = fns[len(fns)-1]
	caller := fmt.Sprintf("[%s:%d:%s] ", color.MagentaString("%s", file), line, color.MagentaString("%s", name))

	fmt.Fprintf(os.Stderr, prefix+caller+format+"\n", v...)
}
