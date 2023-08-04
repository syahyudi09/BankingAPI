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

