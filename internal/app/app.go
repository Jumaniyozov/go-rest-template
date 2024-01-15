package app

import (
	"fmt"
	"github.com/Jumaniyozov/go-rest-template/internal/config"
	"github.com/Jumaniyozov/go-rest-template/internal/database/postgres"
	loggerpkg "github.com/Jumaniyozov/go-rest-template/internal/logger"
	service "github.com/Jumaniyozov/go-rest-template/internal/services"
	"net/http"
	"time"
)

// StartApp starts the server
func StartApp() {
	// Initializing logger
	logger := loggerpkg.SetupLoggger()

	// Setting up configurations from environment variables
	cfg, err := config.SetupConfigs(logger)
	if err != nil {
		logger.Fatal().Err(err).Msgf("Failed to setup configs %v", err)
	}

	// Initializing repositories
	rep, err := postgres.NewPostgresDB(cfg)
	if err != nil {
		logger.Fatal().Err(err).Msgf("Failed to connect to database %v", err)
	}

	// Initializing and setting up services
	services := service.NewService(cfg, logger, rep)

	// Initializing and setting up router
	router := SetupoRouter(cfg, logger, services)

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.ServerPort),
		Handler:      router,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 20 * time.Second,
	}

	logger.Info().Msgf("Starting server on port %d", cfg.ServerPort)

	if err = srv.ListenAndServe(); err != nil {
		logger.Fatal().Err(err).Msgf("Failed to start server %v", err)
	}
}
