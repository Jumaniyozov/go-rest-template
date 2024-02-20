package auth

import (
	"github.com/Jumaniyozov/go-rest-template/internal/config"
	"github.com/Jumaniyozov/go-rest-template/internal/models"
	"github.com/Jumaniyozov/go-rest-template/internal/repository/auth"
	"github.com/rs/zerolog"
)

type AuthI interface {
	AllPermissions(userID int) ([]models.Permissions, error)
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

func (u *authService) AllPermissions(userID int) ([]models.Permissions, error) {
	permissions, err := u.repo.AllPermissions(userID)
	if err != nil {
		return nil, err
	}
	u.log.Info().Msg("ListingAllUsers")

	return permissions, nil
}
