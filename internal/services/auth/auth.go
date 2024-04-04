package auth

import (
	"context"
	"github.com/Jumaniyozov/go-rest-template/internal/models"
	"github.com/Jumaniyozov/go-rest-template/internal/repository/auth"
)

type Auth interface {
	AllPermissions(ctx context.Context, userID int) ([]models.Permissions, error)
}

type authService struct {
	repo auth.Auth
}

func New(r auth.Auth) Auth {
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
