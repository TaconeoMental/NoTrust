package models

import "github.com/golang-jwt/jwt/v5"

type NoTrustClaims struct {
	UUID string `json:"uuid"`
	jwt.RegisteredClaims
}
