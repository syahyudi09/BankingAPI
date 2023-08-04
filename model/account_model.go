package model

import "time"

type AccountModel struct {
	ID           string
	AccountNumber       string
	AccountType         string
	Balance             float64
	CustomerID          string
	LastTransactionDate time.Time
}

type PaymentAccount struct{
	NamaCustomer string 
	AccountNumberSender string `json:"no_rekening_pengirim"`
	AccountNumberReceiver string `json:"no_rekening_penerima"`
	Balance float64	`json:"jumlah"`
	CreateAt time.Time
}