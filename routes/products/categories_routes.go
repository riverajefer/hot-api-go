package products

import (
	"github.com/gin-gonic/gin"
	"github.com/riverajer/hot-bread-api/controllers/products" // Cambiar la ruta de importación
)

func CategoryRoutes(r *gin.Engine) {
	r.POST("/categories", products.CreateCategory)
	r.GET("/categories", products.GetCategories)
	r.GET("/categories/:id", products.GetCategory)
	r.PUT("/categories/:id", products.UpdateCategory)
	r.DELETE("/categories/:id", products.DeleteCategory)
}
