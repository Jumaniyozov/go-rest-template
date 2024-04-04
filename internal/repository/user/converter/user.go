package converter

import (
	db "github.com/Jumaniyozov/go-rest-template/internal/database/sqlc"
	"github.com/Jumaniyozov/go-rest-template/internal/models"
)

func ToUserFromRepo(row *db.ListRow) *models.User {
	return &models.User{}
}
