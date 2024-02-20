package user

import (
	"context"
	db "github.com/Jumaniyozov/go-rest-template/internal/database/sqlc"
)

type User interface {
	List(ctx context.Context, arg db.ListParams) ([]db.ListRow, error)
}
