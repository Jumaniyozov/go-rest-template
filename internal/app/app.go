package app

import (
	"context"
	"errors"
	"fmt"
	"github.com/Jumaniyozov/go-rest-template/internal/app/routes"
	"github.com/Jumaniyozov/go-rest-template/internal/config"
	"github.com/Jumaniyozov/go-rest-template/internal/database/postgres"
	loggerpkg "github.com/Jumaniyozov/go-rest-template/internal/logger"
	service "github.com/Jumaniyozov/go-rest-template/internal/services"
	"github.com/Jumaniyozov/go-rest-template/pkg/closer"
	"net/http"
	"time"
)

const (
	shutdownTimeout = 10 * time.Second
)

// StartApp starts the server
func Start(ctx context.Context) error {

	// Setting up configurations from environment variables
	cfg, err := config.New()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	logger := loggerpkg.New(cfg)

	// Initializing repositories
	rep, err := postgres.New(cfg)
	if err != nil {
		logger.Fatal().Err(err).Msgf("Failed to connect to database %v", err)
	}

	// Initializing and setting up services
	services := service.New(cfg, logger, rep)

	// Initializing and setting up router
	router := routes.New(cfg, logger, services)

	clsr := &closer.Closer{}

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.ServerPort),
		Handler:      router.CreateHttpRouter(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 20 * time.Second,
	}

	clsr.Add(srv.Shutdown)

	// Cleanup function
	clsr.Add(func(ctx context.Context) error {
		//time.Sleep(6 * time.Second)

		return nil
	})

	go func() {
		if err = srv.ListenAndServe(); err != nil && !errors.Is(http.ErrServerClosed, err) {
			logger.Fatal().Err(err).Msgf("Failed to start server %v", err)
		}
	}()

	logger.Info().Msgf("Listening server on port %d", cfg.ServerPort)
	<-ctx.Done()
	logger.Info().Msg("Shutting down server gracefully")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	if err = clsr.Close(shutdownCtx); err != nil {
		return fmt.Errorf("closer: %v", err)
	}

	return nil
}
