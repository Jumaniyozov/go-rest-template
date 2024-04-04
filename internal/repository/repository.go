package repository

import (
	"context"
	"github.com/Jumaniyozov/go-rest-template/internal/models"
)

type RepositoryI struct {
	User User
	Auth Auth
}

type User interface {
	List(ctx context.Context) ([]*models.User, error)
}

type Auth interface {
	AllPermissions(ctx context.Context, userID int) ([]models.Permissions, error)
}
