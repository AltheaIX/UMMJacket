package handlers

import (
	AuthServices "github.com/AltheaIX/UMMJacket/internal/domain/auth/service"
	JacketsServices "github.com/AltheaIX/UMMJacket/internal/domain/jackets/service"
	StatisticServices "github.com/AltheaIX/UMMJacket/internal/domain/statistic/service"
	TransactionServices "github.com/AltheaIX/UMMJacket/internal/domain/transaction/service"
	UserServices "github.com/AltheaIX/UMMJacket/internal/domain/user/service"
	AuthMiddleware "github.com/AltheaIX/UMMJacket/transport/middleware"
	"github.com/gin-gonic/gin"
)

type Handlers struct {
	authMiddleware AuthMiddleware.AuthMiddleware

	userService        UserServices.UserServices
	authService        AuthServices.AuthServices
	statisticService   StatisticServices.StatisticServices
	jacketsService     JacketsServices.JacketsServices
	transactionService TransactionServices.TransactionServices
}

func NewHandlers(
	authMiddleware AuthMiddleware.AuthMiddleware,
	userService UserServices.UserServices,
	authServices AuthServices.AuthServices,
	statisticService StatisticServices.StatisticServices,
	jacketsService JacketsServices.JacketsServices,
	transactionService TransactionServices.TransactionServices,
) *Handlers {
	return &Handlers{
		authMiddleware:     authMiddleware,
		userService:        userService,
		authService:        authServices,
		statisticService:   statisticService,
		jacketsService:     jacketsService,
		transactionService: transactionService,
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

	jacketsHandler := r.Group("/jackets")
	{
		jacketsHandler.Use(h.authMiddleware.GetUserInfo)
		jacketsHandler.POST("", h.CreateJacket)
		jacketsHandler.POST("/filter", h.FiltersJacket)
		jacketsHandler.PUT("/:id", h.UpdateJacket)
		jacketsHandler.DELETE("/:id", h.DeleteJacket)
	}

	transactionsHandler := r.Group("/transactions")
	{
		transactionsHandler.Use(h.authMiddleware.GetUserInfo)
		transactionsHandler.POST("/new", h.NewTransaction)
	}
}
