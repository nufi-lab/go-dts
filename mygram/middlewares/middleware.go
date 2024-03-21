package middlewares

import (
	"errors"
	"mygram/config"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

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

// authMiddleware is middleware to authenticate user using JWT
// func AuthMiddleware(c *gin.Context) (*config.JWTClaim, error) {
// 	// Get token from Authorization header
// 	tokenString := c.GetHeader("Authorization")
// 	if tokenString == "" {
// 		return nil, jwt.ErrTokenRequiredClaimMissing
// 	}

// 	// Extract JWT claims from token
// 	claims := &config.JWTClaim{}
// 	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
// 		return config.JWT_KEY, nil
// 	})
// 	if err != nil || !token.Valid {
// 		return nil, jwt.ErrTokenInvalidClaims
// 	}

// 	return claims, nil
// }
