# kratos-examples

Go微服务框架Kratos（Bilibili开源）最佳实践。封装常用工具，以满足实际开发工作中的需求，提升开发效率。

> 彦祖，给个Star

## 和kratos官方examples项目的区别

1. 官方的示例比较简单，实际开发工作中，需要更完善的功能。
2. 官方的示例仅仅是针对某一个功能的示例，实际工作中，往往需要自己整合。

## 示例项目

1. backup-db: 定时任务示例，根据cron表达式定时执行。
2. user-manage: 微服务示例，基本的CRUD，参数校验，分页查询，数据库变更自管理。

## 具体功能

1. kratos框架的基本使用，常用功能的封装
   1. 封装配置读取、校验
   2. 日志组件初始化
   3. 配置链路追踪，将 trace_id 写入响应头里，可用于返回 trace_id 给前端，方便定位日志
   4. http返回错误信息处理，隐藏500以上的错误，避免底层错误暴露给调用者
   5. 重新实现kratos的log.Fatal()，log.Fatal()内部调用os.Exit()，会造成defer函数无法执行，改为调用panic
   6. 响应枚举值配置，请求响应枚举值的数字
   7. 通过protoc-gen-validate实现参数校验
2. 利用反射，高效处理响应值里的time类型字段，转换为Protocol Buffer的timestamp类型
3. 基于[Google API规范](https://cloud.google.com/apis/design?hl=zh-cn)定义API接口
4. **使用FieldMask，解决golang更新零值问题**
5. **使用FieldMask，解决接口响应字段超出客户端所需，造成带宽浪费，还可避免执行不必要的业务逻辑**
6. 使用[golang-migrate](https://github.com/golang-migrate/migrate)
   ，在服务启动时**维护数据库变更**，自动执行初始化SQL，[详细介绍](https://yuyy.info/?p=2087)
7. 通过grom操作postgres数据库
   1. **封装gorm的辅助工具类，提供了基础的CRUD方法，通过泛型实现。**
       命名参照[mybatisplus的mapper](https://baomidou.com/pages/49cc81/#mapper-crud-%E6%8E%A5%E5%8F%A3)
    2. 使用BeforeCreate钩子函数，自动生成id
   3. **封装分页查询操作**
   4. 使用可选函数封装数据库连接初始化
8. logrus日志库
9. cron定时任务
10. protobuf定义配置文件
11. 通过protoc-gen-validate实现**配置文件校验**
12. wire依赖注入
13. 常用脚本封装到Makefile
14. k8s存活探针接口，日志过滤
15. bytedance的协程池
16. golangci-lint代码检查
17. 打包镜像

