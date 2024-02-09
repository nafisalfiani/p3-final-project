package transactionservice

import (
	"context"

	"github.com/nafisalfiani/p3-final-project/api-gateway-service/entity"
)

func (t *trxSvc) GetWallet(ctx context.Context, in entity.Wallet) (entity.Wallet, error) {
	walletProto, err := t.wallet.GetWallet(ctx, toWalletProto(in))
	if err != nil {
		return entity.Wallet{}, err
	}

	wallet := fromWalletProto(walletProto)
	for _, v := range walletProto.GetHistory() {
		wallet.History = append(wallet.History, entity.WalletHistory{
			Id:              v.GetId(),
			WalletId:        v.GetWalletId(),
			PreviousBalance: v.GetPreviousBalance(),
			CurrentBalance:  v.GetCurrentBalance(),
			TransactionType: v.GetTransactionType(),
			CreatedAt:       v.GetCreatedAt().AsTime(),
			CreatedBy:       v.GetCreatedBy(),
		})
	}

	return wallet, nil
}
