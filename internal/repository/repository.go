package repository

import (
	"github.com/Jumaniyozov/go-rest-template/internal/contracts/repository"
)

type RepositoryI interface {
	UserRepository() repository.UserI
	AuthRepository() repository.AuthI
}
