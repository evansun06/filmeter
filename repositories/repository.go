package repositories

import (
	"log"
	"restful-movie-api/models"

	"database/sql"

	"github.com/lib/pq"
)

// Allow database dependency and isolate user data logic
type Repository struct {
	DB *sql.DB
}

// EFFECT: Get all users in the associated database.
func (repo *Repository) GetAllUsers() ([]*models.User, error) {
	// Query
	rows, queryError := repo.DB.Query("SELECT id, username, email, hashed_password FROM users;")

	if queryError != nil {
		return nil, queryError
	}
	defer rows.Close()

	var users []*models.User
	for rows.Next() {
		user := models.User{}
		if err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.HashedPassword); err != nil {
			log.Printf("Error scanning row: %v", err)
			return nil, err
		}
		users = append(users, &user)
	}

	// Check for erros during iteration
	if err := rows.Err(); err != nil {
		log.Printf("Error during rows iteration: %v", err)
		return nil, err
	}

	return users, nil
}

// EFFECT: Upload a new user to the associated database.
// REQUIRES: Assume the new user is instantiated
func (repo *Repository) UploadNewUser(user *models.User) error {
	stmt := "INSERT INTO users (username, email, hashed_password) VALUES ($1, $2, $3)"
	_, insertError := repo.DB.Exec(stmt, user.Username, user.Email, user.HashedPassword)

	if insertError != nil {
		// Check for duplicate error.
		if pqErr, ok := insertError.(*pq.Error); ok && pqErr.Code == "23505" {
			log.Printf("Duplicate User Error: %v", pqErr)
			return pqErr
		}

		log.Printf("Insertion Error: %v", insertError)
		return insertError
	}

	return nil
}
