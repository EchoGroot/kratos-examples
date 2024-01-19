package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/EchoGroot/kratos-examples/pkg/kratos/bootstrap"
	"github.com/EchoGroot/kratos-examples/pkg/kratos/contrib/log/logrus"
	"github.com/EchoGroot/kratos-examples/pkg/kratos/encoding/json"
	"github.com/EchoGroot/kratos-examples/user-manage/init/postgres"
	"github.com/EchoGroot/kratos-examples/user-manage/internal/admin/conf"
	"github.com/go-kratos/kratos/v2/transport/grpc"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/http"
)

var (
	Name = "user-manage-admin"
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
	kratosLog, err := bootstrap.NewLoggerProvider(&ServiceInfo, Flags.LogLevel)
	if err != nil {
		logrus.Panicf("load logger error: %+v", err)
	}

	// 配置序列化
	json.ConfigMarshal()

	// 初始化链路追踪
	if err = bootstrap.NewTracerProvider(); err != nil {
		// 底层调用panic（kit库里修改的）。不要使用log.Fatal()，内部调用os.Exit()，造成无法执行defer
		// 不要在 main 函数外退出程序，会造成逻辑割裂，和error一个道理。
		log.Log(log.LevelFatal, log.DefaultMessageKey, fmt.Sprintf("create global tracer error: %+v", err))
		// log.Fatal("xxx")
	}

	// 加载配置文件
	var bc conf.Bootstrap
	if err = bootstrap.NewConfigProvider(Flags.Conf, &bc); err != nil {
		log.Log(log.LevelFatal, log.DefaultMessageKey, fmt.Sprintf("load config file error: %+v", err))
	}

	// 初始化数据库
	if err := postgres.InitDb(bc.Data.Postgres.Dsn, bc.Data.Postgres.Init.Timeout.AsDuration()); err != nil {
		log.Log(log.LevelFatal, log.DefaultMessageKey, fmt.Sprintf("init postgres error: %+v", err))
	}

	// 依赖注入
	app, cleanup, err := wireApp(&bc, kratosLog.Logger)
	if err != nil {
		log.Log(log.LevelFatal, log.DefaultMessageKey, fmt.Sprintf("wire app error: %+v", err))
	}
	defer cleanup()

	// 启动项目
	if err := app.Run(); err != nil {
		log.Log(log.LevelFatal, log.DefaultMessageKey, fmt.Sprintf("run app error: %+v", err))
	}
}

func newApp(hs *http.Server, gs *grpc.Server) *kratos.App {
	return kratos.New(
		kratos.ID(ID),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Server(
			hs,
			gs,
		),
	)
}
