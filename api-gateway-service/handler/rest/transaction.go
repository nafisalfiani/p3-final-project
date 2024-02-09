package rest

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nafisalfiani/p3-final-project/api-gateway-service/entity"
	"github.com/nafisalfiani/p3-final-project/lib/codes"
)

// @Summary List transaction
// @Description This endpoint will return a list of tickets not yet bought/sold
// @Tags Transaction
// @Security BearerAuth
// @Produce json
// @Success 200 {object} entity.HTTPResp{}
// @Failure 400 {object} entity.HTTPResp{}
// @Failure 500 {object} entity.HTTPResp{}
// @Router /api/v1/transaction [get]
func (r *rest) ListTransaction(ctx *gin.Context) {
	userInfo, err := r.auth.GetUserAuthInfo(ctx.Request.Context())
	if err != nil {
		r.httpRespError(ctx, err)
		return
	}

	res, err := r.uc.TransactionSvc.ListTransaction(ctx.Request.Context(), entity.Transaction{
		BuyerId: userInfo.User.Id,
	})
	if err != nil {
		r.httpRespError(ctx, err)
		return
	}

	r.httpRespSuccess(ctx, codes.CodeSuccess, res, nil)
}

// @Summary Create transaction
// @Description This endpoint will accept a request to create transaction
// @Tags Transaction
// @Security BearerAuth
// @Param trx body entity.TransactionCreateRequest true "transaction"
// @Produce json
// @Success 200 {object} entity.HTTPResp{}
// @Failure 400 {object} entity.HTTPResp{}
// @Failure 500 {object} entity.HTTPResp{}
// @Router /api/v1/transaction [post]
func (r *rest) CreateTransaction(ctx *gin.Context) {
	var req entity.TransactionCreateRequest
	if err := r.Bind(ctx, &req); err != nil {
		r.httpRespError(ctx, err)
		return
	}

	userInfo, err := r.auth.GetUserAuthInfo(ctx.Request.Context())
	if err != nil {
		r.httpRespError(ctx, err)
		return
	}

	ticket, err := r.uc.ProductSvc.GetTicket(ctx.Request.Context(), entity.Ticket{
		Id: req.TicketId,
	})
	if err != nil {
		r.httpRespError(ctx, err)
	}

	res, err := r.uc.TransactionSvc.CreateTransaction(ctx.Request.Context(), entity.Transaction{
		TicketId:  ticket.Id,
		Amount:    ticket.SellingPrice,
		BuyerId:   userInfo.User.Id,
		Status:    "waiting for payment",
		CreatedAt: time.Now(),
		CreatedBy: userInfo.User.Id,
	})
	if err != nil {
		r.httpRespError(ctx, err)
		return
	}

	r.httpRespSuccess(ctx, codes.CodeSuccess, res, nil)
}

// @Summary Update transaction
// @Description This endpoint will accept a request to update transaction
// @Tags Transaction
// @Security BearerAuth
// @Produce json
// @Success 200 {object} entity.HTTPResp{}
// @Failure 400 {object} entity.HTTPResp{}
// @Failure 500 {object} entity.HTTPResp{}
// @Router /api/v1/transaction [put]
func (r *rest) UpdateTransaction(ctx *gin.Context) {
	var req entity.TransactionCreateRequest
	if err := r.Bind(ctx, &req); err != nil {
		r.httpRespError(ctx, err)
		return
	}

	userInfo, err := r.auth.GetUserAuthInfo(ctx.Request.Context())
	if err != nil {
		r.httpRespError(ctx, err)
		return
	}

	res, err := r.uc.TransactionSvc.CreateTransaction(ctx.Request.Context(), entity.Transaction{
		TicketId:  req.TicketId,
		BuyerId:   userInfo.User.Id,
		Status:    "waiting for payment",
		UpdatedAt: time.Now(),
		UpdatedBy: userInfo.User.Id,
	})
	if err != nil {
		r.httpRespError(ctx, err)
		return
	}

	r.httpRespSuccess(ctx, codes.CodeSuccess, res, nil)
}

// @Summary Update payment status
// @Description Update payment status if applicable
// @Tags Transaction
// @Accept json
// @Produce json
// @Param payment body entity.XenditCheckPayment true "payment request"
// @Success 200 {object} entity.HttpResp
// @Failure 400 {object} entity.HttpResp
// @Failure 500 {object} entity.HttpResp
// @Router /webhooks/payment [post]
func (r *rest) UpdatePaymentStatus(ctx *gin.Context) {
	req := entity.XenditWebhookBody{}
	if err := r.Bind(ctx, &req); err != nil {
		r.httpRespError(ctx, err)
		return
	}

	r.log.Info(ctx, ctx.Request.Header.Get("x-callback-token"))
	r.log.Info(ctx, ctx.Request.Header.Get("webhook-id"))

	paymentType, paymentId, err := req.GetPaymentId()
	if err != nil {
		r.httpRespError(ctx, err)
		return
	}

	trx := entity.Transaction{
		Id:            paymentId,
		Type:          paymentType,
		Status:        req.Status,
		PaymentMethod: req.PaymentMethod,
	}

	newPayment, err := r.uc.TransactionSvc.UpdateTransaction(ctx, trx)
	if err != nil {
		r.httpRespError(ctx, err)
		return
	}

	r.httpRespSuccess(ctx, http.StatusOK, newPayment, nil)
}

// @Summary Wallet
// @Description This endpoint will return user wallet
// @Tags Transaction
// @Security BearerAuth
// @Produce json
// @Success 200 {object} entity.HTTPResp{}
// @Failure 400 {object} entity.HTTPResp{}
// @Failure 500 {object} entity.HTTPResp{}
// @Router /api/v1/wallet [get]
func (r *rest) GetWallet(ctx *gin.Context) {
	userInfo, err := r.auth.GetUserAuthInfo(ctx.Request.Context())
	if err != nil {
		r.httpRespError(ctx, err)
		return
	}

	res, err := r.uc.TransactionSvc.GetWallet(ctx.Request.Context(), entity.Wallet{
		UserId: userInfo.User.Id,
	})
	if err != nil {
		r.httpRespError(ctx, err)
		return
	}

	r.httpRespSuccess(ctx, codes.CodeSuccess, res, nil)
}
