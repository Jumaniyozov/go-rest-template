package user

import (
	"context"
	"github.com/Jumaniyozov/go-rest-template/internal/database/entities"
	db "github.com/Jumaniyozov/go-rest-template/internal/database/sqlc"
	"github.com/Jumaniyozov/go-rest-template/internal/models"
	"github.com/Jumaniyozov/go-rest-template/internal/repository"
	"github.com/Jumaniyozov/go-rest-template/internal/repository/user/converter"
)

type repo struct {
	entity *entities.Entities
}

func New(e *entities.Entities) repository.User {
	return &repo{
		entity: e,
	}
}

func (u *repo) List(ctx context.Context) ([]*models.User, error) {
	repoU, err := u.entity.User.ListUsers(ctx, db.ListUsersParams{
		Offset: 0,
		Limit:  100,
	})
	if err != nil {
		return nil, err
	}

	users := make([]*models.User, len(repoU))

	for _, v := range repoU {
		users = append(users, converter.ToUserFromRepo(&v))
	}

	return users, nil
}
