package logger

import (
	"github.com/rs/zerolog"
	"os"
)

type Logger struct {
	log *zerolog.Logger
}

func New() (log *zerolog.Logger) {
	logger := zerolog.New(os.Stdout)
	return &logger
}
