package rest

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nafisalfiani/p3-final-project/api-gateway-service/entity"
	"github.com/nafisalfiani/p3-final-project/lib/codes"
)

// @Summary List category
// @Description This endpoint will return a list of category
// @Tags Product
// @Security BearerAuth
// @Produce json
// @Success 200 {object} entity.HTTPResp{}
// @Failure 400 {object} entity.HTTPResp{}
// @Failure 500 {object} entity.HTTPResp{}
// @Router /api/v1/category [get]
func (r *rest) ListCategory(ctx *gin.Context) {
	res, err := r.uc.ProductSvc.ListCategory(ctx.Request.Context())
	if err != nil {
		r.httpRespError(ctx, err)
		return
	}

	r.httpRespSuccess(ctx, codes.CodeSuccess, res, nil)
}

// @Summary List region
// @Description This endpoint will return a list of region
// @Tags Product
// @Security BearerAuth
// @Produce json
// @Success 200 {object} entity.HTTPResp{}
// @Failure 400 {object} entity.HTTPResp{}
// @Failure 500 {object} entity.HTTPResp{}
// @Router /api/v1/region [get]
func (r *rest) ListRegion(ctx *gin.Context) {
	res, err := r.uc.ProductSvc.ListRegion(ctx.Request.Context())
	if err != nil {
		r.httpRespError(ctx, err)
		return
	}

	r.httpRespSuccess(ctx, codes.CodeSuccess, res, nil)
}

// @Summary List ticket
// @Description This endpoint will return a list of tickets not yet bought/sold
// @Tags Product
// @Security BearerAuth
// @Produce json
// @Success 200 {object} entity.HTTPResp{}
// @Failure 400 {object} entity.HTTPResp{}
// @Failure 500 {object} entity.HTTPResp{}
// @Router /api/v1/ticket [get]
func (r *rest) ListOpenTicket(ctx *gin.Context) {
	res, err := r.uc.ProductSvc.ListTicket(ctx.Request.Context(), entity.Ticket{
		Status: "open",
	})
	if err != nil {
		r.httpRespError(ctx, err)
		return
	}

	r.httpRespSuccess(ctx, codes.CodeSuccess, res, nil)
}

// @Summary List ticket sold by me
// @Description This endpoint will return a list of tickets sold by logged in user
// @Tags Product
// @Security BearerAuth
// @Produce json
// @Success 200 {object} entity.HTTPResp{}
// @Failure 400 {object} entity.HTTPResp{}
// @Failure 500 {object} entity.HTTPResp{}
// @Router /api/v1/ticket-sold [get]
func (r *rest) ListSoldTicketByMe(ctx *gin.Context) {
	userInfo, err := r.auth.GetUserAuthInfo(ctx.Request.Context())
	if err != nil {
		r.httpRespError(ctx, err)
		return
	}

	res, err := r.uc.ProductSvc.ListTicket(ctx.Request.Context(), entity.Ticket{
		SellerId: userInfo.User.Id,
		Status:   "sold",
	})
	if err != nil {
		r.httpRespError(ctx, err)
		return
	}

	r.httpRespSuccess(ctx, codes.CodeSuccess, res, nil)
}

// @Summary List ticket bought by me
// @Description This endpoint will return a list of tickets bought by logged in user
// @Tags Product
// @Security BearerAuth
// @Produce json
// @Success 200 {object} entity.HTTPResp{}
// @Failure 400 {object} entity.HTTPResp{}
// @Failure 500 {object} entity.HTTPResp{}
// @Router /api/v1/ticket-bought [get]
func (r *rest) ListBoughtTicketByMe(ctx *gin.Context) {
	userInfo, err := r.auth.GetUserAuthInfo(ctx.Request.Context())
	if err != nil {
		r.httpRespError(ctx, err)
		return
	}

	res, err := r.uc.ProductSvc.ListTicket(ctx.Request.Context(), entity.Ticket{
		BuyerId: userInfo.User.Id,
		Status:  "sold",
	})
	if err != nil {
		r.httpRespError(ctx, err)
		return
	}

	r.httpRespSuccess(ctx, codes.CodeSuccess, res, nil)
}

// @Summary Register ticket for sale
// @Description This endpoint will accept a request to put ticket up for sale
// @Tags Product
// @Security BearerAuth
// @Param ticket body entity.TicketCreateRequest true "ticket request"
// @Produce json
// @Success 200 {object} entity.HTTPResp{}
// @Failure 400 {object} entity.HTTPResp{}
// @Failure 500 {object} entity.HTTPResp{}
// @Router /api/v1/ticket [post]
func (r *rest) RegisterTicketForSale(ctx *gin.Context) {
	var req entity.TicketCreateRequest
	err := r.Bind(ctx, &req)
	if err != nil {
		r.httpRespError(ctx, err)
		return
	}

	userInfo, err := r.auth.GetUserAuthInfo(ctx.Request.Context())
	if err != nil {
		r.httpRespError(ctx, err)
		return
	}

	ticketReq := req.ToTicket()
	ticketReq.Status = "open"
	ticketReq.SellerId = userInfo.User.Id
	ticketReq.CreatedAt = time.Now()
	ticketReq.CreatedBy = userInfo.User.Id

	res, err := r.uc.ProductSvc.RegisterTicketForSale(ctx.Request.Context(), ticketReq)
	if err != nil {
		r.httpRespError(ctx, err)
		return
	}

	r.httpRespSuccess(ctx, codes.CodeSuccess, res, nil)
}

// @Summary Update ticket info
// @Description This endpoint will accept request to update specific ticket
// @Tags Product
// @Security BearerAuth
// @Produce json
// @Success 200 {object} entity.HTTPResp{}
// @Failure 400 {object} entity.HTTPResp{}
// @Failure 500 {object} entity.HTTPResp{}
// @Router /api/v1/ticket/:id [put]
func (r *rest) UpdateTicketInfo(ctx *gin.Context) {
	var req entity.TicketUpdateRequest
	if err := r.BindParams(ctx, &req); err != nil {
		r.httpRespError(ctx, err)
		return
	}

	if err := r.Bind(ctx, &req); err != nil {
		r.httpRespError(ctx, err)
		return
	}

	userInfo, err := r.auth.GetUserAuthInfo(ctx.Request.Context())
	if err != nil {
		r.httpRespError(ctx, err)
		return
	}

	ticketReq := req.ToTicket()
	ticketReq.UpdatedAt = time.Now()
	ticketReq.UpdatedBy = userInfo.User.Id

	res, err := r.uc.ProductSvc.UpdateTicketInfo(ctx.Request.Context(), ticketReq)
	if err != nil {
		r.httpRespError(ctx, err)
		return
	}

	r.httpRespSuccess(ctx, codes.CodeSuccess, res, nil)
}

// @Summary Delete ticket
// @Description This endpoint will accept request to take down ticket
// @Tags Product
// @Security BearerAuth
// @Produce json
// @Success 200 {object} entity.HTTPResp{}
// @Failure 400 {object} entity.HTTPResp{}
// @Failure 500 {object} entity.HTTPResp{}
// @Router /api/v1/ticket/:id [delete]
func (r *rest) TakeDownTicket(ctx *gin.Context) {
	var req entity.TicketGetRequest
	if err := r.BindParams(ctx, &req); err != nil {
		r.httpRespError(ctx, err)
		return
	}

	if err := r.uc.ProductSvc.TakeDownTicket(ctx.Request.Context(), entity.Ticket{
		Id: req.Id,
	}); err != nil {
		r.httpRespError(ctx, err)
		return
	}

	r.httpRespSuccess(ctx, codes.CodeSuccess, nil, nil)
}
