# Java 开发者转向 Go 学习路线

本学习路线针对具有 Java 背景的开发者设计，帮助快速掌握 Go 语言的核心概念和实战应用。

## 📚 学习阶段规划

### 第一阶段：基础语法与概念（1-2 周）

#### 1.1 Go 语言基础介绍
- [Go 官方文档 - 入门指南](https://golang.org/doc/tutorial/getting-started)
- [Go 官方之旅（The Go Tour）](https://tour.golang.org/) - 强烈推荐交互式学习
- [Go 简介（Go by Example）](https://gobyexample.com/) - 包含大量代码示例

#### 1.2 与 Java 的对比学习
- [Go vs Java 对比指南](https://www.infoq.com/articles/go-for-java-developers/)
- [从 Java 到 Go 的转换](https://blog.golang.org/hello)

#### 1.3 核心语法
- 变量声明与初始化
- 数据类型（基本类型、复合类型）
- 控制流（if/else、for、switch）
- 函数和方法
- 指针（Go 中的指针与 Java 引用的区别）
- 接口（Interface）
- Goroutine 和通道（Channel）- Go 的核心特性

**推荐资源：**
- [Go 官方博客](https://go.dev/blog) - 官方深度文章
- [Effective Go](https://golang.org/doc/effective_go) - Go 最佳实践

---

### 第二阶段：面向对象与类型系统（1 周）

#### 2.1 结构体和方法
- Go 中的结构体（Struct）vs Java 中的类
- 方法接收者（Receiver）概念
- 值接收者 vs 指针接收者

#### 2.2 接口系统
- 隐式接口实现（vs Java 显式 implements）
- 类型断言和类型开关
- 空接口（interface{}）的使用

#### 2.3 包（Package）体系
- 包管理和可见性（大小写规则）
- 模块化设计

**推荐资源：**
- [《Go 程序设计语言》中文版](https://book.douban.com/subject/26912139/) - 书籍
- [Go 设计模式](https://github.com/tmrts/go-patterns) - GitHub 项目

---

### 第三阶段：并发编程（2 周）

#### 3.1 Goroutine 和通道
- Goroutine 简介（轻量级线程 vs Java Thread）
- 通道（Channel）通信
- 缓冲通道和无缓冲通道
- Select 语句

#### 3.2 并发模式
- Worker 模式
- 扇出/扇入（Fan-out/Fan-in）
- 超时处理
- 等等模式

#### 3.3 同步和互斥
- sync 包（Mutex、WaitGroup、Once 等）
- 原子操作
- Context 包

**推荐资源：**
- [Go 并发模式](https://go.dev/blog/pipelines) - 官方系列文章
- [高级 Go 并发模式](https://go.dev/blog/io2013-talk-concurrency)
- [Context 包深入解析](https://pkg.go.dev/context)

---

### 第四阶段：标准库和实战（2-3 周）

#### 4.1 IO 和文件操作
- io 包和 bufio 包
- 文件读写
- JSON 序列化/反序列化

#### 4.2 网络编程
- net 包基础
- HTTP 客户端和服务器
- 构建简单的 REST API

#### 4.3 测试和调试
- testing 包（单元测试）
- 基准测试（Benchmark）
- 覆盖率报告
- 调试工具（Delve）

**推荐资源：**
- [Go 标准库文档](https://pkg.go.dev/std) - 官方 API 文档
- [Go Web 开发教程](https://golang.org/doc/articles/wiki/)

---

### 第五阶段：Web 框架与项目实战（3-4 周）

#### 5.1 Web 框架选择
- **Gin** - 轻量级、高性能（推荐入门）
  - [Gin 官方文档](https://gin-gonic.com/)
  - [Gin GitHub 仓库](https://github.com/gin-gonic/gin)

- **Echo** - 简洁而强大
  - [Echo 官方文档](https://echo.labstack.com/)

- **标准库** - net/http（学习基础）
  - [Go HTTP 服务教程](https://golang.org/doc/articles/wiki/)

#### 5.2 数据库操作
- database/sql 包
- 流行 ORM：
  - [GORM](https://gorm.io/) - 类似 Java 的 ORM 体验
  - [Ent](https://entgo.io/) - 图基 ORM
- SQL 查询和预处理

#### 5.3 项目实战
- 构建完整的 REST API 项目
- 添加认证和授权
- 错误处理和日志
- 部署和容器化

**推荐资源：**
- [Gin 教程系列](https://tutorialedge.com/golang/) - 教程网站
- [Go 项目布局标准](https://github.com/golang-standards/project-layout) - 项目结构参考

---

### 第六阶段：进阶话题（持续学习）

#### 6.1 性能优化
- 内存管理和垃圾回收
- CPU 分析和内存分析
- 优化建议

#### 6.2 微服务和分布式
- gRPC 和 Protocol Buffers
- 服务发现
- 消息队列集成

#### 6.3 云原生开发
- Docker 容器化
- Kubernetes 基础
- 云部署（AWS、GCP、Azure）

**推荐资源：**
- [gRPC Go 教程](https://grpc.io/docs/languages/go/)
- [Go 性能优化指南](https://github.com/dgryski/go-perfbook)

---

## 🔧 开发环境配置

### 安装 Go
- [官方下载页面](https://golang.org/dl/)
- 版本推荐：1.21 或更新

### IDE 和编辑器
- **VS Code** + Go 扩展（推荐）
  - [Go for VS Code](https://marketplace.visualstudio.com/items?itemName=golang.Go)
- **GoLand** (JetBrains)
  - [GoLand 官方网站](https://www.jetbrains.com/go/)
- **Vim/Neovim** - 配置 gopls

### 必要工具
```bash
# Go modules 管理
go mod init <project-name>
go get <package>

# 代码格式化
go fmt ./...
gofmt -w .

# 代码分析
go vet ./...
golangci-lint run
```

---

## 📖 学习资源总结

| 资源类型 | 推荐资源 | 说明 |
|--------|--------|------|
| 官方文档 | [Go 官方网站](https://golang.org/) | 权威来源 |
| 交互式学习 | [The Go Tour](https://tour.golang.org/) | 必做 |
| 代码示例 | [Go by Example](https://gobyexample.com/) | 参考查询 |
| 书籍 | 《Go 程序设计语言》 | 深入学习 |
| 博客 | [Go 官方博客](https://go.dev/blog) | 最新信息 |
| GitHub | [Go 语言代码仓库](https://github.com/golang) | 源码参考 |
| 社区 | [Go 官方论坛](https://forum.golangbridge.org/) | 问题讨论 |
| 包管理 | [pkg.go.dev](https://pkg.go.dev/) | 包文档查询 |

---

## 💡 从 Java 到 Go 的关键差异

| 特性 | Java | Go |
|------|------|-----|
| 并发 | Thread + 框架 | Goroutine + Channel |
| 类型系统 | 显式继承 | 接口隐式实现 |
| 包管理 | Maven/Gradle | Go Modules |
| 编译 | 字节码 | 原生二进制 |
| 启动速度 | 秒级（需要 JVM） | 毫秒级 |
| 内存占用 | 较大 | 很小 |
| 部署 | JAR 包 + JVM | 单一二进制 |
| 错误处理 | 异常机制 | 错误值（if err != nil） |

---

## ✅ 学习检查清单

- [ ] 完成 Go Tour 交互式教程
- [ ] 理解 Goroutine 和 Channel 概念
- [ ] 编写第一个 HTTP 服务器
- [ ] 实现简单的 REST API（使用 Gin 或标准库）
- [ ] 编写单元测试和基准测试
- [ ] 完成一个完整的项目（如任务管理系统、博客 API 等）
- [ ] 学习 Docker 容器化和部署
- [ ] 探索一个实际的 Go 开源项目

---

## 🚀 快速上手项目推荐

1. **CLI 工具** - 从简单的命令行工具开始
   - [Cobra](https://cobra.dev/) - CLI 框架

2. **REST API** - 构建简单的 TODO API
   - 使用 Gin + SQLite
   - 添加认证功能

3. **WebSocket 应用** - 实时聊天系统
   - 学习 WebSocket 和并发

4. **开源贡献** - 参与 Go 生态项目
   - 从简单的 issue 开始
   - [Good First Issues](https://goodfirstissue.dev/)

---

**祝你学习愉快！Go 是一门高效且优雅的语言，相信你会快速上手！** 🎉
