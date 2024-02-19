package service

import (
	"github.com/Jumaniyozov/go-rest-template/internal/config"
	"github.com/Jumaniyozov/go-rest-template/internal/repository"
	"github.com/Jumaniyozov/go-rest-template/internal/services/auth"
	"github.com/Jumaniyozov/go-rest-template/internal/services/user"
	"github.com/rs/zerolog"
)

type ServiceI interface {
	UserService() user.UserI
	AuthService() auth.AuthI
}

type service struct {
	cfg *config.Config
	log *zerolog.Logger
	rep repository.RepositoryI
}

func New(cfg *config.Config, log *zerolog.Logger, rep repository.RepositoryI) ServiceI {
	return &service{
		cfg: cfg,
		log: log,
		rep: rep,
	}
}

func (s *service) UserService() user.UserI {
	return user.New(s.cfg, s.log, s.rep.UserRepository())
}
func (s *service) AuthService() auth.AuthI {
	return auth.New(s.cfg, s.log, s.rep.AuthRepository())
}
