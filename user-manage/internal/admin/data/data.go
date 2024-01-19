package data

import (
	"github.com/EchoGroot/kratos-examples/pkg/gormx/postgres"
	"github.com/EchoGroot/kratos-examples/user-manage/internal/admin/conf"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

func NewData(c *conf.Bootstrap) (*gorm.DB, func(), error) {
	return postgres.NewPostgres(
		c.Data.Postgres.Dsn,
		postgres.WithMaxIdleConns(int(c.Data.Postgres.Pool.MaxIdleConns)),
		postgres.WithMaxOpenConns(int(c.Data.Postgres.Pool.MaxOpenConns)),
		postgres.WithConnMaxLifetime(c.Data.Postgres.Pool.ConnMaxLifetime.AsDuration()),
		postgres.WithLogger(log.GetLogger()),
	)
}
