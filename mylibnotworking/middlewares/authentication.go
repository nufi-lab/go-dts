package middlewares

import (
	"mylib/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

// var jwtKey = []byte("jfidjeihoahudhuehkdsfsa")

// func Authentication() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		verifyToken, err := helpers.VerifyToken(c)

// 		if err != nil {
// 			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
// 				"error":   "Unauthenticated",
// 				"message": err.Error(),
// 			})
// 			return
// 		}
// 		c.Set("userData", verifyToken)

//			c.Next()
//		}
//	}
//
// Middleware untuk melakukan autentikasi dengan token JWT.
// func Authenticate(c *gin.Context) {
// 	// Mendapatkan token dari header Authorization.
// 	authHeader := c.GetHeader("Authorization")
// 	if len(authHeader) <= len("Bearer ") {
// 		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
// 		return
// 	}
// 	tokenString := authHeader[len("Bearer "):]

// 	// Parsing token JWT.
// 	claims := &helpers.Claims{}
// 	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
// 		return jwtKey, nil
// 	})
// 	if err != nil || !token.Valid {
// 		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
// 		return
// 	}

// 	// Memasukkan username ke dalam context untuk digunakan oleh handler selanjutnya jika diperlukan.
// 	c.Set("username", claims.Username)
// 	c.Next()
// }

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, err := helpers.VerifyToken(c)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthenticated",
				"message": "Unauthorized",
			})
			return
		}

		// Set token claims into the context
		c.Set("userData", claims)

		c.Next()
	}
}
