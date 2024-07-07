package main

import (
	"apps/controllers"
	"apps/utils"

	_ "apps/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1
func main() {
	utils.InitDB()
	// Create a new Gin router instance
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/users", controllers.GetUsers)
	r.POST("/users", controllers.CreateUser)
	r.GET("/users/:id", controllers.GetUser)
	r.PUT("/users/:id", controllers.UpdateUser)
	r.DELETE("/users/Temporary/:id", controllers.DeleteUserTemporary)
	r.DELETE("/users/Permanently/:id", controllers.DeleteUserPermanently)

	// Start the server on port 8080
	r.Run(":8080")
}
