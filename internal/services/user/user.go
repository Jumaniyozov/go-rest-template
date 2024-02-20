package user

import (
	"github.com/Jumaniyozov/go-rest-template/internal/config"
	db "github.com/Jumaniyozov/go-rest-template/internal/database/sqlc"
	"github.com/Jumaniyozov/go-rest-template/internal/repository/user"
	"github.com/rs/zerolog"
)

type UserI interface {
	List() ([]db.ListRow, error)
}

type userService struct {
	cfg  *config.Config
	log  *zerolog.Logger
	repo user.UserI
}

func New(c *config.Config, l *zerolog.Logger, r user.UserI) UserI {
	return &userService{
		cfg:  c,
		log:  l,
		repo: r,
	}
}

func (u *userService) List() ([]db.ListRow, error) {
	users, err := u.repo.List()
	if err != nil {
		return nil, err
	}

	return users, nil
}
