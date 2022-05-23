package logger

import (
	"context"
	"io"
	"os"

	"github.com/rs/zerolog"
)

type Logger struct {
	logger *zerolog.Logger
}

func New(isDebug bool) *Logger {
	zerolog.SetGlobalLevel(getLogLevel(isDebug))
	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()
	return &Logger{logger : &logger}
}

func getLogLevel(isDebug bool) zerolog.Level {
	logLevel := zerolog.InfoLevel
	if isDebug {
		logLevel = zerolog.DebugLevel
	}
	return logLevel
}

func NewConsole (isDebug bool) *Logger {
	zerolog.SetGlobalLevel(getLogLevel(isDebug))
	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()
	return &Logger{logger : &logger}

} 

func (l *Logger) Output (w io.Writer) zerolog.Logger {
	return l.logger.Output(w)
}
