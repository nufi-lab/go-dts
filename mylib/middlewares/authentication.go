package middlewares

import (
	"fmt"
	"mylib/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

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
func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, err := helpers.VerifyToken(c)

		fmt.Printf(claims.GetSubject())

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
