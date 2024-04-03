package auth

import (
	"context"
	"github.com/Jumaniyozov/go-rest-template/internal/config"
	"github.com/Jumaniyozov/go-rest-template/internal/models"
	"github.com/Jumaniyozov/go-rest-template/internal/repository/auth"
	"github.com/rs/zerolog"
)

type AuthI interface {
	AllPermissions(ctx context.Context, userID int) ([]models.Permissions, error)
}

type authService struct {
	cfg  *config.Config
	log  *zerolog.Logger
	repo auth.AuthI
}

func New(c *config.Config, l *zerolog.Logger, r auth.AuthI) AuthI {
	return &authService{
		cfg:  c,
		log:  l,
		repo: r,
	}
}

func (u *authService) AllPermissions(ctx context.Context, userID int) ([]models.Permissions, error) {
	permissions, err := u.repo.AllPermissions(ctx, userID)
	if err != nil {
		return nil, err
	}
	u.log.Info().Msg("ListingAllUsers")

	return permissions, nil
}
