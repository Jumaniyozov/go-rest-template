package user

import (
	"github.com/Jumaniyozov/go-rest-template/internal/config"
	db "github.com/Jumaniyozov/go-rest-template/internal/database/sqlc"
	"github.com/Jumaniyozov/go-rest-template/internal/repository"
	"github.com/rs/zerolog"
)

type userService struct {
	cfg  *config.Config
	log  *zerolog.Logger
	repo repository.UserI
}

type UserServiceI interface {
	ListAllUsers() ([]db.ListUsersRow, error)
}

func NewService(c *config.Config, l *zerolog.Logger, r repository.UserI) UserServiceI {
	return &userService{
		cfg:  c,
		log:  l,
		repo: r,
	}
}

func (u *userService) ListAllUsers() ([]db.ListUsersRow, error) {
	users, err := u.repo.ListAllUsers()
	if err != nil {
		return nil, err
	}
	u.log.Info().Msg("ListingAllUsers")

	return users, nil
}