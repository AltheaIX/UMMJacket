package services

import (
	"context"
	"github.com/AltheaIX/UMMJacket/configs"
	"github.com/AltheaIX/UMMJacket/infras"
	AuthRepository "github.com/AltheaIX/UMMJacket/internal/domain/auth/repository"
	AuthServices "github.com/AltheaIX/UMMJacket/internal/domain/auth/service"
	"github.com/AltheaIX/UMMJacket/internal/domain/user/repository"
	"github.com/AltheaIX/UMMJacket/internal/domain/user/service"
	"github.com/AltheaIX/UMMJacket/shared"
	"testing"
)

func TestLogin(t *testing.T) {
	cfg := configs.GetConfig()
	db, _ := infras.InitPostgres(cfg)
	defer db.Close()

	userRepo := repository.NewUserRepository(db)
	authRepo := AuthRepository.NewAuthRepository(cfg)

	userService := service.NewUserService(userRepo)

	authService := AuthServices.NewAuthServices(authRepo, userService)
	t.Log(authService.Login(context.Background(), "202410370110031", "Malik"))
}

func TestAuthenticateToken(t *testing.T) {
	cfg := configs.GetConfig()
	db, _ := infras.InitPostgres(cfg)
	defer db.Close()

	userRepo := repository.NewUserRepository(db)
	authRepo := AuthRepository.NewAuthRepository(cfg)

	userService := service.NewUserService(userRepo)

	authService := AuthServices.NewAuthServices(authRepo, userService)
	userInfo, err := authService.AuthenticateToken(context.Background(), "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuaW0iOiIyMDI0MTAzNzAxMTAwMzEiLCJpc3MiOiJqYWNrZXRsYWItdW1tIiwic3ViIjoiMSIsImV4cCI6MTc0NTQ2NTMyMCwiaWF0IjoxNzQ1NDYxNzIwfQ.jtQyZdjZv_OOv4qQGm82x5kjjoVhKZb7av-UJ1PPacw")
	if err != nil {
		code := shared.GetCode(err)
		t.Error(code, err.Error())
		return
	}

	t.Log(userInfo)
}
