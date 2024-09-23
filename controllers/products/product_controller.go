package products

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/riverajer/hot-bread-api/initializers"
	"github.com/riverajer/hot-bread-api/models"
)

// Helper function to send error responses
func respondWithError(c *gin.Context, status int, message string) {
	c.JSON(status, gin.H{"error": message})
}

// Helper function to find a product by ID
func findProductByID(c *gin.Context, id string) (*models.Product, error) {
	var product models.Product
	if result := initializers.DB.First(&product, id); result.Error != nil {
		return nil, result.Error
	}
	return &product, nil
}

// CreateProduct creates a new product
func CreateProduct(c *gin.Context) {
	var body struct {
		Name        string  `json:"name" binding:"required"`
		Description string  `json:"description"`
		Price       float64 `json:"price" binding:"required"`
		CategoryID  uint    `json:"category_id" binding:"required"`
	}

	// Bind JSON request body to struct
	if err := c.ShouldBindJSON(&body); err != nil {
		respondWithError(c, http.StatusBadRequest, "Invalid request: "+err.Error())
		return
	}

	// Create and save the product
	product := models.Product{
		Name:        body.Name,
		Description: body.Description,
		Price:       body.Price,
		CategoryID:  body.CategoryID,
	}
	if result := initializers.DB.Create(&product); result.Error != nil {
		respondWithError(c, http.StatusInternalServerError, result.Error.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product created successfully", "product": product})
}

// GetProducts retrieves all products
func GetProducts(c *gin.Context) {
	var products []models.Product
	if result := initializers.DB.Find(&products); result.Error != nil {
		respondWithError(c, http.StatusInternalServerError, result.Error.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"products": products})
}

// GetProduct retrieves a product by ID
func GetProduct(c *gin.Context) {
	product, err := findProductByID(c, c.Param("id"))
	if err != nil {
		respondWithError(c, http.StatusNotFound, "Product not found")
		return
	}
	c.JSON(http.StatusOK, gin.H{"product": product})
}

// UpdateProduct updates a product by ID
func UpdateProduct(c *gin.Context) {
	var body struct {
		Name        string  `json:"name"`
		Description string  `json:"description"`
		Price       float64 `json:"price"`
		CategoryID  uint    `json:"category_id"`
	}

	// Bind JSON request body to struct
	if err := c.ShouldBindJSON(&body); err != nil {
		respondWithError(c, http.StatusBadRequest, "Invalid request: "+err.Error())
		return
	}

	// Find the product by ID
	product, err := findProductByID(c, c.Param("id"))
	if err != nil {
		respondWithError(c, http.StatusNotFound, "Product not found")
		return
	}

	// Update product fields
	product.Name = body.Name
	product.Description = body.Description
	product.Price = body.Price
	product.CategoryID = body.CategoryID

	// Save the changes
	if result := initializers.DB.Save(&product); result.Error != nil {
		respondWithError(c, http.StatusInternalServerError, result.Error.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product updated successfully", "product": product})
}

// DeleteProduct deletes a product by ID
func DeleteProduct(c *gin.Context) {
	product, err := findProductByID(c, c.Param("id"))
	if err != nil {
		respondWithError(c, http.StatusNotFound, "Product not found")
		return
	}

	// Delete the product
	if result := initializers.DB.Delete(&product); result.Error != nil {
		respondWithError(c, http.StatusInternalServerError, result.Error.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}
