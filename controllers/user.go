package controllers

import (
	"database/sql"
	_ "database/sql"
	"log"
	"net/http"
	"restful-movie-api/database"
	"restful-movie-api/models"

	"github.com/gin-gonic/gin"
)

// EFFECT: Endpoint that returns all users in indented json
func GetUsers(c *gin.Context) {
	rows, queryError := database.Db.Query("SELECT * FROM users;")
	if queryError != nil {
		log.Println("Error in query: \"users\"")
	}

	convertedSlice, conversionErr := rowsToUsers(rows)
	if conversionErr != nil {
		log.Println("ConversionError: \"controllers/users.go\"")
	}

	c.IndentedJSON(http.StatusOK, convertedSlice)
	
}

// EFFECT: Converts the queried rows into a a Go Slice
//		- Returns err if wrong sql.Rows is given.
// REQUIRES: assumes that a slice of Users are passed in
//
func rowsToUsers(rows *sql.Rows) ([]*models.User, error){
	//var convertedSlice []*models.User
	for rows.Next() {
		// TODO: figure out .Scan function 
		//var user models.User = rows.Scan(&user.ID, &user.Username, &user.Email, &user.HashedPassword)
	}

	return nil, nil
}


// EFFECT: Upload new User
func UploadUser(c *gin.Context) {

}