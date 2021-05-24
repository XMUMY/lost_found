// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"github.com/XMUMY/lost_found/internal/biz"
	"github.com/XMUMY/lost_found/internal/conf"
	"github.com/XMUMY/lost_found/internal/data"
	"github.com/XMUMY/lost_found/internal/server"
	"github.com/XMUMY/lost_found/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
)

// Injectors from wire.go:

// initApp init kratos application.
func initApp(confServer *conf.Server, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	dataData, cleanup, err := data.NewData(confData, logger)
	if err != nil {
		return nil, nil, err
	}
	itemRepo := data.NewItemRepo(dataData, logger)
	itemUsecase := biz.NewItemUsecase(itemRepo, logger)
	lostAndFoundService, err := service.NewLostAndFoundService(itemUsecase)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	httpServer := server.NewHTTPServer(confServer, lostAndFoundService, logger)
	grpcServer := server.NewGRPCServer(confServer, lostAndFoundService, logger)
	app := newApp(logger, httpServer, grpcServer)
	return app, func() {
		cleanup()
	}, nil
}