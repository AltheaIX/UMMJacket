package model

import "github.com/golang-jwt/jwt/v5"

type Claims struct {
	Nim string `json:"nim"`
	jwt.RegisteredClaims
}
