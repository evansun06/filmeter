package controllers

import (
	"database/sql"
	"log"
	"net/http"
	"restful-movie-api/database"
	"restful-movie-api/models"
	
	"github.com/gin-gonic/gin"
)

// EFFECT: Endpoint that returns all users in indented json
func GetUsers(c *gin.Context) {
	// Send Query
	rows, queryError := database.Db.Query("SELECT * FROM users;")
	defer rows.Close() // Prevents Data Leakage

	if queryError != nil {
		log.Println("Error in query: \"users\"")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database query failed"})
		return
	}

	// Convert SQL -> Go
	convertedSlice, conversionErr := rowsToUsers(rows)

	if conversionErr != nil {
		log.Println("ConversionError: \"controllers/users.go\"")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to convert database rows"})
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

// EFFECT: Upload new User
func UploadUser(c *gin.Context) {

}
