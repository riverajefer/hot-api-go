package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/riverajer/hot-bread-api/initializers"
	"github.com/riverajer/hot-bread-api/routes"
	"github.com/riverajer/hot-bread-api/routes/products"
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

	routes.AuthRoutes(r)
	routes.UserRoutes(r)
	products.CategoryRoutes(r)
	products.ProductsRoutes(r)
	r.Run()

	log.Print("starting server!...")
}
