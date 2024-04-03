package models

type User struct {
	ID       int64  `json:"id" binding:"required"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password" binding:"required"`
}
