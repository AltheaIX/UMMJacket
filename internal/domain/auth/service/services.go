package service

import (
	"context"
	AuthModel "github.com/AltheaIX/UMMJacket/internal/domain/auth/model"
	AuthRepository "github.com/AltheaIX/UMMJacket/internal/domain/auth/repository"
	UserServices "github.com/AltheaIX/UMMJacket/internal/domain/user/service"
)

type AuthServices interface {
	IAuthServices

	AuthenticateToken(ctx context.Context, token string) (*AuthModel.Claims, error)
}

type AuthServicesImpl struct {
	authRepo    AuthRepository.AuthRepository
	userService UserServices.UserService
}

func NewAuthServices(authRepo AuthRepository.AuthRepository, userService UserServices.UserService) AuthServices {
	return &AuthServicesImpl{authRepo: authRepo, userService: userService}
}

func (a *AuthServicesImpl) AuthenticateToken(ctx context.Context, token string) (*AuthModel.Claims, error) {
	userInfo, err := a.authRepo.ValidateJWT(token)
	if err != nil {
		return nil, err
	}

	return userInfo, nil
}
