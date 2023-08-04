package manager

import (
	"sync"

	"github.com/syahyudi09/BankingAPI/model"
	"github.com/syahyudi09/BankingAPI/repository"
)

type RepositoryManager interface {
	GetCustomerRepo() repository.CustomerRepository
	GetAccountRepo() repository.AccountRepository
	GetPaymentRepo() repository.PaymentRepository
	GetMarchantRepo() repository.MarchantRepository
}

type repositoryManagerImp struct {
	filepath string
	customerRepo repository.CustomerRepository
	accountRepo repository.AccountRepository
	paymentRepo repository.PaymentRepository
	paymentAccounts map[int]*model.PaymentAccount
	marchantRepo repository.MarchantRepository

}

var onceLoadCustomerRepo sync.Once
var onceLoadAccountRepo sync.Once
var onceLoadPaymentRepo sync.Once
var onceLoadMarchantRepo sync.Once

func(rm *repositoryManagerImp) GetCustomerRepo() repository.CustomerRepository {
	onceLoadCustomerRepo.Do(func ()  {
		rm.customerRepo = repository.NewCustomerRepository(rm.filepath)
	})
	return rm.customerRepo
}

func (rm *repositoryManagerImp) GetAccountRepo() repository.AccountRepository {
	onceLoadAccountRepo.Do(func() {
		rm.accountRepo = repository.NewAccountRepo(rm.filepath)
	})
	return rm.accountRepo
}

func (rm *repositoryManagerImp) GetPaymentRepo() repository.PaymentRepository{
	onceLoadPaymentRepo.Do(func() {
		rm.paymentRepo = repository.NewPaymentRepository(rm.paymentAccounts)
	})
	return rm.paymentRepo
}

func (rm *repositoryManagerImp) GetMarchantRepo() repository.MarchantRepository{
	onceLoadMarchantRepo.Do(func() {
		rm.marchantRepo = repository.NewMerchantRepository(rm.filepath)
	})
	return rm.marchantRepo
}

func NewRepositoryManager(filepath string) RepositoryManager {
	return &repositoryManagerImp{
	}
}