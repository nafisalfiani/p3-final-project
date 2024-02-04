package auth

import (
	context "context"
	"time"

	"github.com/nafisalfiani/p3-final-project/account-service/entity"
	"github.com/nafisalfiani/p3-final-project/account-service/usecase/role"
	"github.com/nafisalfiani/p3-final-project/account-service/usecase/user"
	"github.com/nafisalfiani/p3-final-project/lib/auth"
	"github.com/nafisalfiani/p3-final-project/lib/codes"
	"github.com/nafisalfiani/p3-final-project/lib/errors"
	"github.com/nafisalfiani/p3-final-project/lib/log"
	"github.com/nafisalfiani/p3-final-project/lib/security"
	"go.mongodb.org/mongo-driver/bson/primitive"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type grpcAuth struct {
	log  log.Interface
	sec  security.Interface
	auth auth.Interface
	user user.Interface
	role role.Interface
}

func Init(log log.Interface, sec security.Interface, auth auth.Interface, user user.Interface, role role.Interface) *grpcAuth {
	return &grpcAuth{
		log:  log,
		sec:  sec,
		auth: auth,
		user: user,
		role: role,
	}
}

func (u *grpcAuth) mustEmbedUnimplementedAuthServiceServer() {}

func (a *grpcAuth) Register(ctx context.Context, in *RegisterRequest) (*RegisterResponse, error) {
	hashedPassword, err := a.sec.HashPassword(ctx, in.GetPassword())
	if err != nil {
		return nil, err
	}

	// get default role
	role, err := a.role.Get(ctx, entity.Role{Code: "member"})
	if err != nil {
		return nil, err
	}

	newId := primitive.NewObjectID()
	createReq := entity.User{
		Id:        newId,
		Username:  in.GetUsername(),
		Name:      in.GetName(),
		Password:  hashedPassword,
		Email:     in.GetEmail(),
		Role:      role,
		CreatedAt: time.Now(),
		CreatedBy: newId.Hex(),
	}
	newUser, err := a.user.Create(ctx, createReq)
	if err != nil {
		return nil, err
	}

	res := &RegisterResponse{
		Id:       newUser.Id.Hex(),
		Name:     newUser.Name,
		Username: newUser.Username,
		Email:    newUser.Email,
	}

	return res, nil
}

func (a *grpcAuth) Login(ctx context.Context, in *LoginRequest) (*LoginResponse, error) {
	user, err := a.user.Get(ctx, entity.User{
		Username: in.GetUsername(),
	})
	if err != nil {
		return nil, err
	}

	if !user.IsEmailVerified {
		return nil, errors.NewWithCode(codes.CodeAuth, "email not verified yet")
	}

	if !a.sec.CompareHashPassword(ctx, user.Password, in.GetPassword()) {
		return nil, errors.NewWithCode(codes.CodeAuthWrongPassword, "password doesn't match")
	}

	token, err := a.auth.CreateToken(ctx, user.ToAuthUser())
	if err != nil {
		return nil, err
	}

	res := &LoginResponse{
		TokenType:        token.TokenType,
		AccessToken:      token.AccessToken,
		AccessExpiresIn:  token.AccessExpiresIn,
		RefreshToken:     token.RefreshToken,
		RefreshExpiresIn: token.RefreshExpiresIn,
	}

	return res, nil
}

func (a *grpcAuth) CreateRole(ctx context.Context, req *Role) (*Role, error) {
	roleReq := entity.Role{
		Code:   req.GetCode(),
		Scopes: req.GetScopes(),
	}

	role, err := a.role.Create(ctx, roleReq)
	if err != nil {
		return nil, err
	}

	res := &Role{
		Id:     role.Id.Hex(),
		Code:   role.Code,
		Scopes: role.Scopes,
	}

	return res, nil
}

func (a *grpcAuth) ListRole(ctx context.Context, in *emptypb.Empty) (*RoleList, error) {
	roles, err := a.role.List(ctx)
	if err != nil {
		return nil, err
	}

	res := &RoleList{}
	for i := range roles {
		res.Roles = append(res.Roles, &Role{
			Id:     roles[i].Id.Hex(),
			Code:   roles[i].Code,
			Scopes: roles[i].Scopes,
		})
	}

	return res, nil
}
