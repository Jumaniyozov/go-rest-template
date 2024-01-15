package repository

import (
	db "github.com/Jumaniyozov/go-rest-template/internal/database/sqlc"
)

type RepositoryI interface {
	UserRepository() UserI
}

type UserI interface {
	ListAllUsers() ([]db.ListUsersRow, error)
}
