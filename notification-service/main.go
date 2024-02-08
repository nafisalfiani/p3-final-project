package main

import (
	"context"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/nafisalfiani/p3-final-project/lib/auth"
	"github.com/nafisalfiani/p3-final-project/lib/broker"
	"github.com/nafisalfiani/p3-final-project/lib/configreader"
	"github.com/nafisalfiani/p3-final-project/lib/email"
	"github.com/nafisalfiani/p3-final-project/lib/log"
	"github.com/nafisalfiani/p3-final-project/lib/parser"
	"github.com/nafisalfiani/p3-final-project/lib/security"
	"github.com/nafisalfiani/p3-final-project/notification-service/config"
	"github.com/nafisalfiani/p3-final-project/notification-service/domain"
	"github.com/nafisalfiani/p3-final-project/notification-service/handler/grpc"
	"github.com/nafisalfiani/p3-final-project/notification-service/usecase"
)

func main() {
	log.DefaultLogger().Info(context.Background(), "starting server...")

	// init config file
	cfg := configreader.Init(configreader.Options{
		Type:       configreader.Viper,
		ConfigFile: "./config.json",
	})

	// read from config
	config := &config.Application{}
	if err := cfg.ReadConfig(config); err != nil {
		log.DefaultLogger().Fatal(context.Background(), err)
	}

	log.DefaultLogger().Info(context.Background(), config)

	// init logger
	logger := log.Init(config.Log, log.Zerolog)

	// init parser
	parser := parser.InitParser(logger, parser.Options{})

	allConf, _ := parser.JSONParser().Marshal(cfg.AllSettings())
	log.DefaultLogger().Info(context.Background(), string(allConf))

	// init validator
	validator := validator.New(validator.WithRequiredStructEnabled())

	// init security
	sec := security.Init(logger, config.Security)

	// init mail
	mail := email.Init(config.Mail, logger)

	// init broker
	broker, err := broker.Init(config.Broker, parser.JSONParser())
	if err != nil {
		logger.Fatal(context.Background(), err)
	}
	defer broker.Close()

	// init auth
	auth := auth.Init(config.Auth, logger, parser.JSONParser(), http.DefaultClient)

	// init domain
	dom := domain.Init(logger, parser.JSONParser(), broker, mail)

	// init usecase
	uc := usecase.Init(logger, dom)

	// TODO: init consumer

	// TODO: init scheduler

	// init grpc
	grpc := grpc.Init(config.Grpc, logger, uc, sec, auth, validator)

	// start grpc server
	grpc.Run()
}