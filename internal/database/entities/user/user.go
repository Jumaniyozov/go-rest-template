package user

import (
	"context"
	db "github.com/Jumaniyozov/go-rest-template/internal/database/sqlc"
)

type User interface {
	ListUsers(ctx context.Context, arg db.ListUsersParams) ([]db.ListUsersRow, error)
}
