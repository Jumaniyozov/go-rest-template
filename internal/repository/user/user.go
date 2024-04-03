package user

import (
	"context"
	"github.com/Jumaniyozov/go-rest-template/internal/database/entities"
	db "github.com/Jumaniyozov/go-rest-template/internal/database/sqlc"
)

type UserI interface {
	List(ctx context.Context) ([]db.ListRow, error)
}

type repository struct {
	entity *entities.Entities
}

func New(e *entities.Entities) UserI {
	return &repository{
		entity: e,
	}
}

func (u *repository) List(ctx context.Context) ([]db.ListRow, error) {
	users, err := u.entity.User.List(ctx, db.ListParams{
		Offset: 0,
		Limit:  100,
	})
	if err != nil {
		return nil, err
	}

	return users, nil
}
