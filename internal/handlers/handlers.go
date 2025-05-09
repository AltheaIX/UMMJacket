package handlers

import (
	AuthServices "github.com/AltheaIX/UMMJacket/internal/domain/auth/service"
	StatisticServices "github.com/AltheaIX/UMMJacket/internal/domain/statistic/service"
	UserServices "github.com/AltheaIX/UMMJacket/internal/domain/user/service"
	AuthMiddleware "github.com/AltheaIX/UMMJacket/transport/middleware"
	"github.com/gin-gonic/gin"
)

type Handlers struct {
	authMiddleware AuthMiddleware.AuthMiddleware

	userService      UserServices.UserServices
	authService      AuthServices.AuthServices
	statisticService StatisticServices.StatisticServices
}

func NewHandlers(
	authMiddleware AuthMiddleware.AuthMiddleware,
	userService UserServices.UserServices,
	authServices AuthServices.AuthServices,
	statisticService StatisticServices.StatisticServices,
) *Handlers {
	return &Handlers{
		authMiddleware:   authMiddleware,
		userService:      userService,
		authService:      authServices,
		statisticService: statisticService,
	}
}

func (h *Handlers) RouterV1(r *gin.RouterGroup) {
	authHandler := r.Group("/auth")
	{
		authHandler.POST("/login", h.Login)
		loginRequireAuthHandler := authHandler.Group("")
		{
			loginRequireAuthHandler.Use(h.authMiddleware.GetUserInfo)
			loginRequireAuthHandler.POST("/refresh", h.Refresh)
			loginRequireAuthHandler.POST("/current", h.CurrentUser)
		}
	}

	statisticsHandler := r.Group("/statistics")
	{
		statisticsHandler.Use(h.authMiddleware.GetUserInfo)
		statisticsHandler.GET("/dashboard", h.Dashboard)
	}
}
