package logger

import (
	"io"
	"log"
	"os"

	"github.com/fatih/color"
)

type Logger struct {
	debug   *log.Logger
	info    *log.Logger
	warning *log.Logger
	err     *log.Logger
	writer  io.Writer
}

// NewLogger function  
func NewLogger() *Logger {
	writer := io.Writer(os.Stdout)
	logger := log.New(writer, "", log.Ldate|log.Ltime)

	cyan := color.New(color.FgCyan).SprintfFunc()
	green := color.New(color.FgGreen).SprintfFunc()
	yellow := color.New(color.FgYellow).SprintfFunc()
	red := color.New(color.FgRed).SprintfFunc()

	return &Logger{
		debug:   log.New(writer, cyan("|> DEBUG: "), logger.Flags()),
		info:    log.New(writer, green("|> INFO: "), logger.Flags()),
		warning: log.New(writer, yellow("|> WARNING: "), logger.Flags()),
		err:     log.New(writer, red("|> ERROR: "), logger.Flags()),
		writer:  writer,
	}
}

// Create Non-Formatted Logs
func (l *Logger) Debug(v ...interface{}) {
	l.debug.Println(v...)
}

func (l *Logger) Info(v ...interface{}) {
	l.info.Println(v...)
}

func (l *Logger) Warn(v ...interface{}) {
	l.warning.Println(v...)
}

func (l *Logger) Error(v ...interface{}) {
	l.err.Println(v...)
}

// Create Format Enabled Logs
func (l *Logger) Debugf(format string, v ...interface{}) {
	l.debug.Printf(format, v...)
}

func (l *Logger) Infof(format string, v ...interface{}) {
	l.info.Printf(format, v...)
}

func (l *Logger) Warnf(format string, v ...interface{}) {
	l.warning.Printf(format, v...)
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	l.err.Printf(format, v...)
}
