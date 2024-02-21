package logger

import (
	"fmt"
	"github.com/Jumaniyozov/go-rest-template/internal/config"
	"github.com/rs/zerolog"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type Logger struct {
	log *zerolog.Logger
}

func New(cfg *config.Config) *zerolog.Logger {
	var logger zerolog.Logger
	if cfg.ProdEnv == "dev" {
		logger = zerolog.New(zerolog.ConsoleWriter{
			Out:        os.Stderr,
			TimeFormat: time.RFC3339,
			FormatLevel: func(i interface{}) string {
				return strings.ToUpper(fmt.Sprintf("[%s]", i))
			},
			FormatMessage: func(i interface{}) string {
				return fmt.Sprintf("| %s |", i)
			},
			FormatCaller: func(i interface{}) string {
				return filepath.Base(fmt.Sprintf("%s", i))
			},
		}).
			Level(zerolog.TraceLevel).
			With().
			Timestamp().
			Caller().
			Int("pid", os.Getpid()).
			Logger()
	} else {
		logger = zerolog.New(os.Stdout).Level(zerolog.InfoLevel).With().Timestamp().Logger()
	}

	return &logger
}
