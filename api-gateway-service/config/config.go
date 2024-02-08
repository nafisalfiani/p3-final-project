package config

import (
	"github.com/nafisalfiani/p3-final-project/api-gateway-service/handler/rest"
	"github.com/nafisalfiani/p3-final-project/api-gateway-service/handler/scheduler"
	"github.com/nafisalfiani/p3-final-project/api-gateway-service/usecase"
	"github.com/nafisalfiani/p3-final-project/lib/auth"
	"github.com/nafisalfiani/p3-final-project/lib/broker"
	"github.com/nafisalfiani/p3-final-project/lib/log"
	"github.com/nafisalfiani/p3-final-project/lib/security"
)

type Application struct {
	Auth     auth.Config
	Log      log.Config
	Security security.Config
	Broker   broker.Config
	Rest     rest.Config
	Jobs     scheduler.Config
	Usecase  usecase.Config
}
