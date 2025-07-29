# Gin Learning Demos

这是一个基于 Gin 框架的学习实践项目，用于记录和分享 Gin 框架的学习过程。目前已实现基础功能，后续会持续扩展更多进阶特性。


## 🌟 目前已实现的功能

- 初始化 Gin 实例并启动 HTTP 服务（默认端口 8080）
- 基础路由与 HTTP 方法示例：
  - `GET /`：返回字符串响应 "Hello, Golang Gin!"
  - `GET /ping`：返回 JSON 响应 `{"message": "pong"}`
  - `GET/POST/PUT/DELETE /resource`：演示 RESTful 风格的资源操作
- HTML 模板渲染：
  - `GET /news`：通过 `templates/news.html` 模板返回动态页面


## 📅 后续开发计划

接下来将逐步扩展以下内容（持续更新）：

- 路由进阶：路由分组、参数绑定、路由嵌套
- 中间件：日志记录、请求耗时统计、跨域处理、身份验证
- 数据交互：表单提交与验证、JSON 数据处理
- 模板引擎：模板继承、公共组件复用、动态数据渲染
- 数据库操作：连接 MySQL/PostgreSQL、ORM 框架使用（如 GORM）
- 错误处理：全局异常捕获、自定义错误响应
- 服务配置：环境变量、配置文件管理
- 部署相关：Docker 容器化、CI/CD 流程



## 📚 学习参考

以下是学习 Gin 框架和 Go Web 开发的常用资源，按类型分类整理：

### 官方文档
- [Gin 官方文档（中文）](https://gin-gonic.com/zh-cn/docs/)  
  包含 Gin 框架的快速入门、核心特性、API 参考等，适合入门到进阶学习。
- [Go 语言官方文档](https://golang.google.cn/doc/)  
  Go 语言基础语法、标准库用法及最佳实践的权威指南。

### 工具与生态
- [GORM 官方文档](https://gorm.io/zh_CN/docs/)  
  Go 语言常用的 ORM 框架，与 Gin 搭配使用可快速实现数据库操作。
- [Postman](https://www.postman.com/)  
  接口调试工具，适合测试本项目中的 RESTful 接口。
- [fresh 热重载工具](https://github.com/pilu/fresh)  
  开发时自动监听代码变化并重启服务，提高开发效率。

欢迎关注本项目的更新，一起学习 Gin 框架的开发实践！如果有任何问题或建议，欢迎提交 Issues 或 PR。