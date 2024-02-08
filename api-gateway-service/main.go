package main

import (
	"context"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/nafisalfiani/p3-final-project/api-gateway-service/config"
	"github.com/nafisalfiani/p3-final-project/api-gateway-service/domain"
	"github.com/nafisalfiani/p3-final-project/api-gateway-service/handler/rest"
	"github.com/nafisalfiani/p3-final-project/api-gateway-service/handler/scheduler"
	"github.com/nafisalfiani/p3-final-project/api-gateway-service/usecase"
	"github.com/nafisalfiani/p3-final-project/lib/auth"
	"github.com/nafisalfiani/p3-final-project/lib/broker"
	"github.com/nafisalfiani/p3-final-project/lib/configreader"
	"github.com/nafisalfiani/p3-final-project/lib/log"
	"github.com/nafisalfiani/p3-final-project/lib/parser"
	"github.com/nafisalfiani/p3-final-project/lib/security"
)

func main() {
	log.DefaultLogger().Info(context.Background(), "starting server...")

	// init config file
	cfg := configreader.Init(configreader.Options{
		Type:       configreader.Viper,
		ConfigFile: "./config.json",
	})

	// read from config
	config := config.Application{}
	if err := cfg.ReadConfig(&config); err != nil {
		log.DefaultLogger().Fatal(context.Background(), err)
	}

	log.DefaultLogger().Info(context.Background(), config)

	// init logger
	logger := log.Init(config.Log, log.Zerolog)

	// init parser
	parser := parser.InitParser(logger, parser.Options{})

	// init validator
	validator := validator.New(validator.WithRequiredStructEnabled())

	// init security
	sec := security.Init(logger, config.Security)

	// init broker
	broker, err := broker.Init(config.Broker, logger, parser.JSONParser())
	if err != nil {
		logger.Fatal(context.Background(), err)
	}
	defer broker.Close()

	// init auth
	auth := auth.Init(config.Auth, logger, parser.JSONParser(), http.DefaultClient)

	// init dom
	dom := domain.Init(logger)

	// init uc
	uc := usecase.Init(config.Usecase, logger, sec, validator, dom)
	defer uc.CloseAllConns()

	// init and run scheduler
	sch := scheduler.Init(config.Jobs, logger, auth, uc)
	sch.Run()

	// init http server
	r := rest.Init(config.Rest, cfg, logger, parser.JSONParser(), auth, uc, sch)

	// run the http server
	r.Run()
}
