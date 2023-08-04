package usecase

import (
	"fmt"
	"time"
	"math/rand"
	"github.com/google/uuid"
	"github.com/syahyudi09/BankingAPI/model"
	"github.com/syahyudi09/BankingAPI/repository"
)

type AccountUsecase interface {
	CreateAccount(*model.AccountModel) error
}

type accountUsecaeImpl struct {
	accountRepo repository.AccountRepository
	
}

func (accountUsecase *accountUsecaeImpl) CreateAccount(account *model.AccountModel) error {
	account.ID = uuid.NewString()

	randomAccountNumber := generateRandomAccountNumber(10)
	account.AccountNumber = randomAccountNumber

	account.Balance = 500000
	if account.Balance < 5000000 {
		fmt.Println("Top Up Harus Rp. 500.000")
	}

	account.LastTransactionDate = time.Now()

	return accountUsecase.accountRepo.Create(account)
}

func generateRandomAccountNumber(length int) string {
	const numbers = "0123456789"
	rand.Seed(time.Now().UnixNano())
	accountNumber := make([]byte, length)
	for i := range accountNumber {
		accountNumber[i] = numbers[rand.Intn(len(numbers))]
	}
	return string(accountNumber)
}

func NewAccountUsecase(accountRepo repository.AccountRepository) AccountUsecase {
	return &accountUsecaeImpl{
		accountRepo: accountRepo,
	}
}