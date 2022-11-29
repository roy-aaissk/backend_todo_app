package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Print("hello")

	// routing set
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
				"message": "pong",
		})
	})
	router.Run("127.0.0.1:8080")
}
