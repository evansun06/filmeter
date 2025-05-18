package models


// Defualt Client Type
type User struct {
	ID int64
	Username string
	Password string
	Email string
	Reviews []Review
}