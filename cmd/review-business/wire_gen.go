// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"review-business/internal/biz"
	"review-business/internal/conf"
	"review-business/internal/data"
	"review-business/internal/server"
	"review-business/internal/service"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, registry *conf.Registry, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	discovery := data.NewDiscovery(registry)
	reviewClient := data.NewReviewServiceClient(discovery)
	dataData, cleanup, err := data.NewData(confData, logger, reviewClient)
	if err != nil {
		return nil, nil, err
	}
	businessRepo := data.NewBusinessRepo(dataData, logger)
	businessUsecase := biz.NewBusinessUsecase(businessRepo, logger)
	businessService := service.NewBusinessService(businessUsecase)
	grpcServer := server.NewGRPCServer(confServer, businessService, logger)
	httpServer := server.NewHTTPServer(confServer, businessService, logger)
	app := newApp(logger, grpcServer, httpServer)
	return app, func() {
		cleanup()
	}, nil
}
