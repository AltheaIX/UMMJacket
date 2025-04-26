package service

import (
	"context"
	"github.com/AltheaIX/UMMJacket/internal/domain/auth/model"
	"github.com/AltheaIX/UMMJacket/shared"
	"github.com/AltheaIX/UMMJacket/shared/crypt"
	"github.com/AltheaIX/UMMJacket/shared/filter"
	"strconv"
	"time"
)

type IAuthServices interface {
	Login(ctx context.Context, nim string, password string) (string, string, error)
	Refresh(ctx context.Context, userInfo *model.Claims) (string, error)
}

func (a *AuthServicesImpl) Login(ctx context.Context, nim string, password string) (string, string, error) {
	filters := &filter.Filters{Filter: []filter.Filter{
		{
			Field:    "nim",
			Operator: "eq",
			Value:    nim,
		},
	}}

	user, err := a.userService.GetUserService(ctx, filters)
	if err != nil {
		return "", "", err
	}

	if user == nil || !crypt.CheckPasswordHash(password, user.Password) {
		return "", "", &shared.AppError{Code: 200, Message: "Invalid username or password"}
	}

	accessToken, refreshToken, err := a.authRepo.GenerateToken(user.Id, user.Nim)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (a *AuthServicesImpl) Refresh(ctx context.Context, userInfo *model.Claims) (string, error) {
	expiredAt, _ := userInfo.GetExpirationTime()
	if expiredAt.Unix() < time.Now().Unix() {
		return "", &shared.AppError{Code: 401, Message: "Token is expired"}
	}

	userId, err := strconv.Atoi(userInfo.Subject)
	if err != nil {
		return "", err
	}

	accessToken, _, err := a.authRepo.GenerateToken(userId, userInfo.Nim)
	if err != nil {
		return "", err
	}

	return accessToken, nil
}
