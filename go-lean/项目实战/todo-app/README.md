# Todo 应用 - 前后端项目

一个完整的全栈 Todo 应用示例，使用 Go 后端和 HTML5 + JavaScript 前端，实现了数据库的完整 CRUD 操作。

## 📁 项目结构

```
todo-app/
├── backend/                    # Go 后端
│   ├── main.go                # 主程序（REST API）
│   ├── go.mod                 # Go 模块配置
│   └── todos.db               # SQLite 数据库（运行后自动创建）
├── frontend/                   # 前端
│   └── index.html             # 单页面应用（SPA）
└── README.md                  # 此文档
```

## 🎯 功能概述

### 后端功能（Go + SQLite）

**CRUD 操作完整实现**：

| 操作 | HTTP 方法 | 端点 | 说明 |
|------|---------|------|------|
| Create | POST | `/api/todos` | 创建新的待办事项 |
| Read (All) | GET | `/api/todos` | 获取所有待办事项 |
| Read (One) | GET | `/api/todos/detail?id=N` | 获取单个待办事项 |
| Update | PUT | `/api/todos/update?id=N` | 更新待办事项 |
| Delete | DELETE | `/api/todos/delete?id=N` | 删除单个待办事项 |
| Delete (Batch) | DELETE | `/api/todos` | 删除所有已完成的任务 |
| Toggle | POST | `/api/todos/toggle?id=N` | 切换完成状态 |

### 前端功能

- ✅ 实时列表展示
- ✅ 添加待办事项（标题+描述）
- ✅ 标记完成/未完成
- ✅ 删除单个任务
- ✅ 批量清空已完成
- ✅ 统计信息（总计、已完成、待完成）
- ✅ 错误提示
- ✅ 响应式设计

## 🚀 快速开始

### 前置要求

- **Go 1.18+**
- **任何现代浏览器**（Chrome、Firefox、Safari、Edge）
- **SQLite3**（可选，Go 驱动已包含）

### 后端启动

#### 1. 安装依赖

```bash
cd /workspaces/cs-lean/go-lean/项目实战/todo-app/backend

# 下载 SQLite 驱动
go get github.com/mattn/go-sqlite3
```

#### 2. 运行服务器

```bash
go run main.go
```

**预期输出**：
```
Database initialized successfully
Server starting on http://localhost:8080
API Documentation:
  GET    /api/todos              - Get all todos
  GET    /api/todos/detail?id=N - Get todo by ID
  POST   /api/todos              - Create todo
  PUT    /api/todos/update?id=N - Update todo
  DELETE /api/todos/delete?id=N - Delete todo
  POST   /api/todos/toggle?id=N - Toggle todo status
  DELETE /api/todos              - Delete all done todos
```

### 前端启动

#### 方式 1：使用浏览器打开文件

```bash
# 在浏览器中打开
open /workspaces/cs-lean/go-lean/项目实战/todo-app/frontend/index.html
```

或者按 `Ctrl+L` 输入：
```
file:///workspaces/cs-lean/go-lean/项目实战/todo-app/frontend/index.html
```

#### 方式 2：使用 Python 简单服务器（推荐）

```bash
cd /workspaces/cs-lean/go-lean/项目实战/todo-app/frontend

# Python 3
python3 -m http.server 3000

# 然后访问 http://localhost:3000
```

#### 方式 3：使用 Go 简单服务器

在 backend 目录创建 `server.go`：
```go
package main

import (
    "net/http"
)

func serveStatic(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "../frontend/index.html")
}

// 在 main() 中添加
mux.HandleFunc("/", serveStatic)
```

## 📝 API 详细文档

### 1. 获取所有待办事项

**请求**：
```bash
GET /api/todos
```

**响应示例**：
```json
{
  "code": 0,
  "message": "Success",
  "data": [
    {
      "id": 1,
      "title": "学习 Go",
      "desc": "完成基础语法课程",
      "done": false
    },
    {
      "id": 2,
      "title": "完成项目",
      "desc": "实现 Todo 应用",
      "done": true
    }
  ]
}
```

### 2. 创建待办事项

**请求**：
```bash
POST /api/todos
Content-Type: application/json

{
  "title": "学习 Go",
  "desc": "完成基础语法课程",
  "done": false
}
```

**响应**：
```json
{
  "code": 0,
  "message": "Todo created successfully",
  "data": {
    "id": 1,
    "title": "学习 Go",
    "desc": "完成基础语法课程",
    "done": false
  }
}
```

### 3. 获取单个待办事项

**请求**：
```bash
GET /api/todos/detail?id=1
```

**响应**：
```json
{
  "code": 0,
  "message": "Success",
  "data": {
    "id": 1,
    "title": "学习 Go",
    "desc": "完成基础语法课程",
    "done": false
  }
}
```

### 4. 更新待办事项

**请求**：
```bash
PUT /api/todos/update?id=1
Content-Type: application/json

{
  "title": "深入学习 Go",
  "desc": "完成高级特性课程",
  "done": true
}
```

**响应**：
```json
{
  "code": 0,
  "message": "Todo updated successfully",
  "data": {
    "id": 1,
    "title": "深入学习 Go",
    "desc": "完成高级特性课程",
    "done": true
  }
}
```

### 5. 删除单个待办事项

**请求**：
```bash
DELETE /api/todos/delete?id=1
```

**响应**：
```json
{
  "code": 0,
  "message": "Todo deleted successfully",
  "data": {
    "id": 1
  }
}
```

### 6. 切换完成状态

**请求**：
```bash
POST /api/todos/toggle?id=1
```

**响应**：
```json
{
  "code": 0,
  "message": "Todo toggled",
  "data": {
    "id": 1,
    "done": true
  }
}
```

### 7. 清空已完成任务

**请求**：
```bash
DELETE /api/todos
```

**响应**：
```json
{
  "code": 0,
  "message": "Done todos deleted",
  "data": {
    "deleted": 3
  }
}
```

## 💻 后端代码分析

### 数据库初始化

```go
func initDB() error {
    // 打开 SQLite 数据库连接
    db, err := sql.Open("sqlite3", "todos.db")
    
    // 创建 todos 表
    createTableSQL := `
    CREATE TABLE IF NOT EXISTS todos (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        title TEXT NOT NULL,
        desc TEXT,
        done BOOLEAN DEFAULT 0,
        created_at DATETIME DEFAULT CURRENT_TIMESTAMP
    );
    `
    
    _, err = db.Exec(createTableSQL)
    return err
}
```

### 数据模型

```go
type Todo struct {
    ID    int    `json:"id"`      // 待办事项 ID
    Title string `json:"title"`   // 标题
    Desc  string `json:"desc"`    // 描述
    Done  bool   `json:"done"`    // 是否完成
}

type Response struct {
    Code    int         `json:"code"`    // 状态码（0 成功）
    Message string      `json:"message"` // 消息
    Data    interface{} `json:"data"`    // 数据
}
```

### 中间件示例

**CORS 中间件**（允许跨域请求）：
```go
func corsMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        next.ServeHTTP(w, r)
    })
}
```

**日志中间件**：
```go
func loggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Printf("[%s] %s %s", r.Method, r.URL.Path, r.RemoteAddr)
        next.ServeHTTP(w, r)
    })
}
```

### CRUD 实现示例

**创建（Create）**：
```go
func createTodo(w http.ResponseWriter, r *http.Request) {
    var todo Todo
    json.NewDecoder(r.Body).Decode(&todo)
    
    result, _ := db.Exec(
        "INSERT INTO todos (title, desc, done) VALUES (?, ?, ?)",
        todo.Title, todo.Desc, todo.Done,
    )
    
    id, _ := result.LastInsertId()
    todo.ID = int(id)
    sendJSON(w, 0, "Success", todo)
}
```

**读取（Read）**：
```go
func getTodos(w http.ResponseWriter, r *http.Request) {
    rows, _ := db.Query("SELECT id, title, desc, done FROM todos ORDER BY id DESC")
    
    todos := []Todo{}
    for rows.Next() {
        var todo Todo
        rows.Scan(&todo.ID, &todo.Title, &todo.Desc, &todo.Done)
        todos = append(todos, todo)
    }
    
    sendJSON(w, 0, "Success", todos)
}
```

**更新（Update）**：
```go
func updateTodo(w http.ResponseWriter, r *http.Request) {
    id := /* 从请求获取 */
    var todo Todo
    json.NewDecoder(r.Body).Decode(&todo)
    
    _, _ = db.Exec(
        "UPDATE todos SET title = ?, desc = ?, done = ? WHERE id = ?",
        todo.Title, todo.Desc, todo.Done, id,
    )
    
    sendJSON(w, 0, "Success", todo)
}
```

**删除（Delete）**：
```go
func deleteTodo(w http.ResponseWriter, r *http.Request) {
    id := /* 从请求获取 */
    
    _, _ = db.Exec("DELETE FROM todos WHERE id = ?", id)
    
    sendJSON(w, 0, "Success", map[string]int{"id": id})
}
```

## 🎨 前端代码分析

### 获取待办事项

```javascript
async function loadTodos() {
    try {
        const response = await fetch('http://localhost:8080/api/todos');
        const result = await response.json();
        
        if (result.code === 0) {
            const todos = result.data || [];
            renderTodos(todos);
            updateStats(todos);
        }
    } catch (error) {
        showError('加载失败');
    }
}
```

### 添加待办事项

```javascript
async function addTodo() {
    const title = document.getElementById('todoInput').value.trim();
    
    const response = await fetch('http://localhost:8080/api/todos', {
        method: 'POST',
        headers: {'Content-Type': 'application/json'},
        body: JSON.stringify({
            title: title,
            desc: desc,
            done: false
        })
    });
    
    const result = await response.json();
    if (result.code === 0) {
        loadTodos();  // 刷新列表
    }
}
```

### 切换完成状态

```javascript
async function toggleTodo(id) {
    const response = await fetch(`http://localhost:8080/api/todos/toggle?id=${id}`, {
        method: 'POST'
    });
    
    const result = await response.json();
    if (result.code === 0) {
        loadTodos();  // 刷新列表
    }
}
```

## 🔧 测试 API

### 使用 curl 测试

```bash
# 获取所有待办事项
curl http://localhost:8080/api/todos

# 创建待办事项
curl -X POST http://localhost:8080/api/todos \
  -H "Content-Type: application/json" \
  -d '{"title":"学习Go","desc":"完成基础课程","done":false}'

# 获取单个待办事项
curl http://localhost:8080/api/todos/detail?id=1

# 更新待办事项
curl -X PUT http://localhost:8080/api/todos/update?id=1 \
  -H "Content-Type: application/json" \
  -d '{"title":"深入学习Go","desc":"完成高级课程","done":true}'

# 切换完成状态
curl -X POST http://localhost:8080/api/todos/toggle?id=1

# 删除待办事项
curl -X DELETE http://localhost:8080/api/todos/delete?id=1

# 清空已完成任务
curl -X DELETE http://localhost:8080/api/todos
```

### 使用 Postman 测试

1. 下载并安装 [Postman](https://www.postman.com/downloads/)
2. 创建新的请求集合
3. 添加上述 API 端点
4. 逐个测试每个端点

### 使用 VS Code REST Client 测试

创建 `test.http` 文件：
```http
### 获取所有待办事项
GET http://localhost:8080/api/todos

### 创建待办事项
POST http://localhost:8080/api/todos
Content-Type: application/json

{
  "title": "学习 Go",
  "desc": "完成基础课程",
  "done": false
}

### 切换完成状态
POST http://localhost:8080/api/todos/toggle?id=1

### 删除待办事项
DELETE http://localhost:8080/api/todos/delete?id=1
```

然后安装 VS Code 扩展 "REST Client" 并运行。

## 🎓 学习重点

### 后端学习内容

1. **HTTP 服务器**
   - `http.HandleFunc()` 注册路由
   - `http.ListenAndServe()` 启动服务器
   - `http.MethodGet`, `http.MethodPost` 等常量

2. **REST API 设计**
   - GET 获取资源
   - POST 创建资源
   - PUT 更新资源
   - DELETE 删除资源

3. **数据库操作**
   - 使用 `database/sql` 包
   - 执行 SQL 语句（INSERT, SELECT, UPDATE, DELETE）
   - 处理查询结果行

4. **JSON 处理**
   - `json.Unmarshal()` 解析 JSON
   - `json.Marshal()` 生成 JSON
   - 结构体标签 `json:"field"`

5. **中间件模式**
   - 函数装饰器
   - CORS 跨域处理
   - 日志记录

6. **错误处理**
   - 检查 `error` 返回值
   - 返回有意义的错误消息

### 前端学习内容

1. **异步编程**
   - `async/await` 语法
   - `fetch()` API

2. **DOM 操作**
   - `document.getElementById()`
   - `innerHTML` 动态渲染
   - 事件监听

3. **CSS 样式**
   - Flexbox 布局
   - 响应式设计
   - 动画和过渡

4. **JavaScript 高级**
   - 高阶函数
   - 数组方法（map, filter 等）
   - 字符串转义防止 XSS

## 📚 进阶优化建议

### 后端优化

1. **添加身份验证**
   ```go
   // 添加 JWT token 验证
   func authMiddleware(next http.Handler) http.Handler {
       // 验证 Authorization header
   }
   ```

2. **分页查询**
   ```go
   func getTodos(w http.ResponseWriter, r *http.Request) {
       limit := r.URL.Query().Get("limit")    // 每页数量
       offset := r.URL.Query().Get("offset")  // 偏移量
   }
   ```

3. **数据验证**
   ```go
   // 使用 validator 包进行输入验证
   if len(todo.Title) > 200 {
       sendError(w, 400, "Title too long")
   }
   ```

4. **使用 ORM 框架**
   ```go
   // 使用 GORM 替代原生 SQL
   type Todo struct {
       gorm.Model
       Title string
       Desc  string
       Done  bool
   }
   db.Create(&todo)
   ```

### 前端优化

1. **添加编辑功能**
   ```javascript
   async function editTodo(id, newTitle) {
       // PUT 请求更新
   }
   ```

2. **本地存储缓存**
   ```javascript
   localStorage.setItem('todos', JSON.stringify(todos));
   ```

3. **使用框架**
   ```javascript
   // 使用 Vue.js, React 或 Angular
   // 提高代码组织性和可维护性
   ```

4. **搜索和过滤**
   ```javascript
   function filterTodos(keyword) {
       return todos.filter(t => t.title.includes(keyword));
   }
   ```

## 🐛 常见问题排查

### 问题 1：CORS 错误

**错误**：`Access to XMLHttpRequest at 'http://localhost:8080' from origin 'file://' has been blocked`

**原因**：浏览器的同源策略限制

**解决方案**：
- 使用 Python/Go 简单服务器而不是 `file://` 协议
- 或者使用浏览器扩展禁用 CORS（仅开发用）

### 问题 2：数据库锁定

**错误**：`database is locked`

**原因**：SQLite 在多进程访问时可能出现

**解决方案**：
- 检查是否有多个程序访问同一数据库
- 使用 MySQL/PostgreSQL 替代 SQLite
- 增加连接超时时间

### 问题 3：JSON 解析错误

**错误**：`unexpected end of JSON input`

**原因**：发送的请求体格式不正确

**解决方案**：
- 检查 Content-Type 是否为 `application/json`
- 验证 JSON 格式（使用 JSONLint）

### 问题 4：404 Not Found

**错误**：`GET http://localhost:8080/api/todos 404`

**原因**：后端服务未运行或路由注册错误

**解决方案**：
- 确认后端已启动（`go run main.go`）
- 检查路由路径拼写
- 查看后端日志输出

## 📊 数据库架构

### todos 表结构

```sql
CREATE TABLE todos (
    id INTEGER PRIMARY KEY AUTOINCREMENT,  -- 自增主键
    title TEXT NOT NULL,                   -- 标题（必填）
    desc TEXT,                             -- 描述（可选）
    done BOOLEAN DEFAULT 0,                -- 是否完成（默认否）
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP  -- 创建时间
);
```

### 示例数据

```sql
INSERT INTO todos (title, desc, done) VALUES
('学习 Go', '完成基础语法课程', 0),
('完成项目', '实现 Todo 应用', 1),
('代码审查', '检查所有 API 端点', 0);
```

### 查询示例

```sql
-- 获取所有未完成的任务
SELECT * FROM todos WHERE done = 0 ORDER BY id DESC;

-- 统计已完成的任务数
SELECT COUNT(*) FROM todos WHERE done = 1;

-- 搜索标题包含关键字的任务
SELECT * FROM todos WHERE title LIKE '%Go%';

-- 删除所有已完成的任务
DELETE FROM todos WHERE done = 1;
```

## 🔐 安全建议

1. **SQL 注入防护**
   - ✅ 本项目使用参数化查询（使用 `?` 占位符）
   - ❌ 不要使用字符串拼接构建 SQL

2. **XSS 防护**
   - ✅ 前端使用 `escapeHtml()` 函数
   - ❌ 不要使用 `innerHTML` 直接插入用户数据

3. **CORS 配置**
   - 生产环境中限制 `Access-Control-Allow-Origin`
   - 不要使用 `*` 允许所有域名

4. **输入验证**
   - 验证字段长度
   - 验证数据类型
   - 验证业务规则

## 📈 性能优化

1. **数据库索引**
   ```sql
   CREATE INDEX idx_done ON todos(done);
   ```

2. **连接池配置**
   ```go
   db.SetMaxOpenConns(25)
   db.SetMaxIdleConns(5)
   ```

3. **缓存机制**
   ```go
   // 使用内存缓存减少数据库查询
   cache := make(map[int]Todo)
   ```

4. **异步处理**
   ```go
   go func() {
       // 后台任务
   }()
   ```

## 🚀 部署到生产环境

### 使用 Docker

创建 `Dockerfile`：
```dockerfile
FROM golang:1.21

WORKDIR /app
COPY . .
RUN go build -o app main.go

EXPOSE 8080
CMD ["./app"]
```

构建和运行：
```bash
docker build -t todo-app .
docker run -p 8080:8080 todo-app
```

### 使用云服务

- **Heroku**：支持 Go 和 SQLite
- **Railway**：现代云平台
- **Replit**：快速部署
- **AWS Lambda**：无服务器架构

## 📖 参考资源

- [Go 标准库 - net/http](https://pkg.go.dev/net/http)
- [SQLite 文档](https://www.sqlite.org/docs.html)
- [RESTful API 设计](https://restfulapi.net/)
- [MDN - Fetch API](https://developer.mozilla.org/en-US/docs/Web/API/Fetch_API)

---

**项目完成！现在你已经掌握了 Go 全栈开发的基本技能。** 🎉

建议的后续学习：
- [ ] 添加用户认证功能
- [ ] 实现标签分类功能
- [ ] 添加提醒功能
- [ ] 改进 UI/UX 设计
- [ ] 使用前端框架（React/Vue）重构
- [ ] 部署到生产环境
