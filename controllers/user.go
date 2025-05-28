package controllers

import (
	"log"
	"net/http"
	"restful-movie-api/models"
	"restful-movie-api/services"

	"github.com/gin-gonic/gin"
)

// Direct access functions for *gin.Engine
type UserController struct {
	Service *services.Services
}

// EFFECT: Calls service layer to get all users
func (uc *UserController) GetAllUsers(c *gin.Context) {
	users, err := uc.Service.GetAllUsers()
	if err != nil {
		log.Printf("Could not get Users:%v", err)
		return
	}

	c.IndentedJSON(http.StatusOK, users)
}

// EFFECT: Calls service layer to post new user
func (uc *UserController) UploadNewUser(c *gin.Context) {
	newUser := &models.User{}

	err := c.BindJSON(newUser)

	if err != nil {
		log.Printf("Bad Request: Could not bind JSON to User: %v", err)
		return
	}

	uc.Service.AddUsers(newUser)
}
