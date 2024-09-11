package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/riverajer/hot-bread-api/initializers"
	"github.com/riverajer/hot-bread-api/models"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {
	// Get the email/pass of req body
	var body struct {
		Email    string
		Password string
		Name     string
	}
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	// Has the password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})
		return
	}

	// Create the User
	user := models.User{
		Name:     body.Name,
		Email:    body.Email,
		Password: string(hash),
	}

	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user: " + result.Error.Error(),
		})
		return
	}

	// Respond
	c.JSON(http.StatusOK, gin.H{})

}
