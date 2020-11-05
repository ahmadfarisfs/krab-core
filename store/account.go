package store

import (
	"github.com/ahmadfarisfs/krab-core/model"
	"gorm.io/gorm"
)

type AccountStore struct {
	db *gorm.DB
}

func NewAccountStore(db *gorm.DB) *AccountStore {
	return &AccountStore{
		db: db,
	}
}

func (ac *AccountStore) ListAccount(offset, limit int, nameSearch *string) ([]model.Account, int, error) {
	ret := []model.Account{}
	var count int64
	query := ac.db
	if nameSearch != nil {
		query = query.Where("name = ?", *nameSearch)
	}
	err := query.Model(&ret).Count(&count).Error
	if err != nil {
		return ret, int(count), err
	}
	err = query.Offset(offset).Limit(limit).Find(&ret).Error
	return ret, int(count), err
}

func (ac *AccountStore) CreateAccount(name string) (model.Account, error) {
	ret := model.Account{
		AccountName: name,
	}
	err := ac.db.Model(model.Account{}).Create(&ret).Error
	if err != nil {
		return model.Account{}, err
	}
	return ret, nil
}

func (ac *AccountStore) GetAccountDetails(id int) (model.Account, error) {
	ret := model.Account{}
	err := ac.db.First(&ret, "id = ?", id).Error
	return ret, err
}
