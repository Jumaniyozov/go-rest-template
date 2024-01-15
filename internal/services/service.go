package service

import (
	"github.com/Jumaniyozov/go-rest-template/internal/config"
	"github.com/Jumaniyozov/go-rest-template/internal/repository"
	"github.com/Jumaniyozov/go-rest-template/internal/services/auth"
	"github.com/Jumaniyozov/go-rest-template/internal/services/user"
	"github.com/rs/zerolog"
)

type ServiceI interface {
	UserService() user.UserServiceI
	AuthService() auth.AuthServiceI
}

type service struct {
	cfg *config.Config
	log *zerolog.Logger
	rep repository.RepositoryI
}

func NewService(cfg *config.Config, log *zerolog.Logger, rep repository.RepositoryI) ServiceI {
	return &service{
		cfg: cfg,
		log: log,
		rep: rep,
	}
}

func (s *service) UserService() user.UserServiceI {
	return user.NewService(s.cfg, s.log, s.rep.UserRepository())
}
func (s *service) AuthService() auth.AuthServiceI {
	return auth.NewService(s.cfg, s.log, s.rep.AuthRepository())
}
