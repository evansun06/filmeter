package controllers

import (
	"database/sql"
	"log"
	"net/http"
	"restful-movie-api/models"
	"restful-movie-api/services"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

// Direct access functions for *gin.Engine
type UserController struct {
	Service *services.UserServices
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

// @Summary Get single user by ID
// @Description Retrieves a single user by a unique ID parameter
// @Tags Users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} models.User "Single User"
// @Failure 400 {object} models.ErrorResponse "Invalid ID input (non integer)"
// @Failure 404 {object} models.ErrorResponse "No user found"
// @Failure 500 {object} models.ErrorResponse "Failed to retrieve user"
// @Router /user/{id} [get]
func (uc *UserController) GetUserByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)

	if err != nil {
		log.Printf("Invalid ID parameter %s: %v", idParam, err)
		c.IndentedJSON(http.StatusBadRequest, models.ErrorResponse{Error: "Invalid ID input"})
		return
	}

	user, err := uc.Service.GetUserByID(id)

	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("No user found with ID %d", id)
			c.IndentedJSON(http.StatusNotFound, models.ErrorResponse{Error: "No user found"})
			return
		}

		log.Printf("Error in searching for user with ID %d: %v", id, err)
		c.IndentedJSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Error searching for user"})
		return
	}

	c.IndentedJSON(http.StatusOK, user)

}

// @Summary Get single user by username
// @Description Retrieves a single user by a username parameter
// @Tags Users
// @Accept json
// @Produce json
// @Param username path string true "Username"
// @Success 200 {object} models.User "Single User"
// @Failure 400 {object} models.ErrorResponse "Invalid username input"
// @Failure 404 {object} models.ErrorResponse "User with given username does not exist"
// @Failure 500 {object} models.ErrorResponse "Failed to retrieve user"
// @Router /user/username/{username} [get]
func (uc *UserController) GetUserByUsername(c *gin.Context) {
	username := c.Param("username")

	if username == "" {
		log.Printf("Invalid username parameter: %s", username)
		c.IndentedJSON(http.StatusBadRequest, models.ErrorResponse{Error: "Invalid username input"})
		return
	}

	user, err := uc.Service.Repo.GetUserByUsername(username)

	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("No user found with %s", username)
			c.IndentedJSON(http.StatusNotFound, models.ErrorResponse{Error: "User with given username does not exist"})
			return
		}

		log.Printf("Error in searching for user with ID %s: %v", username, err)
		c.IndentedJSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Error searching for user"})
		return
	}

	c.IndentedJSON(http.StatusOK, user)
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
		c.IndentedJSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Failed to add user"})
		return
	}

	// Success
	c.JSON(http.StatusCreated, models.SuccessResponse{Message: "User added successfully"})
}

// @Summary Delete user by ID
// @Description Delete a single user by a unique ID parameter
// @Tags Users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 202 {object} models.SuccessResponse "Successfully removed user"
// @Failure 400 {object} models.ErrorResponse "Invalid ID input (non integer)"
// @Failure 404 {object} models.ErrorResponse "No user with given ID exists"
// @Failure 500 {object} models.ErrorResponse "Failed to delete user"
// @Router /user/{id} [delete]
func (uc *UserController) DeleteUserByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)

	if err != nil {
		log.Printf("Invalid ID parameter %s: %v", idParam, err)
		c.IndentedJSON(http.StatusBadRequest, models.ErrorResponse{Error: "Invalid ID input"})
		return
	}

	err = uc.Service.Repo.DeleteUserById(id)

	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("No user found with ID %d", id)
			c.IndentedJSON(http.StatusNotFound, models.ErrorResponse{Error: "No user with given ID exists"})
			return
		}

		log.Printf("Error in searching for user with ID %d: %v", id, err)
		c.IndentedJSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Error deleting user"})
		return

	}

	// Success - No Content
	c.IndentedJSON(http.StatusAccepted, models.SuccessResponse{Message: "Successfully removed user"})
}