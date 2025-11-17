# 代码示例说明文档

本文件夹包含 Java 开发者学习 Go 的完整代码示例，每个文件对应学习路线中的一个阶段。所有代码都有详细的中文注释，并标注对应的 Java 写法，方便快速理解 Go 与 Java 的差异。

## 📂 文件结构

```
代码示例/
├── 01_变量和常量.go          # 第一阶段：基础语法 - 变量、常量、类型推导
├── 02_数据类型.go            # 第一阶段：基础语法 - 数组、切片、映射、结构体、指针
├── 03_控制流.go              # 第一阶段：基础语法 - if/else、switch、for、defer
├── 04_函数和方法.go          # 第二阶段：面向对象 - 函数、方法、接口、闭包
├── 05_并发编程.go            # 第三阶段：并发编程 - Goroutine、Channel、Mutex
├── 06_字符串操作.go          # 第四阶段：标准库 - 字符串处理、转换、拼接
├── 07_文件和JSON操作.go      # 第四阶段：标准库 - 文件读写、JSON 序列化
├── 08_数据库操作.go          # 第五阶段：项目实战 - SQL、事务、ORM（GORM）
├── 09_HTTP和Web服务.go      # 第五阶段：项目实战 - HTTP 客户端、服务器、REST API
├── 10_测试.go                # 第四阶段：标准库 - 单元测试、基准测试、表驱动测试
└── 11_错误处理.go            # 完整阶段：错误处理 - error 接口、panic/recover、自定义错误
```

## 📖 文件详解

### 01_变量和常量.go
**学习阶段**：第一阶段（基础语法与概念）

**核心内容**：
- 变量声明方式（var、:=、多变量声明）
- 常量声明和 const 块
- 类型推导机制
- 零值（Go 中所有变量有默认值）
- 类型转换（显式转换）
- 匿名变量（_）
- 多返回值函数

**关键对比**：
- Java：需要显式初始化变量，常量使用 `final`
- Go：支持短变量声明 `:=`，自动零初始化

**运行方式**：
```bash
go run 01_变量和常量.go
```

---

### 02_数据类型.go
**学习阶段**：第一阶段（基础语法与概念）

**核心内容**：
- 基本数据类型（int、float、bool、string、rune、byte）
- 数组（固定长度）
- 切片（动态数组）- Go 特有
- 映射（Map）
- 结构体（Struct）- Go 的 "类"
- 指针和指针操作
- 嵌套结构体

**关键对比**：
- Java：数组使用 `int[]`；集合使用 `List<T>`、`Map<K,V>`
- Go：切片 `[]int` 更灵活；映射 `map[string]int` 直接内置
- Java：类通过继承；Go：结构体通过嵌入

**运行方式**：
```bash
go run 02_数据类型.go
```

---

### 03_控制流.go
**学习阶段**：第一阶段（基础语法与概念）

**核心内容**：
- if-else 条件语句（条件中可初始化变量）
- switch 语句（多值 case、无表达式 switch）
- for 循环（C 风格、while 风格、无限循环）
- range 循环（遍历数组、切片、映射、字符串）
- continue 和 break
- 标签和跳转（goto）
- defer 延迟执行

**关键对比**：
- Java：有 while 循环；Go：只有 for
- Java：没有 range；Go：range 遍历更优雅
- Java：没有 defer；Go：defer 用于资源清理（类似 finally）

**运行方式**：
```bash
go run 03_控制流.go
```

---

### 04_函数和方法.go
**学习阶段**：第二阶段（面向对象与类型系统）

**核心内容**：
- 函数定义和调用
- 多返回值（Go 特性）
- 命名返回值
- 可变参数（variadic）
- 函数作为参数（高阶函数）
- 函数作为返回值
- 匿名函数和闭包
- 值接收者和指针接收者方法
- 接口定义和实现（隐式实现）
- 空接口 `interface{}`
- 类型断言
- 错误处理（error 接口）
- 递归

**关键对比**：
- Java：单返回值；Go：多返回值
- Java：显式实现接口；Go：隐式实现
- Java：方法在类中；Go：方法关联到类型

**运行方式**：
```bash
go run 04_函数和方法.go
```

---

### 05_并发编程.go
**学习阶段**：第三阶段（并发编程）**[重点文件]**

**核心内容**：
- Goroutine 创建和管理
- 通道（Channel）基础
- 缓冲通道 vs 无缓冲通道
- 通道关闭和检查
- select 语句（监听多个通道）
- 超时处理
- Worker 池模式（并发任务处理）
- sync.Mutex（互斥锁）
- sync.WaitGroup（等待组）
- sync.Once（执行一次）
- sync.RWMutex（读写锁）
- select 与 default
- Pipeline 模式

**关键对比**：
- Java：Thread + 复杂的同步；Go：Goroutine + Channel（更简洁）
- Java：共享内存通信；Go：通过通信共享内存
- Java：需要显式锁管理；Go：通道更安全

**运行方式**：
```bash
go run 05_并发编程.go
```

⚠️ **注意**：这个文件需要一些时间执行，部分示例有 `time.Sleep()` 延迟。

---

### 06_字符串操作.go
**学习阶段**：第四阶段（标准库和实战）

**核心内容**：
- 字符串声明和多行字符串
- 字符串长度和索引
- Unicode 和 Rune（处理中文等）
- 字节和 Rune 转换
- 字符串与其他类型转换（使用 strconv 包）
- strings 包方法：
  - Contains、HasPrefix、HasSuffix
  - Index、Replace、Split、Join
  - ToUpper、ToLower、TrimSpace
  - Fields、Count
- 字符串拼接（+、fmt.Sprintf、strings.Builder）
- 字符串比较（==、strings.Compare、strings.EqualFold）
- 正则表达式基础

**关键对比**：
- Java：String 方法丰富；Go：大部分在 strings 包中
- Java：字符 char；Go：rune 表示 Unicode 字符
- Java：StringBuilder；Go：strings.Builder

**运行方式**：
```bash
go run 06_字符串操作.go
```

---

### 07_文件和JSON操作.go
**学习阶段**：第四阶段（标准库和实战）

**核心内容**：
- 文件创建和写入
- 文件读取（整体读取、逐行读取）
- 文件追加
- 文件信息查询
- 文件删除
- 目录操作（创建、列出、删除）
- 路径操作（连接、分解、相对/绝对路径）
- JSON 序列化（Marshal、MarshalIndent）
- JSON 反序列化（Unmarshal）
- 动态 JSON 处理（map[string]interface{}）
- JSON 数组处理
- JSON 文件读写

**关键对比**：
- Java：FileReader/FileWriter；Go：os + ioutil
- Java：ObjectMapper；Go：encoding/json
- Java：try-with-resources；Go：defer 自动关闭

**运行方式**：
```bash
go run 07_文件和JSON操作.go
```

---

### 08_数据库操作.go
**学习阶段**：第五阶段（Web 框架与项目实战）

**核心内容**：
- 数据库连接（使用 database/sql 包）
- 创建表
- 插入数据
- 查询单行和多行
- 更新数据
- 删除数据
- 预处理语句（防止 SQL 注入）
- 事务处理
- 连接池配置
- ORM 框架介绍（GORM）

**支持的驱动**：
- SQLite：`github.com/mattn/go-sqlite3`
- MySQL：`github.com/go-sql-driver/mysql`
- PostgreSQL：`github.com/lib/pq`

**关键对比**：
- Java：使用 JDBC 或 JPA；Go：使用 database/sql 或 GORM
- Java：ORM 框架复杂（Hibernate）；Go：GORM 简洁
- Java：HikariCP 连接池；Go：内置连接池

**运行方式**：
```bash
go run 08_数据库操作.go
```

⚠️ **注意**：这是演示代码，实际使用需要导入数据库驱动。

---

### 09_HTTP和Web服务.go
**学习阶段**：第五阶段（Web 框架与项目实战）**[重要文件]**

**核心内容**：
- HTTP GET 请求
- HTTP POST 请求
- 自定义请求头
- 超时设置
- 表单数据提交
- JSON 请求
- 错误处理（HTTP 状态码）
- HTTP 服务器创建
- 请求处理器
- 中间件概念
- JSON 响应
- POST 请求处理
- 简单 REST API 示例
- Web 框架对比（Gin、Echo、标准库）

**关键对比**：
- Java：HttpClient + Spring Boot；Go：net/http 标准库
- Java：Spring MVC 注解路由；Go：HandleFunc 注册路由
- Java：Spring Boot 自动配置；Go：手动配置更灵活

**运行方式**：
```bash
go run 09_HTTP和Web服务.go
```

---

### 10_测试.go
**学习阶段**：第四阶段（标准库和实战）

**核心内容**：
- 基本单元测试
- 表驱动测试（参数化测试）
- 错误测试
- 基准测试（性能测试）
- 子基准测试
- 模糊测试（Fuzz Testing）- Go 1.18+
- 测试设置和清理（setupTest）
- 测试工具函数
- 跳过测试
- 并发测试

**测试文件要求**：
- 文件名：`xxx_test.go`
- 函数名：`TestXxx(t *testing.T)`
- 基准测试：`BenchmarkXxx(b *testing.B)`

**关键对比**：
- Java：JUnit + TestNG；Go：内置 testing 包
- Java：@Test 注解；Go：函数命名约定
- Java：JMH 基准测试；Go：内置 testing.B

**运行方式**：
```bash
# 运行所有测试
go test 10_测试.go

# 运行特定测试
go test -run TestAdd 10_测试.go

# 运行基准测试
go test -bench=. 10_测试.go

# 显示覆盖率
go test -cover 10_测试.go
```

⚠️ **注意**：testing 包中的基准测试和模糊测试需要特定的运行方式。

---

### 11_错误处理.go
**学习阶段**：贯穿所有阶段

**核心内容**：
- 基本错误处理（if err != nil）
- 自定义错误类型
- 错误包装和展开（errors.Unwrap）
- 错误链检查（errors.Is、errors.As）
- Panic 和 Recover（不可恢复错误）
- defer 进行资源清理
- 多个 defer 的执行顺序（后进先出）
- 错误处理最佳实践
- 定义异常体系（AppError）
- 日志记录错误

**关键对比**：
- Java：try-catch-finally；Go：if err != nil + defer
- Java：throw 异常；Go：return error
- Java：异常继承体系；Go：error 接口

**运行方式**：
```bash
go run 11_错误处理.go
```

---

## 🚀 快速开始指南

### 环境要求
- Go 1.18 或更新版本
- VS Code（推荐）+ Go 扩展

### 安装 Go
```bash
# 访问官方下载页面
https://golang.org/dl/

# macOS（使用 Homebrew）
brew install go

# Linux（Ubuntu）
sudo apt-get install golang-go

# 验证安装
go version
```

### 运行代码示例

**方式 1：直接运行单个文件**
```bash
cd /workspaces/cs-lean/go-lean/代码示例
go run 01_变量和常量.go
```

**方式 2：使用 go build 编译**
```bash
go build -o demo 01_变量和常量.go
./demo
```

**方式 3：在 VS Code 中运行**
- 打开文件
- 按 `Ctrl+Alt+N`（需要安装 Code Runner 扩展）
- 或在文件中右键选择 "Run Code"

### 推荐学习顺序

1. **第 1 天**：01、02、03（基础语法）
2. **第 2-3 天**：04（函数和方法）
3. **第 4-7 天**：05（并发编程）- 重点，需要时间理解
4. **第 8-9 天**：06、07、08（标准库）
5. **第 10-11 天**：09（HTTP 和 Web）
6. **第 12 天**：10（测试）
7. **贯穿全过程**：11（错误处理）

---

## 💡 学习建议

### 1. 逐行理解代码
- 阅读代码注释，理解每一行的作用
- 查看 Java 对比，加深理解差异
- 自己修改参数，观察输出变化

### 2. 动手实验
```bash
# 不要只读代码，要运行它！
go run 01_变量和常量.go

# 修改代码，观察输出
# 例如改变变量值，看零值如何变化
```

### 3. 深入研究重点
- **并发编程（05）**：Go 最重要的特性，多花时间理解
- **HTTP 和 Web（09）**：实战中最常用，建议自己修改扩展
- **错误处理（11）**：贯穿整个开发，养成良好习惯

### 4. 参考官方资源
- [Go 官方教程](https://tour.golang.org/)
- [Go 标准库文档](https://pkg.go.dev/std)
- [Effective Go](https://golang.org/doc/effective_go)

### 5. 实战项目
完成代码示例学习后，建议做以下项目：
- **CLI 工具**：使用 Cobra 框架
- **REST API**：使用 Gin + SQLite
- **聊天应用**：使用 WebSocket + Goroutine

---

## 📝 代码文件维护说明

本文档会与代码示例同步更新。当添加或修改代码文件时：

1. **更新文件时**：更新本文档中对应文件的说明
2. **新增文件时**：在本文档中添加新文件的说明
3. **删除文件时**：在本文档中删除对应说明

### 更新历史

| 日期 | 更新内容 |
|------|--------|
| 2025-11-17 | 初始创建，包含 11 个核心代码文件 |

---

## 🔗 相关资源链接

### 官方资源
- [Go 官方网站](https://golang.org/)
- [Go 文档](https://golang.org/doc/)
- [Go 标准库](https://pkg.go.dev/std)

### 学习资源
- [The Go Tour](https://tour.golang.org/) - 交互式教程
- [Go by Example](https://gobyexample.com/) - 代码示例
- [Effective Go](https://golang.org/doc/effective_go) - 最佳实践

### Web 框架
- [Gin](https://gin-gonic.com/) - 轻量级 Web 框架
- [Echo](https://echo.labstack.com/) - 高性能 Web 框架

### 数据库
- [GORM](https://gorm.io/) - Go ORM 库
- [Ent](https://entgo.io/) - 实体框架

### 工具和插件
- [VS Code Go 扩展](https://marketplace.visualstudio.com/items?itemName=golang.Go)
- [GoLand IDE](https://www.jetbrains.com/go/)

---

## 💬 常见问题

**Q1：为什么代码文件都是 `main` 包？**
A：为了方便运行每个独立的示例。在实际项目中，应该创建多个包。

**Q2：如何导入这些文件中的函数到其他项目？**
A：将函数定义改为大写首字母（如 `Add` 而不是 `add`），然后在其他文件中导入。

**Q3：05_并发编程.go 运行时间为什么很长？**
A：该文件包含许多模拟并发操作，使用了 `time.Sleep()` 延迟，需要等待几秒。

**Q4：如何在 VS Code 中直接运行？**
A：安装 Code Runner 扩展后，右键选择 "Run Code"，或按 `Ctrl+Alt+N`。

**Q5：如何调试这些代码？**
A：安装 Delve 调试工具：`go install github.com/go-delve/delve/cmd/dlv@latest`

---

**祝你学习愉快！有问题欢迎参考官方文档或社区论坛！** 🎉
