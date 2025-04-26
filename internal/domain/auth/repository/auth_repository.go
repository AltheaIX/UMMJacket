package repository

import (
	"errors"
	"github.com/AltheaIX/UMMJacket/internal/domain/auth/model"
	"github.com/AltheaIX/UMMJacket/shared"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strconv"
	"time"
)

type IAuthRepository interface {
	GenerateToken(userID int, nim string) (string, string, error)
	ValidateJWT(tokenString string) (*model.Claims, error)
}

func (a *AuthRepositoryImpl) GenerateToken(userID int, nim string) (string, string, error) {
	claims := model.Claims{
		Nim: nim,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)), // 1 hour expiry
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "jacketlab-umm",
			Subject:   strconv.Itoa(userID),
		},
	}

	accessClaimToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	accessToken, err := accessClaimToken.SignedString([]byte(a.cfg.JWTSecret))
	if err != nil {
		return "", "", err
	}

	refreshTokenClaims := claims
	refreshTokenClaims.ExpiresAt = jwt.NewNumericDate(time.Now().Add((24 * 7) * time.Hour))

	refreshClaimToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)
	refreshToken, err := refreshClaimToken.SignedString([]byte(a.cfg.JWTSecret))
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (a *AuthRepositoryImpl) ValidateJWT(tokenString string) (*model.Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &model.Claims{}, func(token *jwt.Token) (interface{}, error) {
		// Ensure the signing method is HMAC
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(a.cfg.JWTSecret), nil
	})

	if err != nil {
		err = &shared.AppError{Code: http.StatusUnauthorized, Message: "Invalid token"}
		return nil, err
	}

	if claims, ok := token.Claims.(*model.Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}
