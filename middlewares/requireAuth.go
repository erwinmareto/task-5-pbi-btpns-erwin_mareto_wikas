package middlewares

import (
	"net/http"
	"strings"

	"github.com/erwinmareto/profile-api-go/app"
	"github.com/erwinmareto/profile-api-go/database"
	"github.com/erwinmareto/profile-api-go/helpers"
	"github.com/gin-gonic/gin"
)

func CheckAuthentication(c *gin.Context) {
	if _, ok := c.Get("user"); !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthenticated"})
		c.Abort()
		return
	}
}

func RequireAuth(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")

	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "No Authorization header provided"})
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	accessToken := strings.Split(tokenString, "Bearer ")[1]

	tokenClaims, err := helpers.ParseToken(accessToken)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		c.AbortWithStatus(http.StatusUnauthorized)
		c.Abort()
		return
	}

	var user app.User

	if err := database.DB.Where("email=?", tokenClaims.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User Unauthenticated"})
		c.AbortWithStatus(http.StatusUnauthorized)
		c.Abort()
		return
	}

	c.Set("user", user)
	c.Next()

}
