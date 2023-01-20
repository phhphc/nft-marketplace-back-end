package log

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Logger struct {
	zerolog.Logger
}

var (
	logger *Logger
)

func init() {
	logger = &Logger{
		Logger: log.Logger,
	}
}

func SetPrettyLogging() {
	logger.Logger = logger.Output(zerolog.ConsoleWriter{Out: os.Stderr})
}

func GetLogger() *Logger {
	return logger
}
