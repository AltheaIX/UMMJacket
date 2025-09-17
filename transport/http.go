package transport

import (
	"github.com/AltheaIX/UMMJacket/configs"
	AuthServices "github.com/AltheaIX/UMMJacket/internal/domain/auth/service"
	JacketsServices "github.com/AltheaIX/UMMJacket/internal/domain/jackets/service"
	StatisticServices "github.com/AltheaIX/UMMJacket/internal/domain/statistic/service"
	TransactionServices "github.com/AltheaIX/UMMJacket/internal/domain/transaction/service"
	UserServices "github.com/AltheaIX/UMMJacket/internal/domain/user/service"
	"github.com/AltheaIX/UMMJacket/internal/handlers"
	AuthMiddleware "github.com/AltheaIX/UMMJacket/transport/middleware"
	ginzerolog "github.com/dn365/gin-zerolog"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

type HTTP struct {
	cfg            *configs.Config
	authMiddleware AuthMiddleware.AuthMiddleware

	userService        UserServices.UserServices
	authService        AuthServices.AuthServices
	statisticService   StatisticServices.StatisticServices
	jacketsService     JacketsServices.JacketsServices
	transactionService TransactionServices.TransactionServices
}

func NewHttp(
	cfg *configs.Config,
	authMiddleware AuthMiddleware.AuthMiddleware,
	userService UserServices.UserServices,
	authServices AuthServices.AuthServices,
	statisticService StatisticServices.StatisticServices,
	jacketsService JacketsServices.JacketsServices,
	transactionService TransactionServices.TransactionServices,
) *HTTP {
	return &HTTP{
		cfg:                cfg,
		authMiddleware:     authMiddleware,
		userService:        userService,
		authService:        authServices,
		statisticService:   statisticService,
		jacketsService:     jacketsService,
		transactionService: transactionService,
	}
}

func (h *HTTP) SetupAndServe() {
	gin.SetMode(h.cfg.Server.Mode)
	r := gin.New()
	r.Use(cors.Default())
	r.Use(ginzerolog.Logger("gin"), gin.Recovery())
	v1 := r.Group("/v1")
	{
		handler := handlers.NewHandlers(
			h.authMiddleware,
			h.userService,
			h.authService,
			h.statisticService,
			h.jacketsService,
			h.transactionService,
		)
		handler.RouterV1(v1)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	go func() {
		// Serve pprof on :6060
		if err := http.ListenAndServe("localhost:6060", nil); err != nil {
			log.Fatal().Err(err).Msg("[pprof] failed to start")
		}
	}()

	err := r.Run(":" + h.cfg.Server.Port)
	if err != nil {
		log.Fatal().Err(err).Msg("[SetupAndServe] failed to start")
	}
}
