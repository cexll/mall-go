//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"mall-go/app/lottery/internal/config"
	"mall-go/app/lottery/internal/server"
	"mall-go/app/lottery/internal/service"
	"mall-go/app/lottery/internal/svc"

	"github.com/google/wire"
)

// initApp init app application.
func initApp(c *config.Config) (*server.AppServer, error) {
	panic(wire.Build(svc.ProviderSet, service.ProviderSet, server.ProviderSet, server.NewApp))
}
