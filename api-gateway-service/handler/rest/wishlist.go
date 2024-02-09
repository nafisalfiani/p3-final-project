package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/nafisalfiani/p3-final-project/api-gateway-service/entity"
	"github.com/nafisalfiani/p3-final-project/lib/codes"
)

// @Summary List wishlist
// @Description This endpoint will return a list of wishlist
// @Tags Wishlist
// @Security BearerAuth
// @Produce json
// @Success 200 {object} entity.HTTPResp{}
// @Failure 400 {object} entity.HTTPResp{}
// @Failure 500 {object} entity.HTTPResp{}
// @Router /api/v1/wishlist [get]
func (r *rest) ListWishlist(ctx *gin.Context) {
	res, err := r.uc.ProductSvc.ListWishlist(ctx.Request.Context(), entity.Wishlist{})
	if err != nil {
		r.httpRespError(ctx, err)
		return
	}

	r.httpRespSuccess(ctx, codes.CodeSuccess, res, nil)
}

// @Summary Get wishlist subscriber
// @Description This endpoint will return a list of wishlist subscriber
// @Tags Wishlist
// @Security BearerAuth
// @Produce json
// @Success 200 {object} entity.HTTPResp{}
// @Failure 400 {object} entity.HTTPResp{}
// @Failure 500 {object} entity.HTTPResp{}
// @Router /api/v1/wishlist/:id [get]
func (r *rest) GetWishlistSubscriber(ctx *gin.Context) {
	var req entity.WishlistGetRequest
	if err := r.BindParams(ctx, &req); err != nil {
		r.httpRespError(ctx, err)
	}

	res, err := r.uc.ProductSvc.GetWishlistSubscriber(ctx.Request.Context(), entity.Wishlist{
		Id: req.Id,
	})
	if err != nil {
		r.httpRespError(ctx, err)
		return
	}

	r.httpRespSuccess(ctx, codes.CodeSuccess, res, nil)
}

// @Summary Subscribe to wishlist
// @Description This endpoint will accept a request to subscribe to wishlist
// @Tags Wishlist
// @Security BearerAuth
// @Produce json
// @Success 200 {object} entity.HTTPResp{}
// @Failure 400 {object} entity.HTTPResp{}
// @Failure 500 {object} entity.HTTPResp{}
// @Router /api/v1/wishlist [post]
func (r *rest) SubscribeToWishlist(ctx *gin.Context) {
	var req entity.WishlistSubscribeRequest
	if err := r.BindParams(ctx, &req); err != nil {
		r.httpRespError(ctx, err)
	}

	userInfo, err := r.auth.GetUserAuthInfo(ctx.Request.Context())
	if err != nil {
		r.httpRespError(ctx, err)
	}

	res, err := r.uc.ProductSvc.Subscribe(ctx.Request.Context(), entity.Wishlist{
		CategoryName:    req.CategoryName,
		RegionName:      req.RegionName,
		SubscribedUsers: []string{userInfo.User.Id},
	})
	if err != nil {
		r.httpRespError(ctx, err)
		return
	}

	r.httpRespSuccess(ctx, codes.CodeSuccess, res, nil)
}

// @Summary Unsubscribe from wishlist
// @Description This endpoint will accept a request to unsubscribe from wishlist
// @Tags Wishlist
// @Security BearerAuth
// @Produce json
// @Success 200 {object} entity.HTTPResp{}
// @Failure 400 {object} entity.HTTPResp{}
// @Failure 500 {object} entity.HTTPResp{}
// @Router /api/v1/wishlist [delete]
func (r *rest) UnsubscribeFromWishlist(ctx *gin.Context) {
	var req entity.WishlistSubscribeRequest
	if err := r.BindParams(ctx, &req); err != nil {
		r.httpRespError(ctx, err)
	}

	userInfo, err := r.auth.GetUserAuthInfo(ctx.Request.Context())
	if err != nil {
		r.httpRespError(ctx, err)
	}

	if err := r.uc.ProductSvc.Unsubscribe(ctx.Request.Context(), entity.Wishlist{
		CategoryName:    req.CategoryName,
		RegionName:      req.RegionName,
		SubscribedUsers: []string{userInfo.User.Id},
	}); err != nil {
		r.httpRespError(ctx, err)
		return
	}

	r.httpRespSuccess(ctx, codes.CodeSuccess, nil, nil)
}
