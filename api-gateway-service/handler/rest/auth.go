package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/nafisalfiani/p3-final-project/api-gateway-service/entity"
	"github.com/nafisalfiani/p3-final-project/lib/codes"
)

// @Summary Register new user
// @Description This endpoint will register new user as member
// @Tags Auth
// @Param register_request body entity.RegisterRequest true "Input Username And Password"
// @Produce json
// @Success 200 {object} entity.HTTPResp{}
// @Failure 400 {object} entity.HTTPResp{}
// @Failure 500 {object} entity.HTTPResp{}
// @Router /auth/v1/register [post]
func (r *rest) RegisterUser(ctx *gin.Context) {
	var regReq entity.RegisterRequest
	err := r.Bind(ctx, &regReq)
	if err != nil {
		r.httpRespError(ctx, err)
		return
	}

	res, err := r.uc.AccountSvc.Register(ctx.Request.Context(), regReq)
	if err != nil {
		r.httpRespError(ctx, err)
		return
	}

	r.httpRespSuccess(ctx, codes.CodeSuccess, res, nil)
}

// @Summary Verify user email
// @Description This endpoint will mark user email as verified
// @Tags Auth
// @Param id path string true "user id"
// @Produce json
// @Success 200 {object} entity.HTTPResp{}
// @Failure 400 {object} entity.HTTPResp{}
// @Failure 500 {object} entity.HTTPResp{}
// @Router /auth/v1/verify-email/{id} [post]
func (r *rest) VerifyEmail(ctx *gin.Context) {
	var regReq entity.VerifyEmailRequest
	err := r.BindParams(ctx, &regReq)
	if err != nil {
		r.httpRespError(ctx, err)
		return
	}

	res, err := r.uc.AccountSvc.VerifyEmail(ctx.Request.Context(), regReq)
	if err != nil {
		r.httpRespError(ctx, err)
		return
	}

	r.httpRespSuccess(ctx, codes.CodeSuccess, res, nil)
}

// @Summary Sign In With Password
// @Description This endpoint will sign in user with username and password
// @Tags Auth
// @Param login_request body entity.LoginRequest true "Input Username And Password"
// @Produce json
// @Success 200 {object} entity.HTTPResp{}
// @Failure 400 {object} entity.HTTPResp{}
// @Failure 500 {object} entity.HTTPResp{}
// @Router /auth/v1/login [post]
func (r *rest) Login(ctx *gin.Context) {
	var loginReq entity.LoginRequest
	err := r.Bind(ctx, &loginReq)
	if err != nil {
		r.httpRespError(ctx, err)
		return
	}

	res, err := r.uc.AccountSvc.SignIn(ctx.Request.Context(), loginReq)
	if err != nil {
		r.httpRespError(ctx, err)
		return
	}

	r.httpRespSuccess(ctx, codes.CodeSuccess, res, nil)
}
