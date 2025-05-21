package models

// Defualt Client Type
type User struct {
	ID             int64  `json:"id"`
	Username       string `json:"username"`
	HashedPassword string `json:"-"`
	Email          string `json:"email"`
}
