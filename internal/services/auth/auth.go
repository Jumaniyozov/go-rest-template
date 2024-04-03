package auth

import (
	"context"
	"github.com/Jumaniyozov/go-rest-template/internal/models"
	"github.com/Jumaniyozov/go-rest-template/internal/repository"
	service "github.com/Jumaniyozov/go-rest-template/internal/services"
)

type authService struct {
	repo repository.Auth
}

func New(r repository.Auth) service.Auth {
	return &authService{
		repo: r,
	}
}

func (u *authService) AllPermissions(ctx context.Context, userID int) ([]models.Permissions, error) {
	permissions, err := u.repo.AllPermissions(ctx, userID)
	if err != nil {
		return nil, err
	}

	return permissions, nil
}
