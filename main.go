package main

import (
	"apps/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	utils.InitDB()
	// Create a new Gin router instance
	router := gin.Default()

	// Define a route for the root URL
	router.GET("/", func(ctx *gin.Context) {
		// Respond with a JSON message
		ctx.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	// Start the server on port 8080
	router.Run(":8080")
}
