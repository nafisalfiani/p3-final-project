package main

import (
	"context"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/nafisalfiani/p3-final-project/account-service/config"
	"github.com/nafisalfiani/p3-final-project/account-service/domain"
	"github.com/nafisalfiani/p3-final-project/account-service/handler/grpc"
	"github.com/nafisalfiani/p3-final-project/account-service/usecase"
	"github.com/nafisalfiani/p3-final-project/lib/auth"
	"github.com/nafisalfiani/p3-final-project/lib/broker"
	"github.com/nafisalfiani/p3-final-project/lib/cache"
	"github.com/nafisalfiani/p3-final-project/lib/configreader"
	"github.com/nafisalfiani/p3-final-project/lib/log"
	"github.com/nafisalfiani/p3-final-project/lib/nosql"
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
	config := &config.Application{}
	if err := cfg.ReadConfig(config); err != nil {
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

	// init database connection
	db, err := nosql.Init(config.NoSql)
	if err != nil {
		logger.Fatal(context.Background(), err)
	}

	// init cache
	cache := cache.Init(config.Cache, logger)

	// init broker
	broker, err := broker.Init(config.Broker, logger, parser.JSONParser())
	if err != nil {
		logger.Fatal(context.Background(), err)
	}
	defer broker.Close()

	// init auth
	auth := auth.Init(config.Auth, logger, parser.JSONParser(), http.DefaultClient)

	// init domain
	dom := domain.Init(logger, parser.JSONParser(), db, cache, broker)

	// init usecase
	uc := usecase.Init(logger, broker, dom)

	// TODO: init consumer

	// TODO: init scheduler

	// init grpc
	grpc := grpc.Init(config.Grpc, logger, uc, sec, auth, validator)

	// start grpc server
	grpc.Run()
}
