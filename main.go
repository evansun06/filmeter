package main

import (
	"log"
	"restful-movie-api/database"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	// connect to db
	err := database.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	// ** defer is kw for run after the enclosing function is complete
	defer database.Db.Close()

	// Setup the Gin server
	server := gin.Default()
	server.Run(": 8080")
}
