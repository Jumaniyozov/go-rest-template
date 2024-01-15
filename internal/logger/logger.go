package logger

import (
	"github.com/rs/zerolog"
	"os"
)

func SetupLoggger() *zerolog.Logger {
	logger := zerolog.New(os.Stdout)
	return &logger
}
