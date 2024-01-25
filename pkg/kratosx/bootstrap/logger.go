package bootstrap

import (
	"github.com/EchoGroot/kratos-examples/pkg/bytedance/util/logger/kratoslog"
	"github.com/EchoGroot/kratos-examples/pkg/kratosx/contrib/log/logrus"
	bytedancelog "github.com/bytedance/gopkg/util/logger"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/pkg/errors"
)

type KratosLog struct {
	Logger log.Logger
	Level  log.Level
}

// NewLoggerProvider 创建logger
func NewLoggerProvider(serviceInfo *ServiceInfo, flagsLogLevel string) (*KratosLog, error) {
	logLevel, err := logrus.ParseLevel(flagsLogLevel)
	if err != nil {
		return nil, errors.Wrap(err, "parse log level error")
	}

	logger := log.With(logrus.NewLogrusLogger(logrus.Level(logLevel)),
		"serviceId", serviceInfo.Id,
		"serviceName", serviceInfo.Name,
		"serviceVersion", serviceInfo.Version,
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"trace_id", tracing.TraceID(),
		"span_id", tracing.SpanID(),
	)

	// 设置kratos全局logger
	log.SetLogger(logger)

	kratosLogLevel := logrus.Level2KratosLogLevel(logLevel)

	// 设置bytedance pkg全局logger，可用于协程池的recovery打日志
	bytedancelog.SetDefaultLogger(kratoslog.NewByteDanceLogger(logger, kratosLogLevel))

	return &KratosLog{logger, kratosLogLevel}, nil
}
