package logger_test

import (
	"github.com/sidpatel93/pocket-sized-go-projects/project_3/logger"
	// "testing"
)

func ExampleLogger_Debugf() {
	debugLogger := logger.New(logger.LevelDebug)
	debugLogger.Debugf("Hello, %s", "world")
	// Output: Hello, world
}
