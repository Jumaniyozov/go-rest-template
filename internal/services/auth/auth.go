package auth

import (
	"github.com/Jumaniyozov/go-rest-template/internal/config"
	"github.com/Jumaniyozov/go-rest-template/internal/contracts/repository"
	"github.com/Jumaniyozov/go-rest-template/internal/models"
	"github.com/rs/zerolog"
)

type authService struct {
	cfg  *config.Config
	log  *zerolog.Logger
	repo repository.AuthI
}

func NewService(c *config.Config, l *zerolog.Logger, r repository.AuthI) repository.AuthI {
	return &authService{
		cfg:  c,
		log:  l,
		repo: r,
	}
}

func (u *authService) GetAllPermissions(userID int) ([]models.Permissions, error) {
	permissions, err := u.repo.GetAllPermissions(userID)
	if err != nil {
		return nil, err
	}
	u.log.Info().Msg("ListingAllUsers")

	return permissions, nil
}
