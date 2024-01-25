package logrus

import (
	"io"
	"os"

	"github.com/go-kratos/kratos/v2/log"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
)

var _ log.Logger = (*LogrusLogger)(nil)

type LogrusLogger struct {
	log *logrus.Logger
}

func NewLogrusLogger(options ...Option) log.Logger {
	logger := logrus.New()
	logger.Level = logrus.DebugLevel
	logger.Out = os.Stdout
	logger.Formatter = &logrus.JSONFormatter{}
	for _, option := range options {
		option(logger)
	}

	return &LogrusLogger{
		log: logger,
	}
}

func (l *LogrusLogger) Log(level log.Level, keyvals ...interface{}) (err error) {
	var (
		logrusLevel logrus.Level
		fields      logrus.Fields = make(map[string]interface{})
		msg         string
	)

	logrusLevel = KratosLogLevel2LogrusLevel(level)

	if logrusLevel > l.log.Level {
		return
	}

	if len(keyvals) == 0 {
		return nil
	}
	if len(keyvals)%2 != 0 {
		keyvals = append(keyvals, "")
	}
	for i := 0; i < len(keyvals); i += 2 {
		key, ok := keyvals[i].(string)
		if !ok {
			continue
		}
		if key == logrus.FieldKeyMsg {
			msg, _ = keyvals[i+1].(string)
			continue
		}
		fields[key] = keyvals[i+1]
	}

	if len(fields) > 0 {
		l.log.WithFields(fields).Log(logrusLevel, msg)
	} else {
		l.log.Log(logrusLevel, msg)
	}

	return
}

type Option func(log *logrus.Logger)

func Level(level logrus.Level) Option {
	return func(log *logrus.Logger) {
		log.Level = level
	}
}

func Output(w io.Writer) Option {
	return func(log *logrus.Logger) {
		log.Out = w
	}
}

func FileRotateSize(filePath string, rotationSize int64, rotationCount uint) Option {
	return func(log *logrus.Logger) {
		rotateWriter, err := rotatelogs.New(
			filePath+".%Y%m%d%H%M",
			rotatelogs.WithLinkName(filePath),
			rotatelogs.WithRotationSize(rotationSize),
			rotatelogs.WithRotationCount(rotationCount),
		)
		if err != nil {
			panic(err)
		}
		writer := io.MultiWriter(rotateWriter, os.Stdout)
		log.Out = writer
	}
}

func Formatter(formatter logrus.Formatter) Option {
	return func(log *logrus.Logger) {
		log.Formatter = formatter
	}
}

func ParseLevel(lvl string) (logrus.Level, error) {
	return logrus.ParseLevel(lvl)
}

// Panic logs a message at level Panic on the standard logger.
func Panic(args ...interface{}) {
	logrus.Panic(args...)
}

// Panicf logs a message at level Panic on the standard logger.
func Panicf(format string, args ...interface{}) {
	logrus.Panicf(format, args...)
}

// Panicln logs a message at level Panic on the standard logger.
func Panicln(args ...interface{}) {
	logrus.Panicln(args...)
}

// SetFormatter sets the standard logger formatter.
func SetFormatter(formatter logrus.Formatter) {
	logrus.SetFormatter(formatter)
}

// KratosLogLevel2LogrusLevel 注意，这里将kratos的LevelFatal转换成了logrus.PanicLevel。
// 目的是在main函数里打印fatal日志时，底层使用panic（默认为os.exit），保障defer被执行
func KratosLogLevel2LogrusLevel(level log.Level) logrus.Level {
	switch level {
	case log.LevelDebug:
		return logrus.DebugLevel
	case log.LevelInfo:
		return logrus.InfoLevel
	case log.LevelWarn:
		return logrus.WarnLevel
	case log.LevelError:
		return logrus.ErrorLevel
	case log.LevelFatal:
		return logrus.PanicLevel
	default:
		return logrus.DebugLevel
	}
}

func Level2KratosLogLevel(level logrus.Level) log.Level {
	switch level {
	case logrus.DebugLevel:
		return log.LevelDebug
	case logrus.InfoLevel:
		return log.LevelInfo
	case logrus.WarnLevel:
		return log.LevelWarn
	case logrus.ErrorLevel:
		return log.LevelError
	default:
		return log.LevelDebug
	}
}
