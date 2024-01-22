package repository

import "github.com/Jumaniyozov/go-rest-template/internal/models"

type AuthI interface {
	GetAllPermissions(userID int) ([]models.Permissions, error)
}
