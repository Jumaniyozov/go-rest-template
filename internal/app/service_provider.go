package app

import (
	"context"
	"github.com/Jumaniyozov/go-rest-template/internal/api/rest/user"
	"github.com/Jumaniyozov/go-rest-template/internal/config"
	"github.com/Jumaniyozov/go-rest-template/internal/database/postgres"
	"github.com/Jumaniyozov/go-rest-template/internal/repository"
	userRepository "github.com/Jumaniyozov/go-rest-template/internal/repository/user"
	service "github.com/Jumaniyozov/go-rest-template/internal/services"
	userService "github.com/Jumaniyozov/go-rest-template/internal/services/user"
	"log"
)

type serviceProvider struct {
	config           *config.Config
	dbClient         *postgres.DBClient
	entityClient     postgres.EntityClient
	userRepository   repository.User
	userService      service.User
	userImplentation *user.UserImplementation
	authRepository   repository.Auth
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) Config() *config.Config {
	if s.config == nil {
		cfg, err := config.New()
		if err != nil {
			log.Fatalf("failed to load config: %v", err)
		}
		s.config = cfg
	}

	return s.config
}

func (s *serviceProvider) DBClient(ctx context.Context) *postgres.DBClient {
	if s.dbClient == nil {
		dbpool, err := postgres.New(s.Config())
		if err != nil {
			log.Fatalf("failed to create db client: %v", err)
		}

		err = dbpool.Ping(ctx)
		if err != nil {
			log.Fatalf("ping error: %s", err.Error())
		}

		s.dbClient = postgres.NewDBEntities(dbpool)
	}

	return s.dbClient
}

func (s *serviceProvider) UserRepository(ctx context.Context) repository.User {
	if s.userRepository == nil {
		s.userRepository = userRepository.New(s.dbClient.Entity)
	}

	return s.userRepository
}

func (s *serviceProvider) UserService(ctx context.Context) service.User {
	if s.userService == nil {
		s.userService = userService.New(
			s.UserRepository(ctx),
		)
	}

	return s.userService
}

func (s *serviceProvider) NoteImpl(ctx context.Context) *user.UserImplementation {
	if s.userImplentation == nil {
		s.userImplentation = user.NewImplementation(s.UserService(ctx))
	}

	return s.userImplentation
}
