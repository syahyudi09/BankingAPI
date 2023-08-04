package usecase

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/syahyudi09/BankingAPI/auth"
	"github.com/syahyudi09/BankingAPI/model"
	"github.com/syahyudi09/BankingAPI/repository"
)

type MarchantUsecase interface {
	CreateMarchant(*model.MarchantModel) error
}

type merchantUsecase struct {
	merchantRepo repository.MarchantRepository
	auth auth.Service
}

func (m *merchantUsecase) CreateMarchant(marchant *model.MarchantModel) error {
	marchant.ID = uuid.NewString()
	if marchant.NamaMarchant == "" {
		fmt.Println("Nama Marchant Tidak Boleh Kosong")
	}

	if len(marchant.NamaMarchant) < 4 {
		fmt.Println("Nama Harus Lebih dari 4 kararkter")
	}
	
	err := m.merchantRepo.CreateMarchant(marchant)
	if err != nil {
		return err
	}

	return nil
}

func NewMerchantUsecase(merchantRepo repository.MarchantRepository, auth auth.Service) MarchantUsecase {
	return &merchantUsecase{
		merchantRepo: merchantRepo,
		auth: auth,
	}
}
