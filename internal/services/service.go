package service

import (
	"github.com/Jumaniyozov/go-rest-template/internal/repository"
	"github.com/Jumaniyozov/go-rest-template/internal/services/auth"
	"github.com/Jumaniyozov/go-rest-template/internal/services/user"
)

type Service struct {
	User user.User
	Auth auth.Auth
}

func New(repository *repository.Repository) *Service {
	return &Service{
		User: user.New(repository.User),
		Auth: auth.New(repository.Auth),
	}
}
