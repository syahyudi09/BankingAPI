// handler.go

package handler

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/syahyudi09/BankingAPI/auth"
	"github.com/syahyudi09/BankingAPI/manager"
)

type Server interface {
	Run()
}

type server struct {
	usecaseManager manager.UsecaseManager
	engine         *gin.Engine
	authService auth.Service
}

func (s *server) Run() {
	authService := auth.NewServiceJWT()
	middleware := auth.NewMiddleware(authService)

	authenticated := s.engine.Group("/")
	authenticated.Use(middleware.AuthMiddleware())

	NewCustomerHandler(s.engine, s.usecaseManager.GetCustomerUsecase(), s.usecaseManager.GetAccountUsecase())
	NewPaymantHandler(s.engine, s.usecaseManager.GetPaymentUsecase())
	NewMarchantHandler(s.engine, s.usecaseManager.GetMarchantUsecase())
	s.engine.Run(":8080")
	
}

func NewServer() Server {

	infra := manager.NewInfraManager()
	filepath := infra.GetFile()
	authService := auth.NewServiceJWT()
	repository := manager.NewRepositoryManager(filepath)
	usecase := manager.NewUsecaseManager(repository,authService)
	
	engine := gin.Default()

	store := cookie.NewStore([]byte(auth.SECRET_KEY))
	engine.Use(sessions.Sessions("mysession", store))

	return &server{
		usecaseManager: usecase,
		engine:         engine,
		authService: authService,
	}
}
