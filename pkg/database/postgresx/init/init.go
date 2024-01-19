package postgresx

import (
	"context"
	"embed"
	"github.com/EchoGroot/kratos-examples/pkg/lang/syncx"
	"strings"
	"time"

	"github.com/go-kratos/kratos/v2/log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"

	"github.com/pkg/errors"
)

// InitDb 用于项目启动时初始化数据库
func InitDb(databaseUrl string, timeout time.Duration, fs embed.FS, initSqlPath string) (err error) {
	sourceInstance, err := iofs.New(fs, initSqlPath)
	if err != nil {
		return errors.Wrapf(err, "could not open initSqlPath: %s", initSqlPath)
	}

	url := dsn2Url(databaseUrl)
	m, err := migrate.NewWithSourceInstance("iofs", sourceInstance, url)
	if err != nil {
		return errors.Wrap(err, "could not init db migrate")
	}

	m.Log = &kratosLogger{}

	timeoutCtx, cancelFunc := context.WithDeadline(context.Background(), time.Now().Add(timeout))
	defer cancelFunc()

	for {
		done := make(chan struct{})
		syncx.Go(func() {
			err = m.Up()
			done <- struct{}{}
		})

		select {
		case <-timeoutCtx.Done():
			return errors.New("init db timeout")
		case <-done:
			if err != nil {
				if errors.Is(err, migrate.ErrNoChange) {
					_, _ = m.Close()
					return nil
				}
				if errors.Is(err, migrate.ErrLocked) {
					time.Sleep(1 * time.Second)
					continue
				}
				return errors.Wrap(err, "init db failed")
			}
			_, _ = m.Close()
			return nil
		}
	}
}

func dsn2Url(databaseUrl string) string {
	arr := strings.Split(databaseUrl, " ")
	params := make(map[string]string, len(arr))
	for _, kv := range arr {
		pair := strings.Split(kv, "=")
		// nolint:gomnd
		if len(pair) == 2 {
			params[pair[0]] = pair[1]
		}
	}

	url := "postgres://{user}:{password}@{host}:{port}/{dbname}?sslmode=disable&TimeZone=Asia/Shanghai"
	for k, v := range params {
		url = strings.Replace(url, "{"+k+"}", v, -1)
	}
	return url
}

type kratosLogger struct{}

func (r *kratosLogger) Printf(format string, v ...interface{}) {
	log.Debugf(format, v...)
}

// Verbose should return true when verbose logging output is wanted
func (r *kratosLogger) Verbose() bool {
	return true
}
