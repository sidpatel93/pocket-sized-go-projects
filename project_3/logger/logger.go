package logger

import (
	"fmt"
	"io"
	"os"
)

// Logger represents a simple logger with different log levels.
type Logger struct {
	threshold Level
	output    io.Writer
}

// New creates a new Logger with the given threshold level and output writer.
func New(threshold Level, opts ...Option) *Logger {
	logger := &Logger{
		threshold: threshold,
		output:    os.Stdout,
	}
	for _, configOptions := range opts {
		configOptions(logger)
	}
	return logger
}

// Debugf formats and prints a message if the log level is debug or higher.
func (l *Logger) Debugf(format string, args ...any) {
	// implement me
	if l.output == nil {
		l.output = os.Stdout
	}
	if l.threshold > LevelDebug {
		return
	}
	l.logf(format, args...)
}

// Infof formats and prints a message if the log level is info or higher.
func (l *Logger) Infof(format string, args ...any) {
	// implement me
	if l.output == nil {
		l.output = os.Stdout
	}
	if l.threshold > LevelInfo {
		return
	}
	l.logf(format, args...)
}

func (l *Logger) Errorf(format string, args ...any) {
	// implement me
	if l.output == nil {
		l.output = os.Stdout
	}
	if l.threshold > LevelError {
		return
	}
	l.logf(format, args...)
}

func (l *Logger) logf(format string, args ...any) {
	_, _ = fmt.Fprintf(l.output, format+"\n", args...)
}
