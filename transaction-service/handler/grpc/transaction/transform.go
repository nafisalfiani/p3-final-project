package transaction

import (
	"github.com/google/uuid"
	"github.com/nafisalfiani/p3-final-project/transaction-service/entity"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func toProto(trx entity.Transaction) *Transaction {
	return &Transaction{
		Id:        trx.Id.String(),
		TicketId:  trx.TicketId,
		BuyerId:   trx.BuyerId,
		SellerId:  trx.SellerId,
		Amount:    trx.Amount,
		CreatedAt: timestamppb.New(trx.CreatedAt),
		CreatedBy: trx.CreatedBy,
		UpdatedAt: timestamppb.New(trx.UpdatedAt),
		UpdatedBy: trx.UpdatedBy,
	}
}

func fromProto(in *Transaction) entity.Transaction {
	uid, _ := uuid.Parse(in.GetId())
	return entity.Transaction{
		Id:        uid,
		TicketId:  in.GetTicketId(),
		BuyerId:   in.GetBuyerId(),
		SellerId:  in.GetSellerId(),
		Amount:    in.GetAmount(),
		CreatedAt: in.GetCreatedAt().AsTime(),
		CreatedBy: in.GetCreatedBy(),
		UpdatedAt: in.GetUpdatedAt().AsTime(),
		UpdatedBy: in.GetUpdatedBy(),
	}
}
