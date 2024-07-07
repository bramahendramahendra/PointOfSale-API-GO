package main

import (
	"apps/controllers"
	"apps/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	utils.InitDB()
	// Create a new Gin router instance
	r := gin.Default()

	r.GET("/users", controllers.GetUsers)
	r.POST("/users", controllers.CreateUser)
	r.GET("/users/:id", controllers.GetUser)
	r.PUT("/users/:id", controllers.UpdateUser)
	r.DELETE("/users/Temporary/:id", controllers.DeleteUserTemporary)
	r.DELETE("/users/Permanently/:id", controllers.DeleteUserPermanently)

	// Start the server on port 8080
	r.Run(":8080")
}
