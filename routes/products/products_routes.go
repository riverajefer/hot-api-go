package products

import (
	"github.com/gin-gonic/gin"
	"github.com/riverajer/hot-bread-api/controllers/products"
)

func ProductsRoutes(r *gin.Engine) {
	r.POST("/products", products.CreateProduct)
	r.GET("/products", products.GetProducts)
	r.GET("/products/:id", products.GetProduct)
	r.PUT("/products/:id", products.UpdateProduct)
	r.DELETE("/products/:id", products.DeleteProduct)
}
