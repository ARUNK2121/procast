package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var InvalidAuthorizationToken = "Invalid Authorization Token"
var MissingAuthorizationToken = "Missing Authorization Token"
var UnAuthorizedAccess = "Unauthorized Access"
var TokenProblem = "Some Problem in Token"

func AdminAuthMiddleware(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": MissingAuthorizationToken})
		c.Abort()
		return
	}
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		return []byte("comebuyjersey"), nil
	})

	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": InvalidAuthorizationToken})
		c.Abort()
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": InvalidAuthorizationToken})
		c.Abort()
		return
	}

	role, ok := claims["role"].(string)
	if !ok || role != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": UnAuthorizedAccess})
		c.Abort()
		return
	}

	id, ok := claims["id"].(int)
	if !ok || id == 0 {
		c.JSON(http.StatusForbidden, gin.H{"error": TokenProblem})
		c.Abort()
		return
	}

	c.Set("id", id)

	c.Next()
}

func SuperAdminAuthMiddleware(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": MissingAuthorizationToken})
		c.Abort()
		return
	}
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		return []byte("comebuyjersey"), nil
	})

	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": InvalidAuthorizationToken})
		c.Abort()
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": InvalidAuthorizationToken})
		c.Abort()
		return
	}

	role, ok := claims["role"].(string)
	if !ok || role != "super_admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": UnAuthorizedAccess})
		c.Abort()
		return
	}

	id, ok := claims["id"].(int)
	if !ok || id == 0 {
		c.JSON(http.StatusForbidden, gin.H{"error": TokenProblem})
		c.Abort()
		return
	}

	c.Set("id", id)

	c.Next()
}
