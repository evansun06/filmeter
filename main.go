package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	fmt.Printf("server: %v\n", server)
}