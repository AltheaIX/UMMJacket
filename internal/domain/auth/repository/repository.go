package repository

import (
	"github.com/AltheaIX/UMMJacket/configs"
)

type AuthRepository interface {
	IAuthRepository
}

type AuthRepositoryImpl struct {
	cfg *configs.Config
}

func NewAuthRepository(cfg *configs.Config) AuthRepository {
	return &AuthRepositoryImpl{cfg: cfg}
}
