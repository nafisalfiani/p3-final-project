package transaction

import (
	"github.com/nafisalfiani/p3-final-project/transaction-service/entity"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func toProto(trx entity.Transaction) *Transaction {
	return &Transaction{
		Id:          trx.Id,
		TicketId:    trx.TicketId,
		TicketCount: trx.TicketCount,
		BuyerId:     trx.BuyerId,
		SellerId:    trx.SellerId,
		Amount:      trx.Amount,
		CreatedAt:   timestamppb.New(trx.CreatedAt),
		CreatedBy:   trx.CreatedBy,
		UpdatedAt:   timestamppb.New(trx.UpdatedAt),
		UpdatedBy:   trx.UpdatedBy,
	}
}

func fromProto(in *Transaction) entity.Transaction {
	return entity.Transaction{
		Id:          in.GetId(),
		TicketId:    in.GetTicketId(),
		TicketCount: in.GetTicketCount(),
		BuyerId:     in.GetBuyerId(),
		SellerId:    in.GetSellerId(),
		Amount:      in.GetAmount(),
		CreatedAt:   in.GetCreatedAt().AsTime(),
		CreatedBy:   in.GetCreatedBy(),
		UpdatedAt:   in.GetUpdatedAt().AsTime(),
		UpdatedBy:   in.GetUpdatedBy(),
	}
}
