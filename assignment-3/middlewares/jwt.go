package middlewares

import (
	"assignment-3/config"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// Middleware untuk mengotorisasi pengguna berdasarkan peran (role)
// Middleware for authorizing users based on role
func AuthorizeRoleMiddleware(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get JWT claim from context
		claims, _ := c.Get("claims")
		claimsData, _ := claims.(*config.JWTClaim)

		// Check if user's role is among the allowed roles
		authorized := false
		for _, role := range roles {
			if claimsData.Role == role {
				authorized = true
				break
			}
		}

		// If not authorized, return Unauthorized
		if !authorized {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized, role not allowed"})
			c.Abort()
			return
		}

		// If authorized, proceed to the next handler
		c.Next()
	}
}

// Middleware JWT yang telah dimodifikasi
func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := c.Cookie("token")
		if err != nil {
			if errors.Is(err, http.ErrNoCookie) {
				c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
				c.Abort()
				return
			}
		}

		claims := &config.JWTClaim{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
			return config.JWT_KEY, nil
		})

		if err != nil {
			if errors.Is(err, jwt.ErrSignatureInvalid) {
				c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
				c.Abort()
				return
			} else if errors.Is(err, jwt.ErrTokenExpired) {
				c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized, Token expired!"})
				c.Abort()
				return
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
				c.Abort()
				return
			}
		}

		if !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			c.Abort()
			return
		}

		// Set klaim JWT di konteks untuk digunakan oleh middleware lainnya
		c.Set("claims", claims)

		c.Next()
	}
}
