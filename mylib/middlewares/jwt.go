package middlewares

import (
	"errors"
	"fmt"
	"mylib/config"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthorizeRoleMiddleware(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, exists := c.Get("claims")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "JWT claims not found"})
			c.Abort()
			return
		}

		claimsData, ok := claims.(*config.JWTClaim)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid JWT claims"})
			c.Abort()
			return
		}

		// fmt.Println("Extracted Role:", claimsData.Role)     // Debugging
		// fmt.Println("Extracted User:", claimsData.Username) // Debugging

		authorized := false
		for _, role := range roles {
			if claimsData.Role == role {
				authorized = true
				break
			}
		}

		if !authorized {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized, role not allowed"})
			c.Abort()
			return
		}

		c.Next()
	}
}

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

		c.Header("Authorization", "Bearer "+tokenString)

		// fmt.Println(token)

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

		c.Set("claims", claims)

		c.Next()
	}
}

func VerifyToken(c *gin.Context) (uint, error) {
	claims, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "JWT claims not found"})
		c.Abort()
	}

	claimsData, ok := claims.(*config.JWTClaim)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid JWT claims"})
		c.Abort()
	}

	// fmt.Println("Extracted Role:", claimsData.Role)     // Debugging
	fmt.Println("Extracted User ID:", claimsData.ID) // Debugging

	userID := claimsData.ID

	return userID, nil
}
