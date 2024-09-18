package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/riverajer/hot-bread-api/controllers"
	"github.com/riverajer/hot-bread-api/middleware"
)

func AuthRoutes(r *gin.Engine) {
	// auth routes
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)
}
