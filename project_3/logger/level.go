package logger

// Level represents the available logger levels.
type Level byte

// Available logger levels enum.
const (
	LevelDebug Level = iota
	LevelInfo
	LevelWarn
	LevelError
)
