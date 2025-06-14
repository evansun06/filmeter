package services

import (
	"restful-movie-api/models"
	"restful-movie-api/repositories"
)

// Acts as a service layer for buiness logic
// Personal Note: May not be used much right now, but is an important part of api architecture for scalable services.
type UserServices struct {
	Repo *repositories.UserRepository
}

// EFFECT: Utilizes the UserRepository to aquire data.
func (s *UserServices) GetAllUsers() ([]*models.User, error) {
	return s.Repo.GetAllUsers()
}

// EFFECT: Add new user to the database
func (s *UserServices) AddUsers(newUser *models.User) error {
	return s.Repo.UploadNewUser(newUser)
}

// EFFECT: Retrieve user based on ID
func (s *UserServices) GetUserByID(id int64) (*models.User, error) {
	return s.Repo.GetUserById(id)
}

// EFFECT: Retrieve user based on Username
func (s *UserServices) GetserByUsername(username string) (*models.User, error) {
	return s.Repo.GetUserByUsername(username)
}
