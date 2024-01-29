package server

import (
	"context"

	"github.com/EchoGroot/kratos-examples/pkg/kratosx"

	adminv1 "github.com/EchoGroot/kratos-examples/user-manage/api/user-manage/v1"
	"github.com/EchoGroot/kratos-examples/user-manage/internal/admin/conf"
	"github.com/EchoGroot/kratos-examples/user-manage/internal/admin/service"

	"github.com/go-kratos/kratos/v2/middleware/logging"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(
	c *conf.Bootstrap,
	userService *service.UserService,
	livezService *service.LivezService,
	logger log.Logger,
) *http.Server {
	opts := []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			tracing.Server(),
			kratosx.WriteTraceId2ReplyHeader,
			selector.
				Server(logging.Server(logger)).
				Match(NewWhiteListMatcher()).
				Build(),
			validate.Validator(),
		),
		http.ErrorEncoder(kratosx.ErrorEncoder),
	}
	if c.Server.Http.Network != "" {
		opts = append(opts, http.Network(c.Server.Http.Network))
	}
	if c.Server.Http.Addr != "" {
		opts = append(opts, http.Address(c.Server.Http.Addr))
	}
	if c.Server.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Server.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)

	adminv1.RegisterLivezServiceHTTPServer(srv, livezService)
	adminv1.RegisterUserServiceHTTPServer(srv, userService)

	return srv
}

func NewWhiteListMatcher() selector.MatchFunc {
	whiteList := make(map[string]struct{})
	whiteList["/usermanage.admin.v1.LivezService/Livez"] = struct{}{}
	return func(ctx context.Context, operation string) bool {
		if _, ok := whiteList[operation]; ok {
			return false
		}
		return true
	}
}
