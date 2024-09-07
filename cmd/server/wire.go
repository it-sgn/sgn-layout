//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/it-sgn/sgn-layout/internal/biz"
	"github.com/it-sgn/sgn-layout/internal/conf"
	"github.com/it-sgn/sgn-layout/internal/data"
	"github.com/it-sgn/sgn-layout/internal/server"
	"github.com/it-sgn/sgn-layout/internal/service"
)

// wireApp init kratos application.
//
//	func wireApp(*conf.Env, *conf.Server, *conf.Data, *conf.Bootstrap, log.Logger, *conf.Registry, *conf.Trace) (*kratos.App, func(), error) {
//		panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
//	}
//
//	func wireApp(*conf.Env, *conf.Server, *conf.Data, *conf.Bootstrap, log.Logger, *conf.Registry, *conf.Trace) (*kratos.App, func(), error) {
//		panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
//	}
func wireApp(*conf.Env, *conf.Server, *conf.Data, *conf.Bootstrap, log.Logger, *conf.Registry, *conf.Trace) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
