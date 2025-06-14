package repositories

import (
	"database/sql"
	"log"
	"restful-movie-api/models"

	"github.com/lib/pq"
)

// Allow database dependency and isolate user data logic
type UserRepository struct {
	DB *sql.DB
}

// EFFECT: Get all users in the associated database.
func (repo *UserRepository) GetAllUsers() ([]*models.User, error) {
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


// EFFECT: Retrieves user object by ID search
func (repo *UserRepository) GetUserById(id int64) (*models.User, error ) {
	stmt := "SELECT id, username, email, hashed_password FROM users WHERE id = $1"

	user := &models.User{}

	queryErr := repo.DB.QueryRow(stmt, id).Scan(&user.ID, &user.Username, &user.Email, &user.HashedPassword)

	if queryErr != nil {
		if queryErr == sql.ErrNoRows {
            // No user found
            log.Printf("No user found with ID: %d", id)
            return nil, nil
        }

        // Other errors
        log.Printf("Error querying for user by ID %d: %v", id, queryErr)
        return nil, queryErr
	}

	return user, nil
}

// EFFECT: Retrieve User object using username
func (repo *UserRepository) GetUserByUsername(username string) (*models.User, error) {
	stmt := "SELECT id, username, email, hashed_password FROM users WHERE username = $1"

	user := &models.User{}

	queryErr := repo.DB.QueryRow(stmt, username).Scan(&user.ID, &user.Username, &user.Email, &user.HashedPassword)

	if queryErr != nil {
		if queryErr == sql.ErrNoRows {
            // No user found
            log.Printf("No user found with username: %s", username)
            return nil, nil
        }

        // Other errors
        log.Printf("Error querying for user by username %s: %v", username, queryErr)
        return nil, queryErr
	}

	return user, nil
}

// EFFECT: Upload a new user to the associated database.
// REQUIRES: Assume the new user is instantiated
func (repo *UserRepository) UploadNewUser(user *models.User) error {
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

// EFFECT: Deletes user in database by ID
func (repo *UserRepository) DeleteUserById(id int64) error {
	stmt := "DELETE FROM users WHERE id = $1"

	_, err := repo.DB.Exec(stmt, id)

	if err != nil {
		log.Printf("Error deleting review with ID %d: %v", id, err)
		return err
	}
	return nil
}