package app

import (
	"context"
	"errors"
	"fmt"
	"github.com/Jumaniyozov/go-rest-template/internal/app/routes"
	"github.com/Jumaniyozov/go-rest-template/internal/config"
	"github.com/Jumaniyozov/go-rest-template/internal/database/postgres"
	loggerpkg "github.com/Jumaniyozov/go-rest-template/internal/logger"
	"github.com/Jumaniyozov/go-rest-template/internal/repository"
	service "github.com/Jumaniyozov/go-rest-template/internal/services"
	"github.com/Jumaniyozov/go-rest-template/pkg/closer"
	"log"
	"net/http"
	"time"
)

const (
	shutdownTimeout = 10 * time.Second
)

type App struct {
	serviceProvider *service.Service
	httpServer      *http.Server
	repository      *repository.Repository
	config          *config.Config
	logger          *loggerpkg.Logger
}

func New() *App {
	a := &App{}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := a.initDeps(ctx)
	if err != nil {
		log.Fatalf("failed to initialize dependencies: %v", err)
	}

	return a
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initConfig,
		a.initLogger,
		a.initServices,
		a.initRepository,
		a.initHTTPServer,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initConfig(_ context.Context) error {
	if a.config == nil {
		cfg, err := config.New()
		if err != nil {
			return fmt.Errorf("config: %v", err)
		}

		a.config = cfg
	}

	return nil
}

func (a *App) initLogger(_ context.Context) error {
	if a.logger == nil {
		logger := loggerpkg.New(a.config)
		a.logger = logger
	}

	return nil
}

func (a *App) initRepository(_ context.Context) error {
	if a.repository == nil {
		rep, err := postgres.New(a.config)
		if err != nil {
			return fmt.Errorf("repository: %v", err)
		}

		a.repository = rep
	}

	return nil
}

func (a *App) initServices(_ context.Context) error {
	if a.serviceProvider == nil {
		services := service.New(a.repository)
		a.serviceProvider = services
	}

	return nil
}

func (a *App) initHTTPServer(_ context.Context) error {
	if a.httpServer == nil {
		router := routes.New(a.config, a.logger, a.serviceProvider)

		a.httpServer = &http.Server{
			Addr:         fmt.Sprintf(":%d", a.config.ServerPort),
			Handler:      router.CreateHttpRouter(),
			IdleTimeout:  time.Minute,
			ReadTimeout:  20 * time.Second,
			WriteTimeout: 30 * time.Second,
		}
	}

	return nil

}

func (a *App) Run(ctx context.Context) error {
	logger := a.logger.Logger

	clsr := &closer.Closer{}

	clsr.Add(a.httpServer.Shutdown)

	// Cleanup function
	clsr.Add(func(ctx context.Context) error {
		//time.Sleep(6 * time.Second)

		return nil
	})

	go func() {
		if err := a.httpServer.ListenAndServe(); err != nil && !errors.Is(http.ErrServerClosed, err) {
			logger.Fatal().Err(err).Msgf("failed to start server %v", err)
		}
	}()

	logger.Info().Msgf("Listening server on port %d", a.config.ServerPort)
	<-ctx.Done()
	logger.Info().Msg("Shutting down server gracefully")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	if err := clsr.Close(shutdownCtx); err != nil {
		return fmt.Errorf("closer: %v", err)
	}

	return nil
}
