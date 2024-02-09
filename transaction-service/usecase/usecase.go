package usecase

import (
	"github.com/nafisalfiani/p3-final-project/lib/broker"
	"github.com/nafisalfiani/p3-final-project/lib/log"
	"github.com/nafisalfiani/p3-final-project/transaction-service/domain"
	"github.com/nafisalfiani/p3-final-project/transaction-service/usecase/transaction"
	"github.com/nafisalfiani/p3-final-project/transaction-service/usecase/wallet"
)

type Usecases struct {
	Transaction transaction.Interface
	Wallet      wallet.Interface
}

func Init(logger log.Interface, dom *domain.Domains, broker broker.Interface) *Usecases {
	return &Usecases{
		Transaction: transaction.Init(logger, dom.Transaction, dom.Xendit, broker),
		Wallet:      wallet.Init(dom.Wallet),
	}
}
