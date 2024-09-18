package utils

import (
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// RespondWithError envia una respuesta de error JSON
func RespondWithError(c *gin.Context, status int, message string) {
	c.JSON(status, gin.H{"error": message})
}

// HashPassword genera el hash de una contraseña
func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(hash), err
}

// GenerateJWT genera un token JWT
func GenerateJWT(userID uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(), // Expira en 30 días
	})
	return token.SignedString([]byte(os.Getenv("SECRET")))
}
