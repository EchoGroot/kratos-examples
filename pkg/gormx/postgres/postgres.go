package postgres

import (
	"fmt"
	"github.com/EchoGroot/kratos-examples/pkg/gormx"
	"log"
	"os"
	"time"

	"gorm.io/gorm/logger"

	kratoslog "github.com/go-kratos/kratos/v2/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Postgres postgres.Config
	Pool     *gormx.PoolConfig
}

type Option func(*options)

type options struct {
	maxIdleConns    int
	maxOpenConns    int
	connMaxLifetime time.Duration
	writer          logger.Writer
}

func WithMaxIdleConns(n int) Option {
	return func(o *options) {
		o.maxIdleConns = n
	}
}

func WithMaxOpenConns(n int) Option {
	return func(o *options) {
		o.maxOpenConns = n
	}
}

func WithConnMaxLifetime(d time.Duration) Option {
	return func(o *options) {
		o.connMaxLifetime = d
	}
}

type Writer struct {
	logger kratoslog.Logger
}

func (w *Writer) Printf(format string, args ...interface{}) {
	_ = w.logger.Log(kratoslog.LevelDebug, kratoslog.DefaultMessageKey, fmt.Sprintf(format, args...))
}

func WithLogger(l kratoslog.Logger) Option {
	return func(o *options) {
		o.writer = &Writer{logger: l}
	}
}

func NewPostgres(dsn string, opts ...Option) (*gorm.DB, func(), error) {
	op := options{
		writer: log.New(os.Stdout, "\r\n", log.LstdFlags),
	}
	for _, o := range opts {
		o(&op)
	}

	db, err := gorm.Open(
		postgres.New(postgres.Config{DSN: dsn}),
		&gorm.Config{
			Logger: logger.New(
				op.writer,
				logger.Config{
					// 打印所有sql，如果是warn，则只打印慢sql
					LogLevel:                  logger.Info,
					IgnoreRecordNotFoundError: true,
					Colorful:                  true,
				}),
		})
	if err != nil {
		return nil, nil, err
	}

	connPoolDb, err := db.DB()
	if err != nil {
		return nil, nil, err
	}

	connPoolDb.SetMaxIdleConns(op.maxIdleConns)
	connPoolDb.SetMaxOpenConns(op.maxOpenConns)
	connPoolDb.SetConnMaxLifetime(op.connMaxLifetime)

	return db, func() { connPoolDb.Close() }, nil
}
