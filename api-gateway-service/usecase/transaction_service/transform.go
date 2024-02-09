package transactionservice

import (
	"github.com/nafisalfiani/p3-final-project/api-gateway-service/entity"
	"github.com/nafisalfiani/p3-final-project/transaction-service/handler/grpc/transaction"
	"github.com/nafisalfiani/p3-final-project/transaction-service/handler/grpc/wallet"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func toTransactionProto(in entity.Transaction) *transaction.Transaction {
	return &transaction.Transaction{
		Id:        in.Id,
		TicketId:  in.TicketId,
		BuyerId:   in.BuyerId,
		SellerId:  in.SellerId,
		Amount:    in.Amount,
		CreatedAt: timestamppb.New(in.CreatedAt),
		CreatedBy: in.CreatedBy,
		UpdatedAt: timestamppb.New(in.UpdatedAt),
		UpdatedBy: in.UpdatedBy,
	}
}

func fromTransactionProto(v *transaction.Transaction) entity.Transaction {
	return entity.Transaction{
		Id:        v.GetId(),
		TicketId:  v.GetTicketId(),
		BuyerId:   v.GetBuyerId(),
		SellerId:  v.GetSellerId(),
		Amount:    v.GetAmount(),
		CreatedAt: v.GetCreatedAt().AsTime(),
		CreatedBy: v.GetCreatedBy(),
		UpdatedAt: v.GetUpdatedAt().AsTime(),
		UpdatedBy: v.GetUpdatedBy(),
	}
}

func toWalletProto(in entity.Wallet) *wallet.Wallet {
	return &wallet.Wallet{
		Id:        in.Id,
		UserId:    in.UserId,
		Balance:   in.Balance,
		CreatedAt: timestamppb.New(in.CreatedAt),
		CreatedBy: in.CreatedBy,
		UpdatedAt: timestamppb.New(in.UpdatedAt),
		UpdatedBy: in.UpdatedBy,
	}
}

func fromWalletProto(v *wallet.Wallet) entity.Wallet {
	return entity.Wallet{
		Id:        v.GetId(),
		UserId:    v.GetUserId(),
		Balance:   v.GetBalance(),
		CreatedAt: v.GetCreatedAt().AsTime(),
		CreatedBy: v.GetCreatedBy(),
		UpdatedAt: v.GetUpdatedAt().AsTime(),
		UpdatedBy: v.GetUpdatedBy(),
	}
}
