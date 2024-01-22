package auth

import (
	"github.com/Jumaniyozov/go-rest-template/internal/contracts/repository"
	db "github.com/Jumaniyozov/go-rest-template/internal/database/sqlc"
	"github.com/Jumaniyozov/go-rest-template/internal/models"
)

type aRepository struct {
	q *db.Queries
}

func NewRepository(q *db.Queries) repository.AuthI {
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
