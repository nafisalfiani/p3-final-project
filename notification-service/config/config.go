package config

import (
	"github.com/nafisalfiani/p3-final-project/lib/auth"
	"github.com/nafisalfiani/p3-final-project/lib/broker"
	"github.com/nafisalfiani/p3-final-project/lib/email"
	"github.com/nafisalfiani/p3-final-project/lib/log"
	"github.com/nafisalfiani/p3-final-project/lib/security"
	"github.com/nafisalfiani/p3-final-project/notification-service/handler/grpc"
)

type Application struct {
	ApiGateway ApiGateway
	Auth       auth.Config
	Log        log.Config
	Security   security.Config
	Mail       email.Config
	Broker     broker.Config
	Grpc       grpc.Config
}

type ApiGateway struct {
	Url string
}
