package helpers

import (
	"errors"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var JWT_KEY = []byte("jfidjeihoahudhuehkdsfsa")

type JWTClaim struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Role     string
	jwt.RegisteredClaims
}

func GenerateToken(id uint, username string) (string, error) {
	claims := jwt.MapClaims{
		"id":       id,
		"username": username,
		"exp":      time.Now().Add(time.Hour * 1).Unix(), // Token berlaku selama 24 jam
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(JWT_KEY)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

// func VerifyToken(ctx *gin.Context) (jwt.MapClaims, error) {
// 	errResponse := errors.New("Unauthorized")
// 	headerToken := ctx.GetHeader("Authorization")

// 	if headerToken == "" {
// 		return nil, errResponse
// 	}

// 	if !strings.HasPrefix(headerToken, "Bearer ") {
// 		return nil, errResponse
// 	}

// 	tokenString := strings.TrimPrefix(headerToken, "Bearer ")

// 	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
// 		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, errors.New("invalid token")
// 		}
// 		return JWT_KEY, nil
// 	})

// 	if err != nil || !token.Valid {
// 		return nil, errResponse
// 	}

// 	claims, ok := token.Claims.(jwt.MapClaims)
// 	if !ok {
// 		return nil, errResponse
// 	}

// 	return claims, nil
// }

func VerifyToken(ctx *gin.Context) (jwt.MapClaims, error) {
	errResponse := errors.New("Unauthorized")
	headerToken := ctx.GetHeader("Authorization")

	// Cek apakah token ada di dalam header Authorization
	if headerToken != "" {
		if !strings.HasPrefix(headerToken, "Bearer ") {
			return nil, errResponse
		}

		tokenString := strings.TrimPrefix(headerToken, "Bearer ")

		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("invalid token")
			}
			return JWT_KEY, nil
		})

		if err != nil || !token.Valid {
			return nil, errResponse
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return nil, errResponse
		}

		return claims, nil
	}

	// Cek apakah token ada di dalam query string
	queryToken := ctx.Query("token")
	if queryToken != "" {
		token, err := jwt.Parse(queryToken, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("invalid token")
			}
			return []byte("secret"), nil
		})

		if err != nil || !token.Valid {
			return nil, errResponse
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return nil, errResponse
		}

		return claims, nil
	}

	return nil, errResponse
}
