package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/riverajer/hot-bread-api/controllers"
	"github.com/riverajer/hot-bread-api/initializers"
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

	r.POST("/signup", controllers.Signup)

	r.Run()

	log.Print("starting server!...")
}
