package manager

import (
	"sync"

	"github.com/syahyudi09/BankingAPI/auth"
	"github.com/syahyudi09/BankingAPI/usecase"
)

type UsecaseManager interface {
	GetCustomerUsecase() usecase.CustomerUsecase
	GetAccountUsecase() usecase.AccountUsecase
	GetPaymentUsecase() usecase.PaymentUsecase
	GetMarchantUsecase() usecase.MarchantUsecase
}

type usecaseManagerImpl struct {
	repositoryManager RepositoryManager
	customerUsecase   usecase.CustomerUsecase
	accountUsecase usecase.AccountUsecase
	paymentUsecase usecase.PaymentUsecase
	authService auth.Service
	marchantUsecase usecase.MarchantUsecase
}

var onceLoadCustomerUsecase sync.Once
var onceLoadAccountUsecase sync.Once
var onceLoadPaymentUsecase sync.Once
var onceLoadMarchantUsecase sync.Once

func (um *usecaseManagerImpl) GetCustomerUsecase() usecase.CustomerUsecase {
	onceLoadCustomerUsecase.Do(func() {
		um.customerUsecase = usecase.NewCustomerUsecase(
			um.repositoryManager.GetCustomerRepo(),
			um.authService,
		)
	})
	return um.customerUsecase
}

func (um *usecaseManagerImpl) GetAccountUsecase() usecase.AccountUsecase {
	onceLoadAccountUsecase.Do(func() {
		um.accountUsecase = usecase.NewAccountUsecase(
			um.repositoryManager.GetAccountRepo(),
		)
	})
	return um.accountUsecase
}

func(um *usecaseManagerImpl) GetPaymentUsecase() usecase.PaymentUsecase{
	onceLoadPaymentUsecase.Do(func() {
		um.paymentUsecase = usecase.NewPaymentUsecase(
			um.repositoryManager.GetPaymentRepo(),
			um.authService,
		)
	})
	return um.paymentUsecase
}

func (um *usecaseManagerImpl) GetMarchantUsecase() usecase.MarchantUsecase {
	onceLoadMarchantUsecase.Do(func() {
		um.marchantUsecase = usecase.NewMerchantUsecase(
			um.repositoryManager.GetMarchantRepo(),
			um.authService,
			
		)
	})
	return um.marchantUsecase
}

func NewUsecaseManager(rm RepositoryManager, authService auth.Service) UsecaseManager {
	return &usecaseManagerImpl{
		repositoryManager: rm,
		authService: authService,
	}
}
