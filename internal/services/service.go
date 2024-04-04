package service

import (
	"context"
	"github.com/Jumaniyozov/go-rest-template/internal/models"
)

type ServiceI struct {
	User User
	Auth Auth
}

func New(user User, auth Auth) *ServiceI {
	return &ServiceI{
		User: user,
		Auth: auth,
	}
}

type User interface {
	List(ctx context.Context) ([]*models.User, error)
}

type Auth interface {
	AllPermissions(ctx context.Context, userID int) ([]models.Permissions, error)
}
