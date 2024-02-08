package rest

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/google/uuid"
	"github.com/nafisalfiani/p3-final-project/api-gateway-service/entity"
	"github.com/nafisalfiani/p3-final-project/lib/appcontext"
	"github.com/nafisalfiani/p3-final-project/lib/auth"
	"github.com/nafisalfiani/p3-final-project/lib/checker"
	"github.com/nafisalfiani/p3-final-project/lib/codes"
	"github.com/nafisalfiani/p3-final-project/lib/errors"
	"github.com/nafisalfiani/p3-final-project/lib/header"
)

func (r *rest) VerifyUser(ctx *gin.Context) {
	userAuth, err := r.verifyUserAuth(ctx)
	if err != nil {
		r.httpRespError(ctx, err)
		return
	}

	c := ctx.Request.Context()
	c = r.auth.SetUserAuthInfo(c,
		auth.User{
			Id:       userAuth.User.Id,
			Name:     userAuth.User.Name,
			Email:    userAuth.User.Email,
			Username: userAuth.User.Username,
			Role:     userAuth.User.Role,
			Scopes:   userAuth.User.Scopes,
		}, &auth.Token{
			TokenType:   header.AuthorizationBearer,
			AccessToken: userAuth.Token.AccessToken,
		})

	c = appcontext.SetUserId(c, userAuth.User.Id)
	ctx.Request = ctx.Request.WithContext(c)

	ctx.Next()
}

func (r *rest) verifyUserAuth(ctx *gin.Context) (auth.UserAuthInfo, error) {
	token := ctx.Request.Header.Get(header.KeyAuthorization)
	if token == "" {
		return auth.UserAuthInfo{}, errors.NewWithCode(codes.CodeUnauthorized, "empty token")
	}

	var tokenId string
	_, err := fmt.Sscanf(token, "Bearer %v", &tokenId)
	if err != nil {
		return auth.UserAuthInfo{}, errors.NewWithCode(codes.CodeUnauthorized, "invalid format: %s with err:%v", token, err)
	}

	// verify token with secret
	user, err := r.auth.VerifyToken(ctx.Request.Context(), tokenId)
	if err != nil {
		return auth.UserAuthInfo{}, err
	}

	userAuth := auth.UserAuthInfo{
		User: user,
		Token: &auth.Token{
			TokenType:   header.AuthorizationBearer,
			AccessToken: tokenId,
		},
	}

	return userAuth, nil
}

func (r *rest) AuthorizeScope(actionCode string, f gin.HandlerFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user, err := r.auth.GetUserAuthInfo(ctx.Request.Context())
		if err != nil {
			r.httpRespError(ctx, errors.NewWithCode(codes.CodeAuthFailure, "failed to get user auth info"))
			return
		}

		if ok := checker.ArrayContains(user.User.Scopes, actionCode); ok {
			ctx.Next()
			f(ctx)
			return
		}

		r.httpRespError(ctx, errors.NewWithCode(codes.CodeUnauthorized, "User doesn't have access"))
	}
}

func (r *rest) BodyLogger(ctx *gin.Context) {
	if r.conf.LogRequest {
		r.log.Info(ctx.Request.Context(),
			fmt.Sprintf(infoRequest, ctx.Request.RequestURI, ctx.Request.Method))
	}

	ctx.Next()
	if r.conf.LogResponse {
		if ctx.Writer.Status() < 300 {
			r.log.Info(ctx.Request.Context(),
				fmt.Sprintf(infoResponse, ctx.Request.RequestURI, ctx.Request.Method, ctx.Writer.Status()))
		} else {
			r.log.Error(ctx.Request.Context(),
				fmt.Sprintf(infoResponse, ctx.Request.RequestURI, ctx.Request.Method, ctx.Writer.Status()))
		}
	}
}

// timeout middleware wraps the request context with a timeout
func (r *rest) SetTimeout(ctx *gin.Context) {
	// wrap the request context with a timeout
	c, cancel := context.WithTimeout(ctx.Request.Context(), r.conf.Timeout)

	// cancel to clear resources after finished
	defer cancel()

	c = appcontext.SetRequestStartTime(c, time.Now())

	// replace request with context wrapped request
	ctx.Request = ctx.Request.WithContext(c)
	ctx.Next()

}

func (r *rest) addFieldsToContext(ctx *gin.Context) {
	reqid := ctx.GetHeader(header.KeyRequestID)
	if reqid == "" {
		reqid = uuid.New().String()
	}

	c := ctx.Request.Context()
	c = appcontext.SetRequestId(c, reqid)
	c = appcontext.SetUserAgent(c, ctx.Request.Header.Get(header.KeyUserAgent))
	c = appcontext.SetServiceVersion(c, r.conf.Meta.Version)
	c = appcontext.SetDeviceType(c, ctx.Request.Header.Get(header.KeyDeviceType))
	c = appcontext.SetCacheControl(c, ctx.Request.Header.Get(header.KeyCacheControl))
	ctx.Request = ctx.Request.WithContext(c)
	ctx.Next()
}

func (r *rest) httpRespError(ctx *gin.Context, err error) {
	c := ctx.Request.Context()

	if errors.Is(c.Err(), context.DeadlineExceeded) {
		err = errors.NewWithCode(codes.CodeContextDeadlineExceeded, "Context Deadline Exceeded")
	}

	httpStatus, displayError := errors.Compile(err)
	errResp := &entity.HTTPResp{
		Message: entity.HTTPMessage{
			Title: displayError.Title,
			Body:  displayError.Body,
		},
		Meta: entity.Meta{
			Path:       r.conf.Meta.Host + ctx.Request.URL.String(),
			StatusCode: httpStatus,
			Status:     http.StatusText(httpStatus),
			Message:    fmt.Sprintf("%s %s [%d] %s", ctx.Request.Method, ctx.Request.URL.RequestURI(), httpStatus, http.StatusText(httpStatus)),
			Error: &entity.MetaError{
				Code:    int(displayError.Code),
				Message: err.Error(),
			},
			Timestamp: time.Now().Format(time.RFC3339),
			RequestID: appcontext.GetRequestId(c),
		},
	}

	r.log.Error(c, err)

	c = appcontext.SetAppResponseCode(c, displayError.Code)
	c = appcontext.SetAppErrorMessage(c, fmt.Sprintf("%s - %s", displayError.Title, displayError.Body))
	c = appcontext.SetResponseHttpCode(c, httpStatus)
	ctx.Request = ctx.Request.WithContext(c)

	ctx.Header(header.KeyRequestID, appcontext.GetRequestId(c))
	ctx.AbortWithStatusJSON(httpStatus, errResp)
}

func (r *rest) httpRespSuccess(ctx *gin.Context, code codes.Code, data interface{}, p *entity.Pagination) {
	successApp := codes.Compile(code)
	c := ctx.Request.Context()
	meta := entity.Meta{
		Path:       r.conf.Meta.Host + ctx.Request.URL.String(),
		StatusCode: successApp.StatusCode,
		Status:     http.StatusText(successApp.StatusCode),
		Message:    fmt.Sprintf("%s %s [%d] %s", ctx.Request.Method, ctx.Request.URL.RequestURI(), successApp.StatusCode, http.StatusText(successApp.StatusCode)),
		Timestamp:  time.Now().Format(time.RFC3339),
		RequestID:  appcontext.GetRequestId(c),
	}

	resp := &entity.HTTPResp{
		Message: entity.HTTPMessage{
			Title: successApp.Title,
			Body:  successApp.Body,
		},
		Meta:       meta,
		Data:       data,
		Pagination: p,
	}

	reqstart := appcontext.GetRequestStartTime(c)
	if !time.Time.IsZero(reqstart) {
		resp.Meta.TimeElapsed = fmt.Sprintf("%dms", int64(time.Since(reqstart)/time.Millisecond))
	}

	raw, err := r.json.Marshal(&resp)
	if err != nil {
		r.httpRespError(ctx, errors.NewWithCode(codes.CodeInternalServerError, err.Error()))
		return
	}

	c = appcontext.SetAppResponseCode(c, code)
	c = appcontext.SetResponseHttpCode(c, successApp.StatusCode)
	ctx.Request = ctx.Request.WithContext(c)

	ctx.Header(header.KeyRequestID, appcontext.GetRequestId(c))
	ctx.Data(successApp.StatusCode, header.ContentTypeJSON, raw)
}

// Bind request body to struct using tag 'json'
func (r *rest) Bind(ctx *gin.Context, obj interface{}) error {
	err := ctx.ShouldBindWith(obj, binding.Default(ctx.Request.Method, ctx.ContentType()))
	if err != nil {
		return errors.NewWithCode(codes.CodeBadRequest, err.Error())
	}

	return nil
}

// Bind all query params to struct using tag 'form'
func (r *rest) BindQuery(ctx *gin.Context, obj interface{}) error {
	err := ctx.ShouldBindWith(obj, binding.Query)
	if err != nil {
		return errors.NewWithCode(codes.CodeBadRequest, err.Error())
	}

	return nil
}

// Bind uri params to struct using tag 'uri'
func (r *rest) BindUri(ctx *gin.Context, obj interface{}) error {
	err := ctx.ShouldBindUri(obj)
	if err != nil {
		return errors.NewWithCode(codes.CodeBadRequest, err.Error())
	}

	return nil
}

// Bind all params (query and uri params) to struct using tag 'uri' and 'form'
func (r *rest) BindParams(ctx *gin.Context, obj interface{}) error {
	err := r.BindQuery(ctx, obj)
	if err != nil {
		return errors.NewWithCode(codes.CodeBadRequest, err.Error())
	}

	err = r.BindUri(ctx, obj)
	if err != nil {
		return errors.NewWithCode(codes.CodeBadRequest, err.Error())
	}

	return nil
}

// @Summary Health Check
// @Description This endpoint will hit the server
// @Tags Server
// @Produce json
// @Success 200 {object} entity.Ping
// @Router /ping [GET]
func (r *rest) Ping(ctx *gin.Context) {
	resp := entity.Ping{
		Status:  "OK",
		Message: "PONG!",
		Version: fmt.Sprintf("%s-%s", r.conf.Meta.Version, r.conf.Meta.Environment),
	}

	r.httpRespSuccess(ctx, codes.CodeSuccess, resp, nil)
}
