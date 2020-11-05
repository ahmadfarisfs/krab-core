package handler

import (
	"github.com/ahmadfarisfs/krab-core/contract"
)

type Handler struct {
	accountStore     contract.AccountStore
	transactionStore contract.TransactionStore
}

func NewHandler(as contract.AccountStore, ts contract.TransactionStore) *Handler {
	return &Handler{
		accountStore:     as,
		transactionStore: ts,
	}
}
