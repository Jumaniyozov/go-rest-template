package auth

import (
	db "github.com/Jumaniyozov/go-rest-template/internal/database/sqlc"
)

type aRepository struct {
	q *db.Queries
}

type RAuthI interface {
	GetAllPermissions(userID int) ([]Permissions, error)
}

type Permissions struct {
	Permission string `json:"permission"`
}

func NewRepository(q *db.Queries) RAuthI {
	return &aRepository{
		q: q,
	}
}

func (u *aRepository) GetAllPermissions(userID int) ([]Permissions, error) {
	permissions := []Permissions{
		{
			Permission: "read",
		},
		{
			Permission: "write",
		},
	}

	return permissions, nil
}
