// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"mall-go/app/lottery/internal/config"
	"mall-go/app/lottery/internal/server"
	"mall-go/app/lottery/internal/service"
	"mall-go/app/lottery/internal/svc"
)

// Injectors from wire.go:

// initApp init app application.
func initApp(c *config.Config) (*server.AppServer, error) {
	serviceContext := svc.NewServiceContext(c)
	lotteryService := service.NewLotteryService(serviceContext)
	httpServer := server.NewHttpServer(c, lotteryService)
	rpcServer := server.NewGrpcServer(c, serviceContext)
	appServer, err := server.NewApp(serviceContext, lotteryService, httpServer, rpcServer)
	if err != nil {
		return nil, err
	}
	return appServer, nil
}
