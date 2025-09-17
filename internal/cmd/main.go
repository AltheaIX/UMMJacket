package main

import (
	"github.com/AltheaIX/UMMJacket/configs"
	_ "github.com/AltheaIX/UMMJacket/docs"
	"github.com/AltheaIX/UMMJacket/infras"
	AuthRepository "github.com/AltheaIX/UMMJacket/internal/domain/auth/repository"
	AuthServices "github.com/AltheaIX/UMMJacket/internal/domain/auth/service"
	JacketsRepository "github.com/AltheaIX/UMMJacket/internal/domain/jackets/repository"
	JacketsServices "github.com/AltheaIX/UMMJacket/internal/domain/jackets/service"
	StatisticRepo "github.com/AltheaIX/UMMJacket/internal/domain/statistic/repository"
	StatisticServices "github.com/AltheaIX/UMMJacket/internal/domain/statistic/service"
	TransactionRepository "github.com/AltheaIX/UMMJacket/internal/domain/transaction/repository"
	TransactionServices "github.com/AltheaIX/UMMJacket/internal/domain/transaction/service"
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

	db, err := infras.InitMysql(cfg)
	if err != nil {
		log.Fatal().Err(err).Msg("[InitMysql] Failed initializing postgres")
	}

	log.Trace().Msg("[InitMysql] Mysql Initialized")

	userRepo := UserRepository.NewUserRepository(db)
	authRepo := AuthRepository.NewAuthRepository(cfg)
	statisticRepo := StatisticRepo.NewStatisticRepository(db)
	jacketsRepo := JacketsRepository.NewJacketRepository(db)
	transactionRepo := TransactionRepository.NewTransactionRepositoryImpl(db)

	userService := UserServices.NewUserServices(userRepo)
	authService := AuthServices.NewAuthServices(authRepo, userService)
	statisticService := StatisticServices.NewStatisticServices(statisticRepo)
	jacketsService := JacketsServices.NewJacketsServices(jacketsRepo)
	transactionService := TransactionServices.NewTransactionServices(transactionRepo)

	authMiddleware := AuthMiddleware.NewAuthMiddleware(authService)

	http := transport.NewHttp(
		cfg,
		authMiddleware,
		userService,
		authService,
		statisticService,
		jacketsService,
		transactionService,
	)
	http.SetupAndServe()
}
