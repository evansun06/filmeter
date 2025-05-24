package errors

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// EFFECT: Boilerplate code for logging a query error.
func InternalServerError(c *gin.Context, message string) {
	log.Println(message)
	c.JSON(http.StatusInternalServerError, gin.H{"Query Error": message})
}

func BadRequestError(c *gin.Context, message string) {
	log.Println(message)
	c.JSON(http.StatusBadRequest, gin.H{"Bad Request Error": message})
}

func ConflictError(c *gin.Context, message string) {
	log.Println(message)
	c.JSON(http.StatusBadRequest, gin.H{"Conflic Error": message})
}
