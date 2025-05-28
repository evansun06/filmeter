package main

import (
	"log"
	"restful-movie-api/controllers"
	"restful-movie-api/database"
	"restful-movie-api/repositories"
	"restful-movie-api/services"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

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

	server.Run(":8080")
}
