package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/riverajer/hot-bread-api/initializers"
	"github.com/riverajer/hot-bread-api/models"
)

func RequireAuth(c *gin.Context) {
	fmt.Println("I'm middleware 2")
	// 1. Get the cookie

	tokenString, err := c.Cookie("Authorization")
	fmt.Println("tokenString", tokenString)

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	// 2. Decode and validate it

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		// check the exp
		fmt.Println("user id: ", claims["sub"])
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			fmt.Println("Unautorized has expired")
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		var user models.User
		initializers.DB.First(&user, claims["sub"])

		if user.ID == 0 {
			fmt.Println("Unautorized user not found")
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		// attach to req
		c.Set("user", user)

		c.Next()

	} else {
		fmt.Println("Unautorized")
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

}
