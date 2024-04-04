package user

import (
	"context"
	"github.com/Jumaniyozov/go-rest-template/internal/database/entities"
	db "github.com/Jumaniyozov/go-rest-template/internal/database/sqlc"
	"github.com/Jumaniyozov/go-rest-template/internal/models"
	"github.com/Jumaniyozov/go-rest-template/internal/repository/user/converter"
)

type User interface {
	List(ctx context.Context) ([]*models.User, error)
}

type repo struct {
	entity *entities.Entities
}

func New(e *entities.Entities) User {
	return &repo{
		entity: e,
	}
}

func (u *repo) List(ctx context.Context) ([]*models.User, error) {
	repoUsers, err := u.entity.User.List(ctx, db.ListParams{
		Offset: 0,
		Limit:  100,
	})
	if err != nil {
		return nil, err
	}

	users := make([]*models.User, len(repoUsers))

	for _, v := range repoUsers {
		users = append(users, converter.ToUserFromRepo(&v))
	}

	return users, nil
}
