package postgres

import (
	"embed"
	"time"

	postgresx "github.com/EchoGroot/kratos-examples/pkg/database/postgresx/init"
)

var (
	// `go embed` 仅能嵌入当前目录及其子目录，无法嵌入上层目录。同时也不支持软链接。
	//go:embed sql/*.sql
	fs embed.FS
	// 由于 go:embed 可以配置多个目录，这里还需要指定下
	initSqlPath = "sql"
)

func InitDb(databaseUrl string, timeout time.Duration) error {
	return postgresx.InitDb(databaseUrl, timeout, fs, initSqlPath)
}
