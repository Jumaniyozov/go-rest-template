package repository

import (
	"github.com/Jumaniyozov/go-rest-template/internal/database/entities"
	"github.com/Jumaniyozov/go-rest-template/internal/repository/auth"
	"github.com/Jumaniyozov/go-rest-template/internal/repository/user"
)

type Repository struct {
	User user.User
	Auth auth.Auth
}

func New(e *entities.Entities) *Repository {
	return &Repository{
		User: user.New(e),
		Auth: auth.New(e),
	}
}
