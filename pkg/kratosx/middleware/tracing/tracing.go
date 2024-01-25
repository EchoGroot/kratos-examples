package tracing

import (
	"context"

	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport"
)

// WriteTraceId2ReplyHeader 将 trace_id 写入响应头里
// 可用于返回 trace_id 给前端，方便定位日志
func WriteTraceId2ReplyHeader(handler middleware.Handler) middleware.Handler {
	return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
		reply, err = handler(ctx, req)
		if tr, ok := transport.FromServerContext(ctx); ok {
			tr.ReplyHeader().Set("trace-id", tracing.TraceID()(ctx).(string))
		}
		return
	}
}
