// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

type Permission struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}

type User struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}
