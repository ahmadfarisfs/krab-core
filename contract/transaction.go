package contract

import (
	"time"

	"github.com/ahmadfarisfs/krab-core/model"
)

type TransactionStore interface {
	GetTransactionDetailsbyID(transactionID int) (model.Transaction, error)
	GetTransactionDetailsbyCode(transactionCode string) (model.Transaction, error)
	ListTransaction(offset, limit int, startTime time.Time, endTime time.Time, accountID []int) ([]model.Transaction, int, error)
	CreateTransaction(accountFrom int, accountTo int, amount int, remarks string) (model.Transaction, error)
}
