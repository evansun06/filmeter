package services

import (
	"restful-movie-api/models"
	"restful-movie-api/repositories"
)

// Acts as a service layer for buiness logic
// Personal Note: May not be used much right now, but is an important part of api architecture for scalable services.
type Services struct {
	Repo *repositories.Repository
}

// EFFECT: Utilizes the UserRepository to aquire data.
func (s *Services) GetAllUsers() ([]*models.User, error) {
	return s.Repo.GetAllUsers()
}

// EFFECT: Add new user to the database
func (s *Services) AddUsers(newUser *models.User) error {
	return s.Repo.UploadNewUser(newUser)
}
