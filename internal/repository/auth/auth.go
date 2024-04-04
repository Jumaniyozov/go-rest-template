package auth

import (
	"context"
	"github.com/Jumaniyozov/go-rest-template/internal/database/entities"
	"github.com/Jumaniyozov/go-rest-template/internal/models"
)

type Auth interface {
	AllPermissions(ctx context.Context, userID int) ([]models.Permissions, error)
}

type repo struct {
	entity *entities.Entities
}

func New(e *entities.Entities) Auth {
	return &repo{
		entity: e,
	}
}

func (u *repo) AllPermissions(ctx context.Context, userID int) ([]models.Permissions, error) {
	permissions := []models.Permissions{
		{
			Permission: "read",
		},
		{
			Permission: "write",
		},
	}

	return permissions, nil
}
