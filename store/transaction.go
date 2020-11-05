package store

import (
	"time"

	"github.com/ahmadfarisfs/krab-core/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TransactionStore struct {
	db *gorm.DB
}

func NewTransactionStore(db *gorm.DB) *TransactionStore {
	return &TransactionStore{
		db: db,
	}
}

func (ac *TransactionStore) ListTransaction(offset, limit int, startTime time.Time, endTime time.Time, accountID []int) ([]model.Transaction, int, error) {
	ret := []model.Transaction{}
	var count int64
	err := ac.db.Model(&ret).Count(&count).Error
	if err != nil {
		return ret, int(count), err
	}
	err = ac.db.Preload("Ledger").Preload("Ledger.Account").
		Where("created_date BETWEEN ? and ?", startTime, endTime).
		Where("account_id in (?)", accountID).
		Offset(offset).Limit(limit).
		Find(&ret).Error
	return ret, int(count), err
}

func (ac *TransactionStore) GetTransactionDetailsbyID(transactionID int) (model.Transaction, error) {
	ret := model.Transaction{}
	err := ac.db.Preload("Ledger").Preload("Ledger.Account").First(&ret, "id = ?", transactionID).Error
	return ret, err
}

func (ac *TransactionStore) GetTransactionDetailsbyCode(transactionCode string) (model.Transaction, error) {
	ret := model.Transaction{}
	err := ac.db.Preload("Ledger").Preload("Ledger.Account").First(&ret, "transaction_code = ?", transactionCode).Error
	return ret, err
}

func (ac *TransactionStore) CreateTransaction(accountFrom int, accountTo int, amount int, remarks string) (model.Transaction, error) {
	//create double entry
	var transactionID int

	err := ac.db.Transaction(func(tx *gorm.DB) error {
		//create entry in transaction db
		trxCode := uuid.New().String()
		trxEntry := model.Transaction{Remarks: remarks, TransactionCode: trxCode}
		if err := tx.Create(&trxEntry).Error; err != nil {
			return err
		}
		//create enty account from
		if err := tx.Create(&model.Ledger{
			AccountID:     accountFrom,
			Credit:        amount,
			TransactionID: int(trxEntry.ID),
		}).Error; err != nil {
			return err
		}
		//create enty account to
		if err := tx.Create(&model.Ledger{
			AccountID:     accountTo,
			Debit:         amount,
			TransactionID: int(trxEntry.ID),
		}).Error; err != nil {
			return err
		}
		transactionID = int(trxEntry.ID)
		//result = trxEntry
		// return nil will commit the whole transaction
		return nil
	})
	if err != nil {
		return model.Transaction{}, err
	}
	return ac.GetTransactionDetailsbyID(transactionID)
}
