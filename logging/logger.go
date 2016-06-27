/*
Package logging introduces a convenient logging interface and a suggested DefaultLogger struct for use in logging

Loggers should follow the syslog standard and it is this author's belief that the way this package can be used
trumps the standard log and syslog loggers.

Because golang only supports the Linux syslog implementation, the package does not force its use and defaults instead
to Stdout. However, this can be extended to utilize your system's syslog implementation or even a custom logging system
using syslog's priority standard.
*/
package logging

import (
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

var globalLogger Logger
var loggers map[string]Logger

//Level is the syslog level represented by a uint8 type
type Level uint8

//no need to generate, these will stay constant
func (l Level) String() string {
	switch l {
	case ALERT:
		return "ALERT"
	case CRITICAL:
		return "CRITICAL"
	case ERROR:
		return "ERROR"
	case WARNING:
		return "WARNING"
	case NOTICE:
		return "NOTICE"
	case INFO:
		return "INFO"
	case DEBUG:
		return "DEBUG"
	}
	return fmt.Sprintf("Level(%d)", l)
}

//These are the the logging levels.
const (
	ALERT Level = 1 << iota
	CRITICAL
	ERROR
	WARNING
	NOTICE
	INFO
	DEBUG
)

//Logger is an interface that implements Log(...) which requires a level and content.
type Logger interface {
	Log(Level, ...interface{})
}

//DefaultLogger is the suggested default logger
type DefaultLogger struct {
	sync.Mutex
	output        io.Writer
	Prefix        string
	AllowedLevels uint8
	EventCounts   map[Level]uint64
}

//NewDefaultLogger initializes and returns a DefaultLogger instance
func NewDefaultLogger(out io.Writer) *DefaultLogger {
	logger := new(DefaultLogger)
	logger.output = out
	logger.AllowedLevels = 0xff
	logger.EventCounts = make(map[Level]uint64)

	return logger
}

//Log satisfies the birdnest Logger interface
func (logger *DefaultLogger) Log(lvl Level, content ...interface{}) {
	logger.Lock()
	defer logger.Unlock()
	if uint8(lvl)&logger.AllowedLevels == 0 {
		return
	}
	logger.EventCounts[lvl]++
	logger.output.Write([]byte(fmt.Sprintf("<%s> [%s] %s - %s\n", logger.Prefix, lvl, time.Now().Format(time.RFC3339), fmt.Sprint(content...))))
}

//GetLogger retrieves a logger. If there is no tag given, then the global logger retrieved
func GetLogger(tag ...string) Logger {
	if tag == nil {
		return globalLogger
	}
	return loggers[tag[0]]
}

//RegisterLogger registers a logger. If the tag parameter is not given, then logger becomes the global default.
func RegisterLogger(logger Logger, tag ...string) {
	if tag == nil {
		globalLogger = logger
		return
	}
	if loggers == nil {
		loggers = make(map[string]Logger)
	}
	loggers[tag[0]] = logger
}

//CONVENIENCE METHODS

//Alert is a convenience method for logging using the ALERT level with the global logger
func Alert(content ...interface{}) { globalLogger.Log(ALERT, content...) }

//Critical is a convenience method for logging using the CRITICAL level with the global logger
func Critical(content ...interface{}) { globalLogger.Log(CRITICAL, content...) }

//Error is a convenience method for logging using the ERROR level with the global logger
func Error(content ...interface{}) { globalLogger.Log(ERROR, content...) }

//Warning is a convenience method for logging using the WARNING level with the global logger
func Warning(content ...interface{}) { globalLogger.Log(WARNING, content...) }

//Notice is a convenience method for logging using the NOTICE level with the global logger
func Notice(content ...interface{}) { globalLogger.Log(NOTICE, content...) }

//Info is a convenience method for logging using the INFO level with the global logger
func Info(content ...interface{}) { globalLogger.Log(INFO, content...) }

//Debug is a convenience method for logging using the DEBUG level with the global logger
func Debug(content ...interface{}) { globalLogger.Log(DEBUG, content...) }

func init() {
	//initialize a base logger
	logger := NewDefaultLogger(os.Stdout)
	logger.Prefix = "default"
	globalLogger = logger
}
