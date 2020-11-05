package model

import (
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	AccountName string `gorm:"unique"`
	TotalCredit int
	TotalDebit  int
}
