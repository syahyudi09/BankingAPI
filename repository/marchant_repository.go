package repository

import (
	"encoding/json"
	"io/ioutil"
	"os"


	"github.com/syahyudi09/BankingAPI/model"
)

type MarchantRepository interface {
	CreateMarchant(*model.MarchantModel) error
}

type marchantRepository struct {
	filepath string
}

func (m *marchantRepository) CreateMarchant(marchant *model.MarchantModel ) error {
	marchants, err := readMarchantFromFile("data/marchant.json")
	if err != nil {
		return err
	}

	marchants = append(marchants, marchant)

	err = saveMarchantToFile("data/marchant.json", marchants)
	if err != nil {
		return err
	}

	return nil
}

func readMarchantFromFile(filepath string) ([]*model.MarchantModel, error) {
	file, err := os.OpenFile(filepath, os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var marchants []*model.MarchantModel
	if len(data) > 0 {
		err = json.Unmarshal(data, &marchants)
		if err != nil {
			return nil, err
		}
	}

	return marchants, nil
}

func saveMarchantToFile(filepath string, marchantSave []*model.MarchantModel) error {
	file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	data, err := json.MarshalIndent(marchantSave, "", "  ")
	if err != nil {
		return err
	}

	_, err = file.Write(data)
	if err != nil {
		return err
	}

	return nil
}

func NewMerchantRepository(filepath string) MarchantRepository {
	return &marchantRepository{
		filepath: filepath,
	}
}
