package wallet

import (
	"github.com/google/uuid"
	"github.com/nafisalfiani/p3-final-project/transaction-service/entity"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func toProto(wallet entity.Wallet) *Wallet {
	var history []*History
	for i := range wallet.History {
		history = append(history, &History{
			Id:              wallet.History[i].Id.String(),
			WalletId:        wallet.History[i].WalletId,
			PreviousBalance: wallet.History[i].PreviousBalance,
			CurrentBalance:  wallet.History[i].CurrentBalance,
			TransactionType: wallet.History[i].TransactionType,
			CreatedAt:       timestamppb.New(wallet.History[i].CreatedAt),
			CreatedBy:       wallet.History[i].CreatedBy,
		})
	}

	return &Wallet{
		Id:        wallet.Id.String(),
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
		uid, _ := uuid.Parse(in.GetId())
		history = append(history, entity.WalletHistory{
			Id:              uid,
			WalletId:        v.WalletId,
			PreviousBalance: v.PreviousBalance,
			CurrentBalance:  v.CurrentBalance,
			TransactionType: v.TransactionType,
			CreatedAt:       v.CreatedAt.AsTime(),
			CreatedBy:       v.CreatedBy,
		})
	}

	uid, _ := uuid.Parse(in.GetId())
	return entity.Wallet{
		Id:        uid,
		UserId:    in.GetUserId(),
		Balance:   in.GetBalance(),
		History:   history,
		CreatedAt: in.GetCreatedAt().AsTime(),
		CreatedBy: in.GetCreatedBy(),
		UpdatedAt: in.GetUpdatedAt().AsTime(),
		UpdatedBy: in.GetUpdatedBy(),
	}
}
