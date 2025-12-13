package logger_test

import (
	"os"

	"github.com/sidpatel93/pocket-sized-go-projects/project_3/logger"
	// "testing"
)

func ExampleLogger_Debugf() {
	debugLogger := logger.New(logger.LevelDebug, logger.WithOutput(os.Stdout))
	debugLogger.Debugf("Hello, %s", "world")
	// Output: Hello, world
}
