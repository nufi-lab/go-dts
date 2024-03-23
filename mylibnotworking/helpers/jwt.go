package helpers

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("jfidjeihoahudhuehkdsfsa")

type Claims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

func GenerateToken(username string) (string, error) {
	// Membuat claims JWT dengan menyertakan username pengguna.
	claims := &Claims{
		Username:         username,
		RegisteredClaims: jwt.RegisteredClaims{},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Menandatangani token dengan secret key.
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
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
			return jwtKey, nil
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
