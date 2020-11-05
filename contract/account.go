package contract

import (
	"github.com/ahmadfarisfs/krab-core/model"
)

type AccountStore interface {
	ListAccount(offset, limit int, nameSearch *string) ([]model.Account, int, error)
	CreateAccount(name string) (model.Account, error)
	GetAccountDetails(id int) (model.Account, error)
}
