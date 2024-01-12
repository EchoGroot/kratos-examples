package kratoslog

import (
	"context"
	"fmt"

	"github.com/bytedance/gopkg/util/logger"
	kratoslog "github.com/go-kratos/kratos/v2/log"
)

type byteDanceLogger struct {
	kratosLogger kratoslog.Logger
	level        kratoslog.Level
}

func NewByteDanceLogger(kratosLogger kratoslog.Logger, level kratoslog.Level) logger.Logger {
	return &byteDanceLogger{
		kratosLogger: kratosLogger,
		level:        level,
	}
}

func (r *byteDanceLogger) logf(lv logger.Level, format *string, v ...interface{}) {
	level := Level2KratosLogLevel(lv)
	if r.level > level {
		return
	}
	var msg string
	if format != nil {
		msg = fmt.Sprintf(*format, v...)
	} else {
		msg = fmt.Sprint(v...)
	}
	_ = r.kratosLogger.Log(level, kratoslog.DefaultMessageKey, msg)
}

func Level2KratosLogLevel(level logger.Level) kratoslog.Level {
	switch level {
	case logger.LevelTrace, logger.LevelDebug:
		return kratoslog.LevelDebug
	case logger.LevelInfo, logger.LevelNotice:
		return kratoslog.LevelInfo
	case logger.LevelWarn:
		return kratoslog.LevelWarn
	case logger.LevelError:
		return kratoslog.LevelError
	case logger.LevelFatal:
		return kratoslog.LevelFatal
	default:
		return kratoslog.LevelDebug
	}
}

func (r *byteDanceLogger) Fatal(v ...interface{}) {
	r.logf(logger.LevelFatal, nil, v...)
}

func (r *byteDanceLogger) Error(v ...interface{}) {
	r.logf(logger.LevelError, nil, v...)
}

func (r *byteDanceLogger) Warn(v ...interface{}) {
	r.logf(logger.LevelWarn, nil, v...)
}

func (r *byteDanceLogger) Notice(v ...interface{}) {
	r.logf(logger.LevelNotice, nil, v...)
}

func (r *byteDanceLogger) Info(v ...interface{}) {
	r.logf(logger.LevelInfo, nil, v...)
}

func (r *byteDanceLogger) Debug(v ...interface{}) {
	r.logf(logger.LevelDebug, nil, v...)
}

func (r *byteDanceLogger) Trace(v ...interface{}) {
	r.logf(logger.LevelTrace, nil, v...)
}

func (r *byteDanceLogger) Fatalf(format string, v ...interface{}) {
	r.logf(logger.LevelFatal, &format, v...)
}

func (r *byteDanceLogger) Errorf(format string, v ...interface{}) {
	r.logf(logger.LevelError, &format, v...)
}

func (r *byteDanceLogger) Warnf(format string, v ...interface{}) {
	r.logf(logger.LevelWarn, &format, v...)
}

func (r *byteDanceLogger) Noticef(format string, v ...interface{}) {
	r.logf(logger.LevelNotice, &format, v...)
}

func (r *byteDanceLogger) Infof(format string, v ...interface{}) {
	r.logf(logger.LevelInfo, &format, v...)
}

func (r *byteDanceLogger) Debugf(format string, v ...interface{}) {
	r.logf(logger.LevelDebug, &format, v...)
}

func (r *byteDanceLogger) Tracef(format string, v ...interface{}) {
	r.logf(logger.LevelTrace, &format, v...)
}

func (r *byteDanceLogger) CtxFatalf(ctx context.Context, format string, v ...interface{}) {
	r.logf(logger.LevelFatal, &format, v...)
}

func (r *byteDanceLogger) CtxErrorf(ctx context.Context, format string, v ...interface{}) {
	r.logf(logger.LevelError, &format, v...)
}

func (r *byteDanceLogger) CtxWarnf(ctx context.Context, format string, v ...interface{}) {
	r.logf(logger.LevelWarn, &format, v...)
}

func (r *byteDanceLogger) CtxNoticef(ctx context.Context, format string, v ...interface{}) {
	r.logf(logger.LevelNotice, &format, v...)
}

func (r *byteDanceLogger) CtxInfof(ctx context.Context, format string, v ...interface{}) {
	r.logf(logger.LevelInfo, &format, v...)
}

func (r *byteDanceLogger) CtxDebugf(ctx context.Context, format string, v ...interface{}) {
	r.logf(logger.LevelDebug, &format, v...)
}

func (r *byteDanceLogger) CtxTracef(ctx context.Context, format string, v ...interface{}) {
	r.logf(logger.LevelTrace, &format, v...)
}
