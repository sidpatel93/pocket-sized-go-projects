package logger

import "io"

type Option func(*Logger)

func WithOutput(output io.Writer) Option {
	return func(logger *Logger) {
		logger.output = output
	}
}
