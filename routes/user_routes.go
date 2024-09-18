package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/riverajer/hot-bread-api/controllers"
)

func UserRoutes(r *gin.Engine) {
	// user routes
	r.POST("/users", controllers.CreateUser)
	r.GET("/users", controllers.GetUsers)
	r.GET("/users/:id", controllers.GetUser)
	r.PUT("/users/:id", controllers.UpdateUser)
	r.DELETE("/users/:id", controllers.DeleteUser)
}
