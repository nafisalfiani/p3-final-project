package ticket

import (
	context "context"

	"github.com/nafisalfiani/p3-final-project/lib/log"
	"github.com/nafisalfiani/p3-final-project/product-service/domain/ticket"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type grpcTicket struct {
	log    log.Interface
	ticket ticket.Interface
}

func Init(log log.Interface, ticket ticket.Interface) TicketServiceServer {
	return &grpcTicket{
		log:    log,
		ticket: ticket,
	}
}

func (t *grpcTicket) mustEmbedUnimplementedTicketServiceServer() {}

func (t *grpcTicket) GetTicket(ctx context.Context, in *Ticket) (*Ticket, error) {
	ticket, err := t.ticket.Get(ctx, fromProto(in))
	if err != nil {
		return nil, err
	}

	return toProto(ticket), nil
}

func (t *grpcTicket) CreateTicket(ctx context.Context, in *Ticket) (*Ticket, error) {
	ticket, err := t.ticket.Create(ctx, fromProto(in))
	if err != nil {
		return nil, err
	}

	return toProto(ticket), nil
}

func (t *grpcTicket) UpdateTicket(ctx context.Context, in *Ticket) (*Ticket, error) {
	ticket, err := t.ticket.Update(ctx, fromProto(in))
	if err != nil {
		return nil, err
	}

	return toProto(ticket), nil
}

func (t *grpcTicket) DeleteTicket(ctx context.Context, in *Ticket) (*emptypb.Empty, error) {
	if err := t.ticket.Delete(ctx, fromProto(in)); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (t *grpcTicket) GetTickets(ctx context.Context, in *emptypb.Empty) (*TicketList, error) {
	tickets, err := t.ticket.List(ctx)
	if err != nil {
		return nil, err
	}

	res := &TicketList{}
	for i := range tickets {
		res.Tickets = append(res.Tickets, toProto(tickets[i]))
	}

	return res, nil
}
