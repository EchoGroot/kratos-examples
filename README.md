# kratos-examples

生产实践的Go和Kratos框架项目示例，封装常用工具，以满足实际开发工作中的需求。

## 和kratos官方examples项目的区别

1. 官方的示例比较简单，实际开发工作中，需要更完善的功能。
2. 官方的示例仅仅是针对某一个功能的示例，实际工作中，往往需要自己整合，而来找示例的人多半不熟悉这块。

## 项目结构

项目主要由以下几个部分组成：

1. backup-db: 定时任务示例，根据cron表达式定时执行。
2. user-manage: web示例，基本的CRUD，分页查询，gorm封装，数据库变更自管理。

## 具体功能

1. kratos框架的基本使用，常用功能的封装
    1. http返回错误信息处理，隐藏500以上的错误，避免底层错误暴露给调用者
    2. 配置链路追踪，将 trace_id 写入响应头里，可用于返回 trace_id 给前端，方便定位日志
    3. 响应枚举值配置，请求响应枚举值的数字
   4. 通过protoc-gen-validate实现参数校验
2. logrus日志库
3. cron定时任务
4. golangci-lint代码检查
5. wire依赖注入
6. protobuf定义配置文件
7. 通过protoc-gen-validate实现**配置文件校验**
8. 常用脚本封装到Makefile
9. bytedance的协程池
10. 使用[golang-migrate](https://github.com/golang-migrate/migrate)
    ，在服务启动时**维护数据库变更**，自动执行初始化SQL，[详细介绍](./pkg/database/postgresx/init/README.md)
11. 通过grom操作postgres数据库
12. gorm
    1. 封装gorm的辅助工具类，提供了基础的CRUD方法，通过泛型实现。
       命名参照[mybatisplus的mapper](https://baomidou.com/pages/49cc81/#mapper-crud-%E6%8E%A5%E5%8F%A3)
    2. 使用BeforeCreate钩子函数，自动生成id
13. 打包镜像
14. k8s存活探针接口，日志过滤
