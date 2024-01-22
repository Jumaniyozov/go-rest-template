package user

import (
	"context"
	"github.com/Jumaniyozov/go-rest-template/internal/contracts/repository"
	db "github.com/Jumaniyozov/go-rest-template/internal/database/sqlc"
	"time"
)

type uRepository struct {
	q *db.Queries
}

func NewRepository(q *db.Queries) repository.UserI {
	return &uRepository{
		q: q,
	}
}

func (u *uRepository) ListAllUsers() ([]db.ListUsersRow, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	users, err := u.q.ListUsers(ctx, db.ListUsersParams{
		Offset: 0,
		Limit:  100,
	})
	if err != nil {
		return nil, err
	}

	return users, nil
}
