// Copyright 2012-2015 Apcera Inc. All rights reserved.

//Package logger provides logging facilities for the NATS server
package logger

import (
	"fmt"
	"log/writer"
	"os"
)

// Logger is the server logger
type Logger struct {
	logger     writer.Writer
	debug      bool
	trace      bool
	warnLabel  string
	errorLabel string
	fatalLabel string
	debugLabel string
	traceLabel string
}

// NewStdLogger creates a logger with output directed to Stderr
func NewStdLogger(path string, debug, trace, colors bool) *Logger {
	l := &Logger{
		logger: writer.NewWriter(path, 1<<20, 0),
		debug:  debug,
		trace:  trace,
	}

	if colors {
		setColoredLabelFormats(l)
	} else {
		setPlainLabelFormats(l)
	}

	return l
}

func (l *Logger) Close() {
	l.logger.Close()
}

func setPlainLabelFormats(l *Logger) {
	l.debugLabel = "[DBG] "
	l.traceLabel = "[TRC] "
	l.warnLabel = "[WAR] "
	l.errorLabel = "[ERR] "
	l.fatalLabel = "[FTL] "
}

func setColoredLabelFormats(l *Logger) {
	colorFormat := "[\x1b[%dm%s\x1b[0m] "
	l.debugLabel = fmt.Sprintf(colorFormat, 36, "DBG")
	l.traceLabel = fmt.Sprintf(colorFormat, 33, "TRC")
	l.warnLabel = fmt.Sprintf(colorFormat, 32, "WAR")
	l.errorLabel = fmt.Sprintf(colorFormat, 31, "ERR")
	l.fatalLabel = fmt.Sprintf(colorFormat, 31, "FTL")
}

// Debug logs a debug statement
func (l *Logger) Debug(format string, v ...interface{}) {
	if l.debug {
		l.logger.Write([]byte(l.getContent(nil, format, l.debugLabel, v...)))
	}
}

func (l *Logger) DebugWithField(head map[string]interface{}, format string, v ...interface{}) {
	if l.debug {
		l.logger.Write([]byte(l.getContent(head, format, l.debugLabel, v...)))
	}
}

// Trace logs a trace statement
func (l *Logger) Trace(format string, v ...interface{}) {
	if l.trace {
		l.logger.Write([]byte(l.getContent(nil, format, l.traceLabel, v...)))
	}
}

func (l *Logger) TraceWithField(head map[string]interface{}, format string, v ...interface{}) {
	if l.trace {
		l.logger.Write([]byte(l.getContent(head, format, l.traceLabel, v...)))
	}
}

// Warning logs a notice statement
func (l *Logger) Warning(format string, v ...interface{}) {
	l.logger.Write([]byte(l.getContent(nil, format, l.warnLabel, v...)))
}

// Warning logs a notice statement
func (l *Logger) WarningWithField(head map[string]interface{}, format string, v ...interface{}) {
	l.logger.Write([]byte(l.getContent(head, format, l.warnLabel, v...)))
}

// Error logs an error statement
func (l *Logger) Error(format string, v ...interface{}) {
	l.logger.Write([]byte(l.getContent(nil, format, l.errorLabel, v...)))
}

func (l *Logger) ErrorWithField(head map[string]interface{}, format string, v ...interface{}) {
	l.logger.Write([]byte(l.getContent(head, format, l.errorLabel, v...)))
}

// Fatal logs a fatal error
func (l *Logger) Fatal(format string, v ...interface{}) {
	l.logger.Write([]byte(l.getContent(nil, format, l.fatalLabel, v...)))
	os.Exit(1)
}

func (l *Logger) FatalWithField(head map[string]interface{}, format string, v ...interface{}) {
	l.logger.Write([]byte(l.getContent(head, format, l.fatalLabel, v...)))
	os.Exit(1)
}

func (l *Logger) getContent(head map[string]interface{}, format, label string, v ...interface{}) string {
	if len(head) > 0 {
		return fmt.Sprintf("%s %+v %s%s", label, head, fmt.Sprintf(format, v...), fmt.Sprintln())
	}
	return fmt.Sprintf("%s %s%s", label, fmt.Sprintf(format, v...), fmt.Sprintln())
}
