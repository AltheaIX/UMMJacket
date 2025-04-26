package main

import (
	"github.com/AltheaIX/UMMJacket/configs"
	_ "github.com/AltheaIX/UMMJacket/docs"
	"github.com/AltheaIX/UMMJacket/infras"
	AuthRepository "github.com/AltheaIX/UMMJacket/internal/domain/auth/repository"
	AuthServices "github.com/AltheaIX/UMMJacket/internal/domain/auth/service"
	UserRepository "github.com/AltheaIX/UMMJacket/internal/domain/user/repository"
	UserServices "github.com/AltheaIX/UMMJacket/internal/domain/user/service"
	"github.com/AltheaIX/UMMJacket/shared"
	"github.com/AltheaIX/UMMJacket/transport"
	AuthMiddleware "github.com/AltheaIX/UMMJacket/transport/middleware"
	"github.com/rs/zerolog/log"
	_ "net/http/pprof"
)

func main() {
	cfg := configs.GetConfig()
	shared.InitLogger()

	db, err := infras.InitPostgres(cfg)
	if err != nil {
		log.Fatal().Err(err).Msg("[InitPostgres] Failed initializing postgres")
	}

	log.Trace().Msg("[InitPostgres] Postgres Initialized")

	userRepo := UserRepository.NewUserRepository(db)
	authRepo := AuthRepository.NewAuthRepository(cfg)

	userService := UserServices.NewUserService(userRepo)
	authService := AuthServices.NewAuthServices(authRepo, userService)

	authMiddleware := AuthMiddleware.NewAuthMiddleware(authService)

	http := transport.NewHttp(cfg, authMiddleware, userService, authService)
	http.SetupAndServe()
}
