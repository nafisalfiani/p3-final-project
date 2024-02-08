package wallet

import (
	"context"

	"github.com/nafisalfiani/p3-final-project/lib/codes"
	"github.com/nafisalfiani/p3-final-project/lib/errors"
	"github.com/nafisalfiani/p3-final-project/transaction-service/entity"
	"gorm.io/gorm"
)

type wallet struct {
	db *gorm.DB
}

type Interface interface {
	List(ctx context.Context) ([]entity.Wallet, error)
	Get(ctx context.Context, filter entity.Wallet) (entity.Wallet, error)
	Create(ctx context.Context, wallet entity.Wallet) (entity.Wallet, error)
	Update(ctx context.Context, wallet entity.Wallet) (entity.Wallet, error)
	Delete(ctx context.Context, id string) error
}

// Init create wallet repository
func Init(db *gorm.DB) Interface {
	return &wallet{
		db: db,
	}
}

// List returns list of wallets
func (w *wallet) List(ctx context.Context) ([]entity.Wallet, error) {
	wallets := []entity.Wallet{}
	if err := w.db.Find(&wallets).Error; err != nil {
		return wallets, errors.NewWithCode(codes.CodeSQLRead, err.Error())
	}

	return wallets, nil
}

// Get returns specific wallet
func (w *wallet) Get(ctx context.Context, filter entity.Wallet) (entity.Wallet, error) {
	wallet := entity.Wallet{}
	// cond := w.db.Where("1 = 1")

	// if filter.

	err := w.db.First(&wallet).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return wallet, errors.NewWithCode(codes.CodeSQLRecordDoesNotExist, err.Error())
	} else if err != nil {
		return wallet, errors.NewWithCode(codes.CodeSQLRead, err.Error())
	}

	return wallet, nil
}

// Create creates new data
func (w *wallet) Create(ctx context.Context, wallet entity.Wallet) (entity.Wallet, error) {
	err := w.db.Create(&wallet).Error
	if err != nil && errors.Is(err, gorm.ErrDuplicatedKey) {
		return wallet, errors.NewWithCode(codes.CodeConflict, err.Error())
	} else if err != nil {
		return wallet, errors.NewWithCode(codes.CodeSQLInsert, err.Error())
	}

	return wallet, nil
}

// Update updates existing data
func (w *wallet) Update(ctx context.Context, wallet entity.Wallet) (entity.Wallet, error) {
	if err := w.db.Save(&wallet).Error; err != nil {
		return wallet, errors.NewWithCode(codes.CodeSQLUpdate, err.Error())
	}

	return wallet, nil
}

// Delete deletes existing data
func (w *wallet) Delete(ctx context.Context, id string) error {
	if err := w.db.Delete(&entity.Wallet{Id: id}).Error; err != nil {
		return errors.NewWithCode(codes.CodeSQLDelete, err.Error())
	}

	return nil
}
