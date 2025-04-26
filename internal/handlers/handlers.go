package handlers

import (
	AuthServices "github.com/AltheaIX/UMMJacket/internal/domain/auth/service"
	UserServices "github.com/AltheaIX/UMMJacket/internal/domain/user/service"
	AuthMiddleware "github.com/AltheaIX/UMMJacket/transport/middleware"
	"github.com/gin-gonic/gin"
)

type Handlers struct {
	authMiddleware AuthMiddleware.AuthMiddleware

	userService UserServices.UserService
	authService AuthServices.AuthServices
}

func NewHandlers(authMiddleware AuthMiddleware.AuthMiddleware, userService UserServices.UserService, authServices AuthServices.AuthServices) *Handlers {
	return &Handlers{authMiddleware: authMiddleware, userService: userService, authService: authServices}
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
}
