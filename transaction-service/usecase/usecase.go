package usecase

import (
	"github.com/nafisalfiani/p3-final-project/lib/log"
	"github.com/nafisalfiani/p3-final-project/transaction-service/domain"
	"github.com/nafisalfiani/p3-final-project/transaction-service/usecase/transaction"
	"github.com/nafisalfiani/p3-final-project/transaction-service/usecase/wallet"
)

type Usecases struct {
	Transaction transaction.Interface
	Wallet      wallet.Interface
}

func Init(logger log.Interface, dom *domain.Domains) *Usecases {
	return &Usecases{
		Transaction: transaction.Init(dom.Transaction),
		Wallet:      wallet.Init(dom.Wallet),
	}
}
