package routes

import (
	"restful-movie-api/controllers"

	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine) {
	router.GET("/users", controllers.GetUsers)
}
