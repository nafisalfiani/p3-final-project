package wallet

import (
	"github.com/nafisalfiani/p3-final-project/transaction-service/entity"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func toProto(wallet entity.Wallet) *Wallet {
	var history []*History
	for i := range wallet.History {
		history = append(history, &History{
			Id:              wallet.History[i].Id,
			WalletId:        wallet.History[i].WalletId,
			PreviousBalance: wallet.History[i].PreviousBalance,
			CurrentBalance:  wallet.History[i].CurrentBalance,
			TransactionType: wallet.History[i].TransactionType,
			CreatedAt:       timestamppb.New(wallet.History[i].CreatedAt),
			CreatedBy:       wallet.History[i].CreatedBy,
		})
	}

	return &Wallet{
		Id:        wallet.Id,
		UserId:    wallet.UserId,
		Balance:   wallet.Balance,
		History:   history,
		CreatedAt: timestamppb.New(wallet.CreatedAt),
		CreatedBy: wallet.CreatedBy,
		UpdatedAt: timestamppb.New(wallet.UpdatedAt),
		UpdatedBy: wallet.UpdatedBy,
	}
}

func fromProto(in *Wallet) entity.Wallet {
	var history []entity.WalletHistory
	for _, v := range in.GetHistory() {
		history = append(history, entity.WalletHistory{
			Id:              v.Id,
			WalletId:        v.WalletId,
			PreviousBalance: v.PreviousBalance,
			CurrentBalance:  v.CurrentBalance,
			TransactionType: v.TransactionType,
			CreatedAt:       v.CreatedAt.AsTime(),
			CreatedBy:       v.CreatedBy,
		})
	}

	return entity.Wallet{
		Id:        in.GetId(),
		UserId:    in.GetUserId(),
		Balance:   in.GetBalance(),
		History:   history,
		CreatedAt: in.GetCreatedAt().AsTime(),
		CreatedBy: in.GetCreatedBy(),
		UpdatedAt: in.GetUpdatedAt().AsTime(),
		UpdatedBy: in.GetUpdatedBy(),
	}
}
