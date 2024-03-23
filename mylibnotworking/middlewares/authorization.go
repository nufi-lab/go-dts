package middlewares

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AddAuthorizationHeader() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Ambil token dari header Authorization
		token := c.GetHeader("Authorization")
		if token != "" {
			// Tambahkan header Authorization ke setiap permintaan
			c.Request.Header.Set("Authorization", token)
		}
		c.Next()
	}
}

func Authorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get user data from the context
		userData, exists := c.Get("userData")
		if !exists {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthenticated",
				"message": "User data not found in context",
			})
			return
		}

		// Try to get user ID from the user data
		userIDFloat, ok := userData.(jwt.MapClaims)["id"].(float64)
		if !ok {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error":   "Internal Server Error",
				"message": "Failed to cast user ID from token to float64",
			})
			return
		}

		userID := int(userIDFloat)

		// Get the user ID from the request parameters
		paramUserIDStr := c.Param("id")
		paramUserID, err := strconv.Atoi(paramUserIDStr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": "Invalid user ID in request",
			})
			return
		}

		// Compare user ID from the token with the user ID from the request parameters
		if userID != paramUserID {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error":   "Forbidden",
				"message": "You are not authorized to perform this action",
			})
			return
		}

		// Continue execution if authorization is successful
		c.Next()
	}
}
