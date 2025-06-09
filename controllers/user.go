package controllers

import (
	"log"
	"net/http"
	"restful-movie-api/models"
	"restful-movie-api/services"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

// Direct access functions for *gin.Engine
type UserController struct {
	Service *services.Services
}

// @Summary Get all users
// @Description Retrieves a list of all users.
// @Tags Users
// @Accept json
// @Produce json 
// @Success 200 {array} models.User "List of all users"
// @Failure 500 {object} models.ErrorResponse "Failed to retrieve users"
// @Router /users [get]
func (uc *UserController) GetAllUsers(c *gin.Context) {
	users, err := uc.Service.GetAllUsers()
	if err != nil {
		log.Printf("Could not get Users:%v", err)
		c.IndentedJSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Failed to retrieve users"})
		return
	}

	c.IndentedJSON(http.StatusOK, users)
}


// @Summary Upload a single new user
// @Description Adds a new user to the database while handling duplicates
// @Tags Users
// @Accept json
// @Produce json
// @Param user body models.User true "User object containing email, id, and username"
// @Success 201 {object} models.SuccessResponse "User added successfully"
// @Failure 400 {object} models.ErrorResponse "Invalid request payload"
// @Failure 409 {object} models.ErrorResponse "User already exists"
// @Failure 500 {object} models.ErrorResponse "Failed to add user"
// @Router /user [post]
func (uc *UserController) UploadNewUser(c *gin.Context) {
	newUser := &models.User{}

	err := c.BindJSON(newUser)

	if err != nil {
		log.Printf("Bad Request: Could not bind JSON to User: %v", err)
		c.IndentedJSON(http.StatusBadRequest, models.ErrorResponse{Error: "Invalid request payload"})
		return
	}

	err = uc.Service.AddUsers(newUser)

	if err != nil {
		// Check Duplicate Error
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
			log.Printf("Duplicate User Error: %v", pqErr)
			c.IndentedJSON(http.StatusConflict, models.ErrorResponse{Error: "User already exists"})
			return
		}

		// Handle other errors
		log.Printf("Could not upload user to database: %v", err)
		c.IndentedJSON(http.StatusInternalServerError, models.ErrorResponse{Error:"Failed to add user"})
		return
	}
	
	// Sucess
	c.JSON(http.StatusCreated, models.SuccessResponse{Message: "User added successfully"})
}
