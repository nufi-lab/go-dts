package config

import "github.com/golang-jwt/jwt/v5"

var JWT_KEY = []byte("hfsubwilakewuajhuehiqkah")

type JWTClaim struct {
	UserID   uint
	Username string
	jwt.RegisteredClaims
}
