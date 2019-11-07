package models

import jwt "github.com/dgrijalva/jwt-go"

type Claim struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
