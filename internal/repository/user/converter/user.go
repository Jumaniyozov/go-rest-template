package converter

import (
	db "github.com/Jumaniyozov/go-rest-template/internal/database/sqlc"
	"github.com/Jumaniyozov/go-rest-template/internal/models"
)

func ToUserFromRepo(user *db.ListUsersRow) *models.User {
	return &models.User{}
}
