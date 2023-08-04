package usecase

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/syahyudi09/BankingAPI/auth"
	"github.com/syahyudi09/BankingAPI/model"
	"github.com/syahyudi09/BankingAPI/repository"
	"golang.org/x/crypto/bcrypt"
)

type CustomerUsecase interface {
	RegisterCustomer(*model.RegistrasiCustomerModel) error
	LoginCustomer(model.CustomerLogin) (string, error)
	GetAllCustomer() ([]*model.RegistrasiCustomerModel, error)
	
}

type customerUsecaseImpl struct {
	customerRepo repository.CustomerRepository
	auth auth.Service
}

// regisstrasi customer
func (customerUsecase *customerUsecaseImpl) RegisterCustomer(customer *model.RegistrasiCustomerModel) error {
	customer.ID = uuid.NewString()
	
	passhash, err := generatePasswordHash(customer.Password)
	if err != nil {
		return fmt.Errorf("Failed to generate password hash: %w", err)
	}
	customer.Password = passhash

	err = customerUsecase.customerRepo.RegisterCustomer(customer)
	if err != nil {
		return fmt.Errorf("Failed to create customer in repository: %w", err)
	}
	return nil
}

// untuk generate password
func generatePasswordHash(password string) (string, error){
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil{
		return "", err
	}
	return string(hash), nil
}


// login customer
func (customerUsecase *customerUsecaseImpl) LoginCustomer(input model.CustomerLogin) (string, error) {
	email := input.Email
	password := input.Password

	// mencari email 
	customer, err := customerUsecase.customerRepo.FindCustomerByEmail(email)
	if err != nil {
		return "", err
	}

	// jika email tidak ada
	if len(customer.ID) == 0 {
		return "", nil
	}
	
	// mencocokan password
	err = bcrypt.CompareHashAndPassword([]byte(customer.Password), []byte(password))
	if err != nil {
		return "", err
	}

	// mengambil token
	token, err := customerUsecase.auth.GenerateToken(customer.ID)
	if err != nil {
		return "", fmt.Errorf("Failed to generate token: %w", err)
	}

	return token, nil
}

// untuk mendapatkan semua data customer
func(cusstomerUsecase *customerUsecaseImpl) GetAllCustomer() ([]*model.RegistrasiCustomerModel, error) {
	return cusstomerUsecase.customerRepo.GetAllCustomer()
}

func NewCustomerUsecase(customerRepo repository.CustomerRepository, auth auth.Service) CustomerUsecase {
	return &customerUsecaseImpl{
		customerRepo: customerRepo,
		auth: auth,
	}
}