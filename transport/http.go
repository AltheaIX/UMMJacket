package transport

import (
	"github.com/AltheaIX/UMMJacket/configs"
	AuthServices "github.com/AltheaIX/UMMJacket/internal/domain/auth/service"
	UserServices "github.com/AltheaIX/UMMJacket/internal/domain/user/service"
	"github.com/AltheaIX/UMMJacket/internal/handlers"
	AuthMiddleware "github.com/AltheaIX/UMMJacket/transport/middleware"
	ginzerolog "github.com/dn365/gin-zerolog"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

type HTTP struct {
	cfg            *configs.Config
	authMiddleware AuthMiddleware.AuthMiddleware

	userService UserServices.UserService
	authService AuthServices.AuthServices
}

func NewHttp(cfg *configs.Config, authMiddleware AuthMiddleware.AuthMiddleware, userService UserServices.UserService, authServices AuthServices.AuthServices) *HTTP {
	return &HTTP{cfg: cfg, authMiddleware: authMiddleware, userService: userService, authService: authServices}
}

func (h *HTTP) SetupAndServe() {
	gin.SetMode(h.cfg.Server.Mode)
	r := gin.New()
	r.Use(ginzerolog.Logger("gin"), gin.Recovery())
	v1 := r.Group("/v1")
	{
		handler := handlers.NewHandlers(h.authMiddleware, h.userService, h.authService)
		handler.RouterV1(v1)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	go func() {
		// Serve pprof on :6060
		if err := http.ListenAndServe("localhost:6060", nil); err != nil {
			log.Fatal().Err(err).Msg("[pprof] failed to start")
		}
	}()

	r.Run(":" + h.cfg.Server.Port)
}
