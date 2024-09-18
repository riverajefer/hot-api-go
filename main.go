package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/riverajer/hot-bread-api/controllers"
	"github.com/riverajer/hot-bread-api/initializers"
	"github.com/riverajer/hot-bread-api/middleware"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// auth routes
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)

	// user routes
	r.POST("/users", controllers.CreateUser)
	r.GET("/users", controllers.GetUsers)
	r.GET("/users/:id", controllers.GetUser)
	r.PUT("/users/:id", controllers.UpdateUser)
	r.DELETE("/users/:id", controllers.DeleteUser)

	r.Run()

	log.Print("starting server!...")
}
