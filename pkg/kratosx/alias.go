package kratosx

import (
	"github.com/EchoGroot/kratos-examples/pkg/kratosx/bootstrap"
	"github.com/EchoGroot/kratos-examples/pkg/kratosx/client"
	"github.com/EchoGroot/kratos-examples/pkg/kratosx/contrib/log/logrus"
	"github.com/EchoGroot/kratos-examples/pkg/kratosx/encoding/json"
	"github.com/EchoGroot/kratos-examples/pkg/kratosx/log"
	"github.com/EchoGroot/kratos-examples/pkg/kratosx/middleware/tracing"
	"github.com/EchoGroot/kratos-examples/pkg/kratosx/transport/http"
)

type CommandFlags = bootstrap.CommandFlags

type KratosLog = bootstrap.KratosLog

type ServiceInfo = bootstrap.ServiceInfo

type LogrusLogger = logrus.LogrusLogger

var (
	NewConfigProvider           = bootstrap.NewConfigProvider
	NewCommandFlags             = bootstrap.NewCommandFlags
	NewLoggerProvider           = bootstrap.NewLoggerProvider
	NewServiceInfo              = bootstrap.NewServiceInfo
	NewTracerProvider           = bootstrap.NewTracerProvider
	NewKeyAuthPerRPCCredentials = client.NewKeyAuthPerRPCCredentials
	NewLogrusLogger             = logrus.NewLogrusLogger
	ConfigMarshal               = json.ConfigMarshal
	WriteTraceId2ReplyHeader    = tracing.WriteTraceId2ReplyHeader
	ErrorEncoder                = http.ErrorEncoder
	NilFilterResponseEncoder    = http.NilFilterResponseEncoder
	LogFatal                    = log.LogFatal
	LogFatalf                   = log.LogFatalf
	LogFatalw                   = log.LogFatalw
)
