package repository

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sync"
	"time"

	"github.com/syahyudi09/BankingAPI/model"
)

type PaymentRepository interface {
	Payment(model.PaymentAccount) error
}

type paymentRepositoryImpl struct {
	account map[int]*model.PaymentAccount
	mutex sync.Mutex
}

func (pm *paymentRepositoryImpl) Payment(payment model.PaymentAccount) error {
	pm.mutex.Lock()
	defer pm.mutex.Unlock()


	// untuk membaca di account.json
	accounts, err := readAccountFromFile("data/account.json")
	if err != nil {
		return err
	}

	var sender, receiver *model.AccountModel

	for _, acc := range accounts {
		if acc.AccountNumber == payment.AccountNumberSender {
			sender = acc
		}
		if acc.AccountNumber == payment.AccountNumberReceiver {
			receiver = acc
		}
	}

	if sender == nil {
		return fmt.Errorf("akun pengirim %s tidak ditemukan", payment.AccountNumberSender)
	}

	if receiver == nil {
		return fmt.Errorf("akun penerima %s tidak ditemukan", payment.AccountNumberReceiver)
	}

	if receiver.Balance < payment.Balance {
		return fmt.Errorf("saldo tidak mencukupi")
	}

	receiver.Balance -= payment.Balance
	sender.Balance += payment.Balance

	// untuk menulis ulang balance
	err = writeAccountsToFile("data/account.json", accounts)
	if err != nil {
		return err
	}

	paymentHistory, err := readPaymentHistoryFromFile()
	if err != nil {
		return err
	}

	paymentHistory = append(paymentHistory, model.PaymentAccount{
		AccountNumberSender:   sender.AccountNumber,
		AccountNumberReceiver: receiver.AccountNumber,
		Balance:     payment.Balance,
		CreateAt: time.Now(),
	})

	err = savePaymentHistoryToFile(paymentHistory)
	if err != nil {
		return err
	}

	return nil
}

func readPaymentHistoryFromFile() ([]model.PaymentAccount, error) {
	filepath := "data/payment.json"

	
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	var paymentHistory []model.PaymentAccount
	if len(data) > 0 {
		err = json.Unmarshal(data, &paymentHistory)
		if err != nil {
			return nil, err
		}
	}
	return paymentHistory, nil
}

func savePaymentHistoryToFile(paymentHistory []model.PaymentAccount) error {
	filepath := "data/payment.json"

	// Simpan data history pembayaran kembali ke file
	data, err := json.MarshalIndent(paymentHistory, "", "  ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filepath, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

func NewPaymentRepository(account map[int]*model.PaymentAccount) PaymentRepository {
	return &paymentRepositoryImpl{
		account: account,
	}
}