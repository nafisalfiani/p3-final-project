package productservice

import (
	"context"
	"time"

	"github.com/nafisalfiani/p3-final-project/api-gateway-service/entity"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (p *productSvc) ListCategory(ctx context.Context) ([]entity.Category, error) {
	categories, err := p.category.GetCategories(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, err
	}

	var res []entity.Category
	for _, v := range categories.GetCategories() {
		res = append(res, entity.Category{
			Id:   v.GetId(),
			Name: v.GetName(),
		})
	}

	return res, nil
}
func (p *productSvc) ListRegion(ctx context.Context) ([]entity.Region, error) {
	regions, err := p.region.GetRegions(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, err
	}

	var res []entity.Region
	for _, v := range regions.GetRegions() {
		res = append(res, entity.Region{
			Id:   v.GetId(),
			Name: v.GetName(),
		})
	}

	return res, nil
}

func (p *productSvc) ListTicket(ctx context.Context, in entity.Ticket) ([]entity.Ticket, error) {
	tickets, err := p.ticket.GetTickets(ctx, toTicketProto(in))
	if err != nil {
		return nil, err
	}

	var res []entity.Ticket
	for _, v := range tickets.GetTickets() {
		res = append(res, fromTicketProto(v))
	}

	return res, nil
}

func (p *productSvc) GetTicket(ctx context.Context, in entity.Ticket) (entity.Ticket, error) {
	ticket, err := p.ticket.GetTicket(ctx, toTicketProto(in))
	if err != nil {
		return entity.Ticket{}, err
	}

	return fromTicketProto(ticket), nil
}

func (p *productSvc) RegisterTicketForSale(ctx context.Context, in entity.Ticket) (entity.Ticket, error) {
	ticket, err := p.ticket.CreateTicket(ctx, toTicketProto(in))
	if err != nil {
		return entity.Ticket{}, err
	}

	return fromTicketProto(ticket), nil
}

func (p *productSvc) UpdateTicketInfo(ctx context.Context, in entity.Ticket) (entity.Ticket, error) {
	userInfo, err := p.auth.GetUserAuthInfo(ctx)
	if err != nil {
		return entity.Ticket{}, err
	}

	in.UpdatedAt = time.Now()
	in.UpdatedBy = userInfo.User.Id
	ticket, err := p.ticket.UpdateTicket(ctx, toTicketProto(in))
	if err != nil {
		return fromTicketProto(ticket), err
	}

	return fromTicketProto(ticket), nil
}

func (p *productSvc) TakeDownTicket(ctx context.Context, in entity.Ticket) error {
	_, err := p.ticket.DeleteTicket(ctx, toTicketProto(in))
	if err != nil {
		return err
	}

	return nil
}
