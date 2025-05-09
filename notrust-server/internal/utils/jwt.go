package utils

import (
	"errors"
	"time"
	"notrust-server/internal/models"
	"notrust-server/internal/config"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var jwtSecret = []byte(config.GetEnv("JWT_SECRET", "DefaultJWTKey"))

func GenerateJWT() (string, error) {
	claims := models.NoTrustClaims{
		UUID: uuid.NewString(),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func ValidateJWT(tokenString string) error {
	token, err := jwt.ParseWithClaims(tokenString, &models.NoTrustClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return err
	}
	// _ == claims
	if _, ok := token.Claims.(*models.NoTrustClaims); ok && token.Valid {
		return nil
	}
	return errors.New("token inválido")
}

