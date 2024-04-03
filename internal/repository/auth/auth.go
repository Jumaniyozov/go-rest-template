package auth

import (
	"context"
	"github.com/Jumaniyozov/go-rest-template/internal/database/entities"
	"github.com/Jumaniyozov/go-rest-template/internal/models"
)

type AuthI interface {
	AllPermissions(ctx context.Context, userID int) ([]models.Permissions, error)
}

type repository struct {
	entity *entities.Entities
}

func New(e *entities.Entities) AuthI {
	return &repository{
		entity: e,
	}
}

func (u *repository) AllPermissions(ctx context.Context, userID int) ([]models.Permissions, error) {
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
