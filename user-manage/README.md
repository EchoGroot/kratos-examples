# user-manage

用户管理。

## 功能

1. 创建用户
    1. 参数校验
2. 修改用户
3. 删除用户
4. 查询用户详情
5. 列表分页查询用户
6. 用于k8s存活探针的livez接口

## 启动项目

1. 创建依赖的数据库容器，`docker compose up -d`
2. 数据库表结构`user-manage/init/postgres/sql`会在项目启动时自动创建，实现**数据库版本变更自管理**

## 构建

1. 构建二进制可执行文件`MODULE_PREFIX=user-manage SUB_MODULE=admin make build`
2. 构建镜像`MODULE_PREFIX=user-manage SUB_MODULE=admin make buildImage`
