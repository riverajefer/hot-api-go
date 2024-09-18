package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/riverajer/hot-bread-api/initializers"
	"github.com/riverajer/hot-bread-api/models"
	"github.com/riverajer/hot-bread-api/utils"
)

// Crud user controller
func CreateUser(c *gin.Context) {
	var body struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
		Location string `json:"location"`
		Phone    string `json:"phone"`
		UserType string `json:"user_type"` // "Customer" o "Merchant"
	}

	// Bind JSON request body to struct
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Hash password
	hashedPassword, err := utils.HashPassword(body.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error hashing password"})
		return
	}

	// Create the user
	user := models.User{
		Name:     body.Name,
		Email:    body.Email,
		Password: hashedPassword,
		Location: body.Location,
		Phone:    body.Phone,
		UserType: models.UserType(body.UserType),
	}

	// Save to database
	if result := initializers.DB.Create(&user); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User created successfully", "user": user})
}

func GetUsers(c *gin.Context) {
	var users []models.User

	// Fetch all users from the database
	if result := initializers.DB.Find(&users); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"users": users})
}

func GetUser(c *gin.Context) {
	id := c.Param("id")

	var user models.User
	// Find user by ID
	if result := initializers.DB.First(&user, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")

	var body struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Location string `json:"location"`
		Phone    string `json:"phone"`
		UserType string `json:"user_type"`
	}

	// Bind JSON request body to struct
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	var user models.User
	// Find user by ID
	if result := initializers.DB.First(&user, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Update fields
	user.Name = body.Name
	user.Email = body.Email
	user.Location = body.Location
	user.Phone = body.Phone
	user.UserType = models.UserType(body.UserType)

	// Save to database
	if result := initializers.DB.Save(&user); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully", "user": user})
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	var user models.User
	// Find user by ID
	if result := initializers.DB.First(&user, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Delete the user
	if result := initializers.DB.Delete(&user); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
