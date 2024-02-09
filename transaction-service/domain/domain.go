package domain

import (
	"github.com/nafisalfiani/p3-final-project/lib/broker"
	"github.com/nafisalfiani/p3-final-project/lib/log"
	"github.com/nafisalfiani/p3-final-project/lib/parser"
	"github.com/nafisalfiani/p3-final-project/transaction-service/domain/transaction"
	"github.com/nafisalfiani/p3-final-project/transaction-service/domain/wallet"
	xenditDom "github.com/nafisalfiani/p3-final-project/transaction-service/domain/xendit"
	"github.com/xendit/xendit-go/v4"
	"gorm.io/gorm"
)

type Domains struct {
	Transaction transaction.Interface
	Wallet      wallet.Interface
	Xendit      xenditDom.Interface
}

func Init(logger log.Interface, json parser.JSONInterface, db *gorm.DB, broker broker.Interface, xnd *xendit.APIClient) *Domains {
	return &Domains{
		Transaction: transaction.Init(db),
		Wallet:      wallet.Init(db),
		Xendit:      xenditDom.Init(xnd, logger),
	}
}
