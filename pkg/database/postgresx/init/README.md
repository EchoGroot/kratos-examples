Golang 项目启动时维护数据库变更

## 背景

数据库变更管理是软件发布必不可少的环节，理想状态是只需发布一个镜像，就能更新应用和数据库。我们项目使用 gorm
来操作数据库，gorm是具有数据库迁移功能的，但是没有 SQL 脚本直观。另外我们的应用是同库多服务的微服务，还有些服务存在多个实例的情况，这就需要考虑数据竞争问题了。经过调研，最终选择了
Github 10k star 的 [golang-migrate](https://github.com/golang-migrate)。

## 使用

### 准备 SQL 脚本

将初始化脚本、升级脚本放在项目里的 init/postgres/sql 目录下。

```
init
└── postgres
    ├── init.go
    └── sql
        ├── 20230113084913_1.0.0.up.sql
        └── 20230114084930_1.1.0.up.sql
```

+ 脚本命名
    + 前面部分是一个整数，体现 version 的大小关系，这里用时间表示，你也可以用 001，002。
    + 后面部分是描述信息，仅仅是给程序员看，这里用的应用版本号。

+ MySQL 和 Oracle 不支持 DDL 回滚，但 PG 是可以的。

+ 整个脚本用事务包裹，保证原子性。

+ 使用 `create table if not exits`，`insert into on conflict do nothing`，支持重复执行。

+ 20230113084913_1.0.0.up.sql
    ```sql
    BEGIN;
    CREATE TABLE IF NOT EXISTS users(
    xxx
    );
    CREATE TABLE IF NOT EXISTS users_1(
    xxx
    );
    COMMIT;
    ```

### 写代码

```go
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
```

+ databaseUrl：gorm.io/driver/postgres 里的 dsn`"host=xxxx port=xxx user=postgres password=xxx dbname=xx sslmode=disable
  TimeZone=Asia/Shanghai"`
+ 数据竞争问题：golang-migrate 具备锁功能，它是通过 pg
  的咨询锁实现的数据库级别的锁。如果获取锁失败会返回 migrate.ErrLocked 错误，于是通过它加了个轮询，多实例对同一个库更新就没问题了。

## 注意事项

1. 不同git仓库的程序使用同一个数据库时，各自的SQL脚本放在自己的仓库，同时需要设置不同的`db_version`表。
    ```go
    postgres.DefaultMigrationsTable = "schema_migrations_xxx"
    ```
