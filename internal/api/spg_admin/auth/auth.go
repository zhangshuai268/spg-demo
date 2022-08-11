package auth

import "github.com/dgrijalva/jwt-go"

type AdminClaim struct {
	UserId int `json:"user_id"`
	jwt.StandardClaims
}


