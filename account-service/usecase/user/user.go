package user

import (
	"context"
	"fmt"

	userDom "github.com/nafisalfiani/p3-final-project/account-service/domain/user"
	"github.com/nafisalfiani/p3-final-project/account-service/entity"
	"github.com/nafisalfiani/p3-final-project/lib/broker"
	"github.com/nafisalfiani/p3-final-project/lib/log"
)

type user struct {
	logger log.Interface
	user   userDom.Interface
	broker broker.Interface
}

type Interface interface {
	List(ctx context.Context) ([]entity.User, error)
	Get(ctx context.Context, filter entity.User) (entity.User, error)
	Create(ctx context.Context, user entity.User) (entity.User, error)
	Update(ctx context.Context, user entity.User) (entity.User, error)
	Delete(ctx context.Context, user entity.User) error
}

// Init creates user usecase
func Init(logger log.Interface, userDom userDom.Interface, broker broker.Interface) Interface {
	return &user{
		logger: logger,
		user:   userDom,
		broker: broker,
	}
}

func (u *user) List(ctx context.Context) ([]entity.User, error) {
	return u.user.List(ctx)
}

func (u *user) Get(ctx context.Context, filter entity.User) (entity.User, error) {
	return u.user.Get(ctx, filter)
}

func (u *user) Create(ctx context.Context, user entity.User) (entity.User, error) {
	user, err := u.user.Create(ctx, user)
	if err != nil {
		return user, err
	}

	u.logger.Info(ctx, fmt.Sprintf("publishing new user with id %v", user.Id.Hex()))
	if err := u.broker.PublishMessage(entity.TopicNewRegistration, user); err != nil {
		u.logger.Error(ctx, err)
	}

	return user, nil
}

func (u *user) Update(ctx context.Context, user entity.User) (entity.User, error) {
	return u.user.Update(ctx, user)
}

func (u *user) Delete(ctx context.Context, user entity.User) error {
	return u.user.Delete(ctx, user)
}
