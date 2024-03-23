package middlewares

import (
	"assignment-3/config"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// Middleware untuk mengotorisasi pengguna berdasarkan peran (role)
// Middleware for authorizing users based on role
// func AuthorizeRoleMiddleware(roles ...string) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		// Get JWT claim from context
// 		claims, _ := c.Get("claims")
// 		claimsData, _ := claims.(*config.JWTClaim)

// 		// Check if user's role is among the allowed roles
// 		authorized := false
// 		for _, role := range roles {
// 			if claimsData.Role == role {
// 				authorized = true
// 				break
// 			}
// 		}

// 		// If not authorized, return Unauthorized
// 		if !authorized {
// 			c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized, role not allowed"})
// 			c.Abort()
// 			return
// 		}

// 		// If authorized, proceed to the next handler
// 		c.Next()
// 	}
// }

func AuthorizeRoleMiddleware(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get JWT claim from context
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

		fmt.Println("Extracted Role:", claimsData.Role)     // Debugging
		fmt.Println("Extracted User:", claimsData.Username) // Debugging

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

		fmt.Println(token)

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

		// if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// 	// Menempatkan informasi pengguna ke dalam konteks gin
		// 	userID, _ := claims["user_id"].(string) // Disesuaikan dengan cara Anda menyimpan ID pengguna di token JWT
		// 	c.Set("user_id", userID)
		// } else {
		// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
		// 	c.Abort()
		// 	return
		// }

		c.Set("claims", claims)

		c.Next()
	}
}
