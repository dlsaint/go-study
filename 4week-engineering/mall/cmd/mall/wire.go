//go:build wireinject
// +build wireinject

package main

import (
	"geekstudy.example/go/4week-engineering/mall/internal/biz"
	"geekstudy.example/go/4week-engineering/mall/internal/conf"
	"geekstudy.example/go/4week-engineering/mall/internal/data"
	"geekstudy.example/go/4week-engineering/mall/internal/server"
	"geekstudy.example/go/4week-engineering/mall/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

func initApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
