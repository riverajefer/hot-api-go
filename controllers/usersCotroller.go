package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/riverajer/hot-bread-api/initializers"
	"github.com/riverajer/hot-bread-api/models"
	"github.com/riverajer/hot-bread-api/utils"
	"golang.org/x/crypto/bcrypt"
)

// Estructuras de request para Signup y Login
type LoginBody struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type SignupBody struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	Name     string `json:"name" binding:"required"`
}

// Signup
func Signup(c *gin.Context) {
	var body SignupBody
	if err := c.ShouldBindJSON(&body); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid request body")
		return
	}

	hash, err := utils.HashPassword(body.Password)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to hash password")
		return
	}

	user := models.User{
		Name:     body.Name,
		Email:    body.Email,
		Password: hash,
	}

	if result := initializers.DB.Create(&user); result.Error != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to create user")
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}

// Login
func Login(c *gin.Context) {
	var body LoginBody
	if err := c.ShouldBindJSON(&body); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid request body")
		return
	}

	var user models.User
	if err := initializers.DB.First(&user, "email = ?", body.Email).Error; err != nil {
		utils.RespondWithError(c, http.StatusUnauthorized, "Invalid email or password")
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
		utils.RespondWithError(c, http.StatusUnauthorized, "Invalid email or password")
		return
	}

	tokenString, err := utils.GenerateJWT(user.ID)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to create token")
		return
	}

	// Guardar el token en una cookie
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})
}

// Validate
func Validate(c *gin.Context) {
	user, _ := c.Get("user")
	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}
