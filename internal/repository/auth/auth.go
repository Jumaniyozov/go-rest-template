package auth

import (
	db "github.com/Jumaniyozov/go-rest-template/internal/database/sqlc"
	"github.com/Jumaniyozov/go-rest-template/internal/models"
)

type AuthI interface {
	GetAllPermissions(userID int) ([]models.Permissions, error)
}

type aRepository struct {
	q *db.Queries
}

func NewRepository(q *db.Queries) AuthI {
	return &aRepository{
		q: q,
	}
}

func (u *aRepository) GetAllPermissions(userID int) ([]models.Permissions, error) {
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
