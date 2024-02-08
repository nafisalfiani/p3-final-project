package wallet

import (
	"context"

	walletDom "github.com/nafisalfiani/p3-final-project/transaction-service/domain/wallet"
	"github.com/nafisalfiani/p3-final-project/transaction-service/entity"
)

type Interface interface {
	List(ctx context.Context) ([]entity.Wallet, error)
	Get(ctx context.Context, filter entity.Wallet) (entity.Wallet, error)
	Create(ctx context.Context, wallet entity.Wallet) (entity.Wallet, error)
	Update(ctx context.Context, wallet entity.Wallet) (entity.Wallet, error)
	Delete(ctx context.Context, walletId string) error
}

type wallet struct {
	wallet walletDom.Interface
}

func Init(prd walletDom.Interface) Interface {
	return &wallet{
		wallet: prd,
	}
}

func (c *wallet) List(ctx context.Context) ([]entity.Wallet, error) {
	return c.wallet.List(ctx)
}

func (c *wallet) Get(ctx context.Context, filter entity.Wallet) (entity.Wallet, error) {
	return c.wallet.Get(ctx, filter)
}

func (c *wallet) Create(ctx context.Context, wallet entity.Wallet) (entity.Wallet, error) {
	return c.wallet.Create(ctx, wallet)
}

func (c *wallet) Update(ctx context.Context, wallet entity.Wallet) (entity.Wallet, error) {
	return c.wallet.Update(ctx, wallet)
}

func (c *wallet) Delete(ctx context.Context, walletId string) error {
	return c.wallet.Delete(ctx, walletId)
}
