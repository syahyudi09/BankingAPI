package repository

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/syahyudi09/BankingAPI/model"
)

type AccountRepository interface {
	Create(*model.AccountModel) error
}

type accountRepository struct {
	filepath string
}

func (ar *accountRepository) Create(account *model.AccountModel) error {
	accounts, err := readAccountFromFile("data/account.json")
	if err != nil {
		return err
	}

	accounts = append(accounts, account)

	err = saveAccountToFile("data/account.json", accounts)
	if err != nil {
		return err
	}

	return nil
}

func readAccountFromFile(filepath string) ([]*model.AccountModel, error) {
	file, err := os.OpenFile(filepath, os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var accounts []*model.AccountModel
	if len(data) > 0 {
		err = json.Unmarshal(data, &accounts)
		if err != nil {
			return nil, err
		}
	}

	return accounts, nil
}

func saveAccountToFile(filepath string, accounts []*model.AccountModel) error {
	file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	data, err := json.MarshalIndent(accounts, "", "  ")
	if err != nil {
		return err
	}

	_, err = file.Write(data)
	if err != nil {
		return err
	}

	return nil
}

func writeAccountsToFile(filepath string, accounts []*model.AccountModel) error {
	data, err := json.Marshal(accounts)
	if err != nil {
		return err
	}

	file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = json.MarshalIndent(accounts, "", "  ")
	if err != nil {
		return err
	}

	_, err = file.Write(data)
	return err
}

func NewAccountRepo(filepath string) AccountRepository {
	return &accountRepository{
		filepath: filepath,
	}
}
