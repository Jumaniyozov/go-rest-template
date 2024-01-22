package service

import (
	"github.com/Jumaniyozov/go-rest-template/internal/contracts/repository"
)

type ServiceI interface {
	UserService() repository.UserI
	AuthService() repository.AuthI
}
