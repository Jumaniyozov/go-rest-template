package repository

import (
	"github.com/Jumaniyozov/go-rest-template/internal/repository/auth"
	"github.com/Jumaniyozov/go-rest-template/internal/repository/user"
)

type RepositoryI interface {
	UserRepository() user.UserI
	AuthRepository() auth.AuthI
}
