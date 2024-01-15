package auth

import (
	"github.com/Jumaniyozov/go-rest-template/internal/config"
	"github.com/Jumaniyozov/go-rest-template/internal/repository/auth"
	"github.com/rs/zerolog"
)

type authService struct {
	cfg  *config.Config
	log  *zerolog.Logger
	repo auth.RAuthI
}

type AuthServiceI interface {
	GetAllPermissions(userID int) ([]auth.Permissions, error)
}

func NewService(c *config.Config, l *zerolog.Logger, r auth.RAuthI) AuthServiceI {
	return &authService{
		cfg:  c,
		log:  l,
		repo: r,
	}
}

func (u *authService) GetAllPermissions(userID int) ([]auth.Permissions, error) {
	permissions, err := u.repo.GetAllPermissions(userID)
	if err != nil {
		return nil, err
	}
	u.log.Info().Msg("ListingAllUsers")

	return permissions, nil
}
