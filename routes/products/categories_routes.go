package products

import (
	"github.com/gin-gonic/gin"
	"github.com/riverajer/hot-bread-api/controllers/products" // Cambiar la ruta de importación
)

func CategoryRoutes(r *gin.Engine) {
	// Create category
	r.POST("/categories", products.CreateCategory)

	// Get all categories
	r.GET("/categories", products.GetCategories)

	// Get category by ID
	r.GET("/categories/:id", products.GetCategory)

	// Update category
	r.PUT("/categories/:id", products.UpdateCategory)

	// Delete category
	r.DELETE("/categories/:id", products.DeleteCategory)
}
