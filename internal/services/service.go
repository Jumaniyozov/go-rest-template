package service

import (
	"github.com/Jumaniyozov/go-rest-template/internal/config"
	contractRepository "github.com/Jumaniyozov/go-rest-template/internal/contracts/repository"
	contractService "github.com/Jumaniyozov/go-rest-template/internal/contracts/service"
	"github.com/Jumaniyozov/go-rest-template/internal/repository"
	"github.com/Jumaniyozov/go-rest-template/internal/services/auth"
	"github.com/Jumaniyozov/go-rest-template/internal/services/user"
	"github.com/rs/zerolog"
)

type service struct {
	cfg *config.Config
	log *zerolog.Logger
	rep repository.RepositoryI
}

func NewService(cfg *config.Config, log *zerolog.Logger, rep repository.RepositoryI) contractService.ServiceI {
	return &service{
		cfg: cfg,
		log: log,
		rep: rep,
	}
}

func (s *service) UserService() contractRepository.UserI {
	return user.NewService(s.cfg, s.log, s.rep.UserRepository())
}
func (s *service) AuthService() contractRepository.AuthI {
	return auth.NewService(s.cfg, s.log, s.rep.AuthRepository())
}
