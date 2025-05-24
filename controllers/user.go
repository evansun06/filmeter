package controllers

import (
	"database/sql"
	"net/http"
	"restful-movie-api/database"
	"restful-movie-api/errors"
	"restful-movie-api/models"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

// EFFECT: Endpoint that returns all users in indented json
func GetUsers(c *gin.Context) {
	// Send Query
	rows, queryError := database.Db.Query("SELECT * FROM users;")

	if queryError != nil {
		errors.InternalServerError(c, "Error in query: \"users\"")
		return
	}

	defer rows.Close() // Prevents Data Leakage

	// Convert SQL -> Go
	convertedSlice, conversionErr := rowsToUsers(rows)

	if conversionErr != nil {
		errors.InternalServerError(c, "ConversionError: \"controllers/users.go\"")
	}

	c.IndentedJSON(http.StatusOK, convertedSlice)
}

// EFFECT: Converts the queried rows into a a Go Slice
//   - Returns err if wrong sql.Rows is given.
//
// REQUIRES: assumes that a slice of Users are passed in
func rowsToUsers(rows *sql.Rows) ([]*models.User, error) {
	// slice of pointers
	var users []*models.User

	for rows.Next() {
		// pointer to empty User struct with fields set to zero values.
		user := &models.User{}
		err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.HashedPassword)

		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

// EFFECT: Uploads new User
//         If duplicate username/email, will handle by returning conflict error
func HandleUserUpload(c *gin.Context) {
	newUser := &models.User{}

	err := c.BindJSON(newUser)

	if err != nil {
		errors.BadRequestError(c, "Bad Request Made")
		return
	}

	stmt := "INSERT INTO users (username, email, hashed_password) VALUES ($1, $2, $3)"

	_, insertError := database.Db.Exec(stmt, newUser.Username, newUser.Email, newUser.HashedPassword)

	if insertError != nil {
		if pqErr, ok := insertError.(*pq.Error); ok && pqErr.Code == "23505" {
			errors.ConflictError(c, "duplicate username/email")
			return
		}

		errors.InternalServerError(c, "Error when inserting new username")
	}

	c.IndentedJSON(http.StatusCreated, newUser)
}

