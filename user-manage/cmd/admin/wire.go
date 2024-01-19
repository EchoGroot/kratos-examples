//go:build wireinject
// +build wireinject

package main

import (
	"github.com/EchoGroot/kratos-examples/user-manage/internal/admin/biz"
	"github.com/EchoGroot/kratos-examples/user-manage/internal/admin/conf"
	"github.com/EchoGroot/kratos-examples/user-manage/internal/admin/data"
	"github.com/EchoGroot/kratos-examples/user-manage/internal/admin/server"
	"github.com/EchoGroot/kratos-examples/user-manage/internal/admin/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

func wireApp(*conf.Bootstrap, log.Logger) (*kratos.App, func(), error) {
	panic(
		wire.Build(
			newApp,
			biz.ProviderSet,
			data.ProviderSet,
			server.ProviderSet,
			service.ProviderSet,
		),
	)
}
