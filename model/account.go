package model

import (
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	AccountName string
	TotalCredit int
	TotalDebit  int
}
