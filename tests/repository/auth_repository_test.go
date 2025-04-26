package repository

import (
	"github.com/AltheaIX/UMMJacket/internal/domain/auth/repository"
	"testing"
)

func TestGenerateToken(t *testing.T) {
	authRepo := repository.NewAuthRepository()

	accessToken, refreshToken, _ := authRepo.GenerateToken(1, "202410370110031")
	t.Log(accessToken)
	t.Log(refreshToken)
}

func TestValidateToken(t *testing.T) {
	authRepo := repository.NewAuthRepository()

	token, err := authRepo.ValidateJWT("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuaW0iOiIyMDI0MTAzNzAxMTAwMzEiLCJpc3MiOiJqYWNrZXRsYWItdW1tIiwic3ViIjoiMSIsImV4cCI6MTc0NjAyNjYzMywiaWF0IjoxNzQ1NDIxODMzfQ.ZAOed8L0LLsjGkE8II6hDGKXpMosfTfZRoarWjMlDKa")
	t.Log(token, err)
}
