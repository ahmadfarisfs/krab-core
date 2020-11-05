package model

import "gorm.io/gorm"

type Ledger struct {
	gorm.Model
	Debit         int
	Credit        int
	AccountID     int
	Account       Account
	TransactionID int
	Transaction   Transaction
}
