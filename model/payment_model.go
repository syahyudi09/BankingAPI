package model

import "time"

type PaymentAccount struct{
	AccountNumberSender string `json:"no_rekening_pengirim"`
	AccountNumberReceiver string `json:"no_rekening_penerima"`
	Balance float64	`json:"jumlah"`
	CreateAt time.Time
}