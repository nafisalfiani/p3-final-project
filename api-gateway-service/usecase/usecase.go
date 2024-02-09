package usecase

import (
	"github.com/go-playground/validator/v10"
	"github.com/nafisalfiani/p3-final-project/api-gateway-service/domain"
	accountservice "github.com/nafisalfiani/p3-final-project/api-gateway-service/usecase/account_service"
	productservice "github.com/nafisalfiani/p3-final-project/api-gateway-service/usecase/product_service"
	transactionservice "github.com/nafisalfiani/p3-final-project/api-gateway-service/usecase/transaction_service"
	"github.com/nafisalfiani/p3-final-project/lib/auth"
	"github.com/nafisalfiani/p3-final-project/lib/log"
)

type Usecases struct {
	AccountSvc     accountservice.Interface
	ProductSvc     productservice.Interface
	TransactionSvc transactionservice.Interface
}

type Config struct {
	AccountConfig     accountservice.Config
	ProductConfig     productservice.Config
	TransactionConfig transactionservice.Config
}

func Init(cfg Config, logger log.Interface, auth auth.Interface, validator *validator.Validate, dom *domain.Domains) *Usecases {
	return &Usecases{
		AccountSvc:     accountservice.Init(cfg.AccountConfig, logger),
		ProductSvc:     productservice.Init(cfg.ProductConfig, logger, auth),
		TransactionSvc: transactionservice.Init(cfg.TransactionConfig, logger),
	}
}

func (u *Usecases) CloseAllConns() {
	u.AccountSvc.Close()
	u.ProductSvc.Close()
	u.TransactionSvc.Close()
}
