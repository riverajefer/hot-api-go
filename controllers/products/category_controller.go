package products

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/riverajer/hot-bread-api/initializers"
	"github.com/riverajer/hot-bread-api/models"
)

// Crud category controller

// CreateCategory godoc
func CreateCategory(c *gin.Context) {
	var body struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}

	// Bind JSON request body to struct
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Create the category
	category := models.Category{
		Name:        body.Name,
		Description: body.Description,
	}

	// Save to database
	if result := initializers.DB.Create(&category); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Category created successfully", "category": category})

}

// GetCategories godoc
func GetCategories(c *gin.Context) {
	var categories []models.Category

	// Fetch all categories from the database
	if result := initializers.DB.Find(&categories); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"categories": categories})
}

func GetCategory(c *gin.Context) {
	var category models.Category

	// Fetch category by ID
	if result := initializers.DB.First(&category, c.Param("id")); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"category": category})
}

// UpdateCategory godoc
func UpdateCategory(c *gin.Context) {
	var body struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}

	// Bind JSON request body to struct
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	var category models.Category
	// Find category by ID
	if result := initializers.DB.First(&category, c.Param("id")); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}

	// Update category
	category.Name = body.Name
	category.Description = body.Description

	// Save to database
	if result := initializers.DB.Save(&category); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Category updated successfully", "category": category})
}

// DeleteCategory godoc

func DeleteCategory(c *gin.Context) {
	var category models.Category
	// Find category by ID
	if result := initializers.DB.First(&category, c.Param("id")); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}

	// Delete category
	if result := initializers.DB.Delete(&category); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Category deleted successfully"})
}
