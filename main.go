package main

import (
	"log"
	"restful-movie-api/controllers"
	"restful-movie-api/database"
	"restful-movie-api/repositories"
	"restful-movie-api/services"

	_ "restful-movie-api/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // Swagger UI files
	ginSwagger "github.com/swaggo/gin-swagger" // Gin middleware for Swagger
)

// @title Filmeter RESTful API
// @version 0.0
// @description A lightweight API for managing movie social media data
// @host localhost:8080
// @BasePath 
func main() {
	// connect to db
	db, err := database.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	// ** defer is kw for run after the enclosing function is complete
	defer db.Close()

	// Setup the Gin server
	server := gin.Default()

	// Initialize all layers
	repo := repositories.Repository{DB: db}
	sc := services.Services{Repo: &repo}
	uc := controllers.UserController{Service: &sc}

	// Initialize Routes
	server.GET("/users", uc.GetAllUsers)
	server.POST("/user", uc.UploadNewUser)

	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	server.Run(":8080")
}
