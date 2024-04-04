package service

import (
	"context"
	"github.com/Jumaniyozov/go-rest-template/internal/models"
	"github.com/Jumaniyozov/go-rest-template/internal/repository"
)

type Service struct {
	User User
	Auth Auth
}

func New(repository *repository.Repository) *Service {
	return &Service{
		User: repository.User,
		Auth: repository.Auth,
	}
}

type User interface {
	List(ctx context.Context) ([]*models.User, error)
}

type Auth interface {
	AllPermissions(ctx context.Context, userID int) ([]models.Permissions, error)
}
