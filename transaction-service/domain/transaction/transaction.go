package transaction

import (
	"context"

	"github.com/google/uuid"
	"github.com/nafisalfiani/p3-final-project/lib/codes"
	"github.com/nafisalfiani/p3-final-project/lib/errors"
	"github.com/nafisalfiani/p3-final-project/transaction-service/entity"
	"gorm.io/gorm"
)

type transaction struct {
	db *gorm.DB
}

type Interface interface {
	List(ctx context.Context) ([]entity.Transaction, error)
	Get(ctx context.Context, filter entity.Transaction) (entity.Transaction, error)
	Create(ctx context.Context, transaction entity.Transaction) (entity.Transaction, error)
	Update(ctx context.Context, transaction entity.Transaction) (entity.Transaction, error)
	Delete(ctx context.Context, id string) error
}

// Init create transaction repository
func Init(db *gorm.DB) Interface {
	return &transaction{
		db: db,
	}
}

// List returns list of transactions
func (w *transaction) List(ctx context.Context) ([]entity.Transaction, error) {
	transactions := []entity.Transaction{}
	if err := w.db.Find(&transactions).Error; err != nil {
		return transactions, errors.NewWithCode(codes.CodeSQLRead, err.Error())
	}

	return transactions, nil
}

// Get returns specific transaction
func (w *transaction) Get(ctx context.Context, filter entity.Transaction) (entity.Transaction, error) {
	transaction := entity.Transaction{}

	err := w.db.First(&transaction).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return transaction, errors.NewWithCode(codes.CodeSQLRecordDoesNotExist, err.Error())
	} else if err != nil {
		return transaction, errors.NewWithCode(codes.CodeSQLRead, err.Error())
	}

	return transaction, nil
}

// Create creates new data
func (w *transaction) Create(ctx context.Context, transaction entity.Transaction) (entity.Transaction, error) {
	err := w.db.Create(&transaction).Error
	if err != nil && errors.Is(err, gorm.ErrDuplicatedKey) {
		return transaction, errors.NewWithCode(codes.CodeConflict, err.Error())
	} else if err != nil {
		return transaction, errors.NewWithCode(codes.CodeSQLInsert, err.Error())
	}

	return transaction, nil
}

// Update updates existing data
func (w *transaction) Update(ctx context.Context, transaction entity.Transaction) (entity.Transaction, error) {
	if err := w.db.Save(&transaction).Error; err != nil {
		return transaction, errors.NewWithCode(codes.CodeSQLUpdate, err.Error())
	}

	return transaction, nil
}

// Delete deletes existing data
func (w *transaction) Delete(ctx context.Context, id string) error {
	uid, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	if err := w.db.Delete(&entity.Transaction{Id: uid}).Error; err != nil {
		return errors.NewWithCode(codes.CodeSQLDelete, err.Error())
	}

	return nil
}
