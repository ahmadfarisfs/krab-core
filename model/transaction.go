package model

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	TransactionCode string
	Remarks         string
	Ledger          []Ledger
}
