# 项目实战

本文件夹包含完整的 Go 全栈项目示例，展示如何使用 Go 构建实际的生产级应用。

## 📂 项目列表

### 1. [Todo 应用](./todo-app/README.md)

一个完整的待办事项管理应用，演示前后端交互和数据库 CRUD 操作。

**技术栈**：
- 后端：Go + SQLite + REST API
- 前端：HTML5 + CSS3 + JavaScript
- 特性：完整的 CRUD 操作、CORS 支持、错误处理

**文件结构**：
```
todo-app/
├── backend/
│   ├── main.go           # 完整的 REST API 服务器
│   └── go.mod            # Go 模块配置
├── frontend/
│   └── index.html        # 单页面应用，无依赖
└── README.md             # 详细文档
```

**核心功能**：
- ✅ 创建、读取、更新、删除待办事项
- ✅ 标记任务完成/未完成
- ✅ 批量清空已完成任务
- ✅ 实时统计信息
- ✅ 错误提示和验证

**学习价值**：
- HTTP 服务器和路由设计
- REST API 设计原则
- SQLite 数据库操作
- JSON 序列化和反序列化
- CORS 跨域资源共享
- 中间件模式
- 前后端交互
- 异步编程（async/await）

**启动指南**：

1. **启动后端**：
   ```bash
   cd todo-app/backend
   go get github.com/mattn/go-sqlite3
   go run main.go
   ```

2. **启动前端**：
   ```bash
   cd todo-app/frontend
   python3 -m http.server 3000
   # 访问 http://localhost:3000
   ```

**测试 API**：
```bash
# 获取所有待办事项
curl http://localhost:8080/api/todos

# 创建待办事项
curl -X POST http://localhost:8080/api/todos \
  -H "Content-Type: application/json" \
  -d '{"title":"学习Go","desc":"","done":false}'

# 完整的 API 文档见 todo-app/README.md
```

---

### 2. [Blog API](./blog-api/README.md)

使用 Gin 框架和 GORM ORM 构建的现代博客系统，演示如何使用流行框架提高开发效率。

**技术栈**：
- 后端：Go + Gin + GORM + SQLite
- 前端：Vue.js 3 + Axios
- 特性：框架化开发、分页、搜索、统计、CORS

**文件结构**：
```
blog-api/
├── backend/
│   ├── main.go           # Gin + GORM 完整应用
│   └── go.mod            # Go 模块配置
├── frontend/
│   ├── index.html        # Vue.js 3 单页应用
│   └── README.md         # 前端文档
└── README.md             # 详细文档
```

**核心功能**：
- ✅ 文章列表（分页）
- ✅ 创建、编辑、删除文章
- ✅ 全文搜索功能
- ✅ 按分类过滤
- ✅ 浏览次数统计
- ✅ 统计信息面板

**学习价值**：
- Gin Web 框架使用
- GORM ORM 数据库操作
- 自动迁移数据库
- Vue.js 3 组合式 API
- 现代 Web 开发流程
- 前端框架集成
- 分页和搜索实现

**启动指南**：

1. **启动后端**：
   ```bash
   cd blog-api/backend
   go mod download
   go run main.go
   # 访问 http://localhost:8080
   ```

2. **启动前端**：
   ```bash
   cd blog-api/frontend
   python3 -m http.server 8000
   # 访问 http://localhost:8000
   # 或直接在浏览器打开 index.html
   ```

**测试 API**：
```bash
# 获取文章列表
curl http://localhost:8080/api/articles?page=1&limit=10

# 创建文章
curl -X POST http://localhost:8080/api/articles \
  -H "Content-Type: application/json" \
  -d '{"title":"Go教程","content":"内容","author":"张三","category":"Go"}'

# 搜索文章
curl http://localhost:8080/api/search?q=Go

# 完整的 API 文档见 blog-api/README.md
```

---

## 🗂️ 项目规划

本文件夹后续会添加更多项目示例：

### 计划中的项目

| 项目 | 描述 | 技术栈 | 难度 | 状态 |
|------|------|--------|------|------|
| Todo App | 待办事项管理 | Go + SQLite | 初级 | ✅ 完成 |
| Blog API | RESTful 博客系统 | Go + Gin + GORM | 中级 | ✅ 完成 |
| 聊天应用 | 实时聊天系统 | Go + WebSocket | 中级 | 📋 计划中 |
| 微服务架构 | 分布式系统示例 | Go + gRPC + Kubernetes | 高级 | 📋 计划中 |
| CLI 工具 | 命令行工具 | Go + Cobra | 初级 | 📋 计划中 |
| Web 爬虫 | 网页爬虫示例 | Go + colly | 中级 | 📋 计划中 |

---

## 📚 项目学习顺序

### 初级开发者（新手）

1. 📖 先阅读 `../代码示例/` 中的基础文件（01-11）
2. 🚀 运行 **Todo App** 项目（原生 Go）
3. 🚀 运行 **Blog API** 项目（框架开发）
4. 📝 理解两个项目的区别
5. 💻 尝试修改代码（如添加新字段）
6. 🔧 自己实现新功能

### 中级开发者

1. ✅ 完成上述初级学习
2. 🎯 在两个项目基础上扩展
   - 添加用户认证
   - 添加标签/分类功能
   - 实现搜索和高级过滤
   - 添加评论功能
3. 📚 学习更多 Go 框架（Echo、Beego）
4. 🚀 部署到云服务

### 高级开发者

1. 👨‍💻 设计自己的项目架构
2. 📊 考虑可扩展性和性能
3. 🔐 实现安全最佳实践
4. 🚀 使用高级特性（并发、缓存等）
5. 📈 实现监控和日志系统

---

## 🎯 项目开发流程

每个项目都遵循以下开发流程：

### 1. 需求分析
- 明确功能需求
- 设计数据模型
- 规划 API 接口

### 2. 后端开发
- 设计数据库架构
- 实现 CRUD 操作
- 添加业务逻辑
- 实现错误处理

### 3. 前端开发
- 设计用户界面
- 实现交互逻辑
- 调用后端 API
- 处理错误和加载状态

### 4. 测试
- 单元测试
- 集成测试
- 手动测试

### 5. 部署
- 本地测试
- 云服务部署
- 性能优化

---

## 💡 项目实战的价值

### 学习真实开发流程
- 从需求到实现
- 从设计到部署
- 从开发到优化

### 掌握最佳实践
- 代码组织
- 错误处理
- 安全考虑
- 性能优化

### 积累项目经验
- 实践各种技术
- 解决真实问题
- 构建作品集

---

## 🚀 快速开始

### 步骤 1：选择项目

建议首先学习 **Todo App** 项目（难度最低，功能完整）。

### 步骤 2：阅读文档

每个项目都有详细的 `README.md` 文档：
- 项目结构说明
- 功能清单
- 快速开始指南
- API 文档
- 代码分析
- 常见问题

### 步骤 3：运行项目

按照各项目的快速开始指南运行。

### 步骤 4：修改代码

- 修改数据模型
- 添加新的 API 端点
- 改进 UI 设计
- 实现新功能

### 步骤 5：扩展项目

- 添加认证功能
- 集成第三方服务
- 部署到生产环境

---

## 📖 学习资源

### Go 相关

- [Go 官方文档](https://golang.org/doc/)
- [Go 标准库](https://pkg.go.dev/std)
- [Effective Go](https://golang.org/doc/effective_go)

### 数据库

- [SQLite 文档](https://www.sqlite.org/docs.html)
- [SQL 教程](https://www.w3schools.com/sql/)
- [GORM 文档](https://gorm.io/)

### Web 开发

- [HTTP 协议](https://developer.mozilla.org/en-US/docs/Web/HTTP)
- [REST API 设计](https://restfulapi.net/)
- [MDN Web 文档](https://developer.mozilla.org/)

### 前端

- [JavaScript 教程](https://developer.mozilla.org/en-US/docs/Web/JavaScript)
- [CSS 参考](https://developer.mozilla.org/en-US/docs/Web/CSS)
- [HTML 文档](https://developer.mozilla.org/en-US/docs/Web/HTML)

---

## 🤝 贡献指南

如果你想为项目实战添加新项目：

1. 在此文件夹创建新目录
2. 实现项目的后端和前端代码
3. 编写详细的 `README.md` 文档
4. 确保代码可以直接运行
5. 提交 Pull Request

---

## 📊 项目对比

### 技术栈对比

| 维度 | Todo App | Blog API |
|------|---------|----------|
| 框架 | 原生 net/http | Gin |
| ORM | database/sql | GORM |
| 数据库 | SQLite | SQLite |
| 前端 | 原生 HTML/JS | Vue.js 3 |
| API 端点 | 7 个 | 8 个 |
| 自动迁移 | ❌ 无 | ✅ 有 |
| 中间件 | 手写 | 内置 |
| 分页 | 手写逻辑 | 内置支持 |
| 难度 | ⭐ 简单 | ⭐⭐ 中等 |

### 学习路径对比

```
Todo App (原生开发)
  ↓
理解 HTTP、数据库基础
  ↓
Blog API (框架开发)
  ↓
理解框架的便利性和抽象
  ↓
能够选择和使用合适的框架
```

### 项目特性对比

| 特性 | Todo App | Blog API |
|------|---------|----------|
| 列表展示 | ✅ | ✅ |
| 分页 | ❌ | ✅ |
| 搜索 | ❌ | ✅ |
| 分类过滤 | ❌ | ✅ |
| 统计信息 | ✅ | ✅ |
| 浏览计数 | ❌ | ✅ |
| CORS 支持 | ✅ | ✅ |
| 错误处理 | ✅ | ✅ |

---

## ❓ 常见问题

**Q: 项目能直接用于生产环境吗？**

A: 这些项目主要用于学习。生产部署需要：
- 添加认证和授权
- 实现更强的错误处理
- 添加日志和监控
- 进行性能优化
- 实现数据备份

**Q: 可以修改项目代码吗？**

A: 完全可以！建议：
- 修改源代码理解原理
- 添加新功能练习
- 重构代码提高质量

**Q: 如何部署到云服务？**

A: 每个项目的文档都有部署说明。常用方案：
- Docker 容器化
- Heroku、Railway 等平台
- AWS、Google Cloud 等云服务

---

## 📝 更新历史

| 日期 | 项目 | 更新 |
|------|------|------|
| 2025-01-15 | Blog API | 创建 Blog API 项目（Gin + GORM + Vue.js） |
| 2025-01-14 | Todo App | 创建 Todo App 项目（原生 Go + Vue.js） |
| 2025-01-13 | 初始化 | 创建项目实战目录 |

---

## 🎯 核心对比：从原生到框架

### Todo App - 原生开发

```go
// ❌ 手动处理 HTTP 方法
func handleTodos(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodGet {
        // 手写 GET 逻辑
    } else if r.Method == http.MethodPost {
        // 手写 POST 逻辑
    }
}

// ❌ 手写 SQL 和 Scan
rows, err := db.Query("SELECT id, title, done FROM todos")
for rows.Next() {
    rows.Scan(&todo.ID, &todo.Title, &todo.Done)
}
```

**优点**: 理解底层原理

**缺点**: 代码重复多，开发效率低

### Blog API - 框架开发

```go
// ✅ 清晰的路由定义
router.GET("/api/articles", GetArticles)
router.POST("/api/articles", CreateArticle)

// ✅ ORM 自动映射
db.Limit(limit).Offset(offset).Find(&articles)
```

**优点**: 高效开发，代码简洁

**缺点**: 隐藏了底层细节

---

## 💼 项目实战总结

通过这两个项目，你将学到：

### Todo App (初级)
- ✅ HTTP 服务器基础
- ✅ REST API 设计原则
- ✅ 数据库基本操作 (CRUD)
- ✅ 前后端交互
- ✅ 错误处理
- ✅ 中间件模式

### Blog API (中级)
- ✅ Web 框架的优势
- ✅ ORM 的使用
- ✅ 高级查询 (分页、搜索)
- ✅ 现代前端框架 (Vue.js)
- ✅ 开发工作流优化
- ✅ 生产就绪的代码

### 从原生到框架

通过对比两个项目，你将理解：
- **为什么使用框架**：提高开发效率
- **框架的成本**：学习曲线，增加依赖
- **何时选择原生**：性能关键、简单项目
- **何时选择框架**：复杂项目、团队协作

---

## 🚀 下一步行动

1. **运行两个项目**
   - 先运行 Todo App，理解基础
   - 再运行 Blog API，对比差异

2. **修改和扩展**
   - 添加新的字段
   - 实现新的功能
   - 优化用户界面

3. **学习更多框架**
   - Echo: 类似 Gin 的轻量框架
   - Beego: 全功能的 Web 框架
   - Buffalo: Ruby on Rails 风格

4. **进阶任务**
   - 添加用户认证
   - 部署到云服务
   - 性能优化
   - 编写测试

---

## 📚 相关资源

### 官方文档
- [Gin 框架](https://gin-gonic.com/)
- [GORM ORM](https://gorm.io/zh_CN/)
- [Go 标准库](https://pkg.go.dev/std)

### 学习资源
- [Go Web 开发最佳实践](https://www.topgoer.com/)
- [REST API 设计指南](https://restfulapi.net/)
- [Vue.js 3 官方教程](https://vuejs.org/guide/introduction.html)

### 部署方案
- [Docker 容器化](https://www.docker.com/)
- [Heroku 平台](https://www.heroku.com/)
- [Railway 平台](https://railway.app/)

---

**每个项目都是知识的积累，持续学习，不断进步！** 🚀

````

---
