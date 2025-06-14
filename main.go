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
	gin.SetMode(gin.ReleaseMode)
	server.SetTrustedProxies([]string{"127.0.0.1"})
	

	// Initialize all layers
	repo := repositories.UserRepository{DB: db}
	sc := services.UserServices{Repo: &repo}
	uc := controllers.UserController{Service: &sc}

	// Initialize Routes
	server.GET("/users", uc.GetAllUsers)
	server.GET("/user/:id", uc.GetUserByID)
	server.GET("/user/username/:username", uc.GetUserByUsername)

	server.POST("/user", uc.UploadNewUser)
	server.DELETE("/user/:id", uc.DeleteUserByID)
	

	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	
	if err := server.Run(":8080"); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}
