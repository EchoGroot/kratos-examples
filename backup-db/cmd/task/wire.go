//go:build wireinject
// +build wireinject

package main

import (
	"github.com/EchoGroot/kratos-examples/backup-db/internal/task/biz"
	"github.com/EchoGroot/kratos-examples/backup-db/internal/task/conf"
	"github.com/EchoGroot/kratos-examples/backup-db/internal/task/server"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

func wireApp(*conf.Bootstrap, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(
		newApp,
		server.ProviderSet,
		biz.ProviderSet,
	))
}
