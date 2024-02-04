package user

import (
	"context"

	userDom "github.com/nafisalfiani/p3-final-project/account-service/domain/user"
	"github.com/nafisalfiani/p3-final-project/account-service/entity"
	"github.com/nafisalfiani/p3-final-project/lib/log"
)

type user struct {
	logger log.Interface
	user   userDom.Interface
}

type Interface interface {
	List(ctx context.Context) ([]entity.User, error)
	Get(ctx context.Context, filter entity.User) (entity.User, error)
	Create(ctx context.Context, user entity.User) (entity.User, error)
	Update(ctx context.Context, user entity.User) (entity.User, error)
	Delete(ctx context.Context, user entity.User) error
}

// Init creates user usecase
func Init(logger log.Interface, userDom userDom.Interface) Interface {
	return &user{
		logger: logger,
		user:   userDom,
	}
}

func (u *user) List(ctx context.Context) ([]entity.User, error) {
	return u.user.List(ctx)
}

func (u *user) Get(ctx context.Context, filter entity.User) (entity.User, error) {
	return u.user.Get(ctx, filter)
}

func (u *user) Create(ctx context.Context, user entity.User) (entity.User, error) {
	return u.user.Create(ctx, user)
}

func (u *user) Update(ctx context.Context, user entity.User) (entity.User, error) {
	return u.user.Update(ctx, user)
}

func (u *user) Delete(ctx context.Context, user entity.User) error {
	return u.user.Delete(ctx, user)
}
