package usecase

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/syahyudi09/BankingAPI/model"
	"github.com/syahyudi09/BankingAPI/repository"
)

type MarchantUsecase interface {
	Create(*model.MarchantModel) error
}

type merchantUsecase struct {
	merchantRepo repository.MarchantRepository
}

func (m *merchantUsecase) Create(marchant *model.MarchantModel) error {
	marchant.ID = uuid.NewString()
	if marchant.NamaMarchant == "" {
		fmt.Println("Nama Marchant Tidak Boleh Kosong")
	}

	if len(marchant.NamaMarchant) < 4 {
		fmt.Println("Nama Harus Lebih dari 4 kararkter")
	}
	return m.merchantRepo.Create(marchant)
}

func NewMerchantUsecase(merchantRepo repository.MarchantRepository) MarchantUsecase {
	return &merchantUsecase{
		merchantRepo: merchantRepo,
	}
}
