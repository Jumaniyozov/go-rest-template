package user

import (
	"context"
	"github.com/Jumaniyozov/go-rest-template/internal/models"
	"github.com/Jumaniyozov/go-rest-template/internal/repository"
	service "github.com/Jumaniyozov/go-rest-template/internal/services"
)

type userService struct {
	repo repository.User
}

func New(r repository.User) service.User {
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
