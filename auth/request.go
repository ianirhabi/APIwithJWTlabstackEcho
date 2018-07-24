package auth

import jwt "github.com/dgrijalva/jwt-go"

type JwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.StandardClaims
}
