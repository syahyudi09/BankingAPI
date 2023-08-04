package repository

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/syahyudi09/BankingAPI/model"
)

type CustomerRepository interface {
	RegisterCustomer(*model.RegistrasiCustomerModel) error
	FindCustomerByEmail(string) (*model.RegistrasiCustomerModel, error)
	GetAllCustomer() ([]*model.RegistrasiCustomerModel, error)
}

type customerRepositoryImpl struct {
	filepath string
}

func (customerRepo *customerRepositoryImpl) RegisterCustomer(customer *model.RegistrasiCustomerModel) error {
	customers, err := readCustomersFromFile("data/customer.json")
	if err != nil {
		return err
	}

	customers = append(customers, customer)

	err = saveCustomersToFile("data/customer.json", customers)
	if err != nil {
		return err
	}

	return nil
}

func (customerRepo *customerRepositoryImpl) FindCustomerByEmail(email string) (*model.RegistrasiCustomerModel, error){
	customers, err := readCustomersFromFile("data/customer.json")
	if err != nil{
		return nil, err
	}

	for _, customer := range customers{
		if customer.Email == email {
			return customer, nil
		}
	}

	return nil, fmt.Errorf("customer dengan email %s tersebut tidak ditemukan", email)
}

func(customerRepo *customerRepositoryImpl) GetAllCustomer() ([]*model.RegistrasiCustomerModel, error){
	customers, err := readCustomersFromFile("data/customer.json")
	if err != nil{
		return nil, err
	}
	return customers, nil
}

func readCustomersFromFile(filepath string) ([]*model.RegistrasiCustomerModel, error) {
	file, err := os.OpenFile(filepath, os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var customers []*model.RegistrasiCustomerModel
	if len(data) > 0 {
		err = json.Unmarshal(data, &customers)
		if err != nil {
			return nil, err
		}
	}

	return customers, nil
}

func saveCustomersToFile(filepath string, customers []*model.RegistrasiCustomerModel) error {
	file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	data, err := json.MarshalIndent(customers, "", "  ")
	if err != nil {
		return err
	}

	_, err = file.Write(data)
	if err != nil {
		return err
	}

	return nil
}

func NewCustomerRepository(filepath string) CustomerRepository {
	return &customerRepositoryImpl{
		filepath: filepath,
}
}