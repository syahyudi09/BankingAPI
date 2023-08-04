package usecase

import (
	"time"

	"github.com/syahyudi09/BankingAPI/model"
	"github.com/syahyudi09/BankingAPI/repository"
)

type PaymentUsecase interface {
	Payment (model.PaymentAccount) error
}

type paymentUsecaseImpl struct {
	PaymentRepo repository.PaymentRepository
}

func(pu *paymentUsecaseImpl) Payment(payment model.PaymentAccount) error {
	payment.CreateAt = time.Now()
	return pu.PaymentRepo.Payment(payment)
}

func NewPaymentUsecase(paymentRepo repository.PaymentRepository) PaymentUsecase {
	return &paymentUsecaseImpl{
		PaymentRepo: paymentRepo,
	}
}
