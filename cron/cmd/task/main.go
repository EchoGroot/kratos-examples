package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/EchoGroot/kratos-examples/pkg/kratos/bootstrap"
	"github.com/EchoGroot/kratos-examples/pkg/kratos/contrib/log/logrus"

	"github.com/EchoGroot/kratos-examples/cron/internal/task/conf"
	"github.com/EchoGroot/kratos-examples/cron/internal/task/server"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	Name = "cron-task"
	// Version 可通过编译参数覆盖
	Version     = "dev"
	ID, _       = os.Hostname()
	ServiceInfo = bootstrap.NewServiceInfo(Name, Version, ID)
	Flags       = bootstrap.NewCommandFlags()
)

func init() {
	Flags.Init()
}

func main() {
	// 读取命令行参数
	flag.Parse()

	// 加载日志
	kratosLogger, err := bootstrap.NewLoggerProvider(&ServiceInfo, Flags.LogLevel)
	if err != nil {
		logrus.Panicf("load logger error: %+v", err)
	}

	// 加载配置文件
	var bc conf.Bootstrap
	if err = bootstrap.NewConfigProvider(Flags.Conf, &bc); err != nil {
		log.Log(log.LevelFatal, log.DefaultMessageKey, fmt.Sprintf("load config file error: %+v", err))
	}

	// 依赖注入
	app, cleanup, err := wireApp(&bc, kratosLogger.Logger)
	if err != nil {
		log.Log(log.LevelFatal, log.DefaultMessageKey, fmt.Sprintf("wire app error: %+v", err))
	}
	defer cleanup()

	// 启动项目
	if err := app.Run(); err != nil {
		log.Log(log.LevelFatal, log.DefaultMessageKey, fmt.Sprintf("run app error: %+v", err))
	}
}

func newApp(c *server.BackupDbServer) *kratos.App {
	return kratos.New(
		kratos.ID(ID),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Server(
			c,
		),
	)
}
