package logger

type Logger struct {
}

// Level represents the available logger levels.
type Level byte

// Available logger levels enum.
const (
	LevelDebug Level = iota
	LevelInfo
	LevelWarn
	LevelError
)

// Debugf formats and prints a message if the log level is debug or higher.
func (l *Logger) Debugf(format string, args ...any) {
	// implement me
}

// Infof formats and prints a message if the log level is info or higher.
func (l *Logger) Infof(format string, args ...any) {
	// implement me
}
