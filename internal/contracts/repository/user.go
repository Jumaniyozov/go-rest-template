package repository

import db "github.com/Jumaniyozov/go-rest-template/internal/database/sqlc"

type UserI interface {
	ListAllUsers() ([]db.ListUsersRow, error)
}
