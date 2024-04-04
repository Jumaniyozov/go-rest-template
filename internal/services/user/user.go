package user

import (
	"context"
	"github.com/Jumaniyozov/go-rest-template/internal/models"
	"github.com/Jumaniyozov/go-rest-template/internal/repository/user"
)

type User interface {
	List(ctx context.Context) ([]*models.User, error)
}

type userService struct {
	repo user.User
}

func New(r user.User) User {
	return &userService{
		repo: r,
	}
}

func (u *userService) List(ctx context.Context) ([]*models.User, error) {
	users, err := u.repo.List(ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}
