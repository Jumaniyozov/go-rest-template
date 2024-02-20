package entities

import (
	"github.com/Jumaniyozov/go-rest-template/internal/database/entities/user"
	db "github.com/Jumaniyozov/go-rest-template/internal/database/sqlc"
)

type Entities struct {
	User user.User
}

func New(q *db.Queries) *Entities {
	return &Entities{
		User: q,
	}
}
