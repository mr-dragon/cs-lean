# Blog API 项目文档

本项目使用 **Gin 框架** + **GORM ORM** + **Vue.js 3** 实现一个完整的博客系统，展示现代 Go Web 开发最佳实践。

## 📚 项目特点

### 后端特性
- ✅ **Gin 框架**: 轻量级、高性能的 Web 框架
- ✅ **GORM ORM**: 功能强大的对象关系映射库，支持自动迁移
- ✅ **SQLite 数据库**: 轻量级数据库，无需额外部署
- ✅ **RESTful API**: 标准化的 API 设计
- ✅ **CORS 支持**: 跨域请求处理
- ✅ **分页功能**: 大数据量下的分页查询
- ✅ **全文搜索**: 按标题、内容、作者搜索
- ✅ **分类过滤**: 按分类过滤文章
- ✅ **浏览计数**: 统计每篇文章的浏览次数

### 前端特性
- ✅ **Vue.js 3**: 现代化的前端框架
- ✅ **Axios**: 优雅的 HTTP 客户端
- ✅ **响应式设计**: 完美适配各种设备
- ✅ **实时交互**: 无刷新更新
- ✅ **完整功能**: 列表、搜索、创建、编辑、删除

## 🏗️ 项目结构

```
blog-api/
├── backend/
│   ├── main.go              # 后端服务（Gin + GORM）
│   ├── go.mod              # Go 模块配置
│   └── blog.db             # SQLite 数据库（自动创建）
├── frontend/
│   ├── index.html          # Vue.js 单页应用
│   └── README.md           # 前端文档
└── README.md               # 项目文档（本文件）
```

## 🔧 技术栈对比

### 与 Java 对比
| 功能 | Java | Go | 优势 |
|------|------|-----|------|
| Web 框架 | Spring Boot | Gin | Go 内存占用更低，启动更快 |
| ORM | Hibernate/JPA | GORM | GORM API 更简洁 |
| 并发模型 | 线程池 | Goroutines | Goroutines 轻量级，数量可达百万 |
| 编译产物 | Jar (100+MB) | 二进制 (10-20MB) | Go 二进制小，无需运行时 |
| 部署复杂度 | 需要 JVM | 一个可执行文件 | Go 部署简单 |

### 与项目 1 (Todo App) 对比
| 特性 | Todo App | Blog API | 区别 |
|------|---------|----------|------|
| 框架 | 原生 net/http | Gin 框架 | Gin 提供更多中间件、路由管理 |
| ORM | database/sql | GORM | GORM 提供更高层抽象 |
| 自动迁移 | 手动 SQL | GORM 自动 | 开发效率高 |
| API 端点 | 7 个 | 8 个 | Blog API 功能更丰富 |
| 前端框架 | 原生 JS | Vue.js 3 | Vue 组件化开发更高效 |

## 🚀 快速开始

### 1. 安装依赖

```bash
# 进入后端目录
cd go-lean/项目实战/blog-api/backend

# 下载 Go 模块依赖
go mod download
```

### 2. 启动后端服务

```bash
# 编译并运行
go run main.go

# 或先编译再运行
go build -o blog-api
./blog-api
```

输出示例：
```
2024/01/15 10:30:45 Starting Blog API server on :8080
2024/01/15 10:30:45 Database initialized successfully
```

### 3. 打开前端

在浏览器中打开：
```
file:///workspaces/cs-lean/go-lean/项目实战/blog-api/frontend/index.html
```

或启动一个本地 Web 服务器：
```bash
# 使用 Python 3
cd frontend
python3 -m http.server 8000

# 然后访问 http://localhost:8000
```

## 📡 API 端点详解

### 1. 获取文章列表 (分页)

**请求**
```
GET /api/articles?page=1&limit=10
```

**参数**
| 参数 | 类型 | 说明 |
|------|------|------|
| page | int | 页码，从 1 开始 |
| limit | int | 每页数量，默认 10 |

**响应成功 (200)**
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "articles": [
      {
        "id": 1,
        "title": "Go 并发编程",
        "content": "Goroutines 是轻量级线程...",
        "author": "张三",
        "category": "Go",
        "view_count": 42,
        "created_at": "2024-01-15T10:30:45Z",
        "updated_at": "2024-01-15T10:30:45Z"
      }
    ],
    "total": 5
  }
}
```

**响应失败 (500)**
```json
{
  "code": -1,
  "message": "数据库查询失败",
  "data": null
}
```

### 2. 获取单篇文章 (自动增加浏览数)

**请求**
```
GET /api/articles/:id
```

**参数**
| 参数 | 类型 | 说明 |
|------|------|------|
| id | uint | 文章 ID |

**功能**: 返回文章详情，同时自动将 `view_count` 加 1

**响应成功 (200)**
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "id": 1,
    "title": "Go 并发编程",
    "content": "完整的文章内容...",
    "author": "张三",
    "category": "Go",
    "view_count": 43,  // 已加 1
    "created_at": "2024-01-15T10:30:45Z",
    "updated_at": "2024-01-15T10:30:45Z"
  }
}
```

### 3. 创建文章

**请求**
```
POST /api/articles
Content-Type: application/json

{
  "title": "新文章标题",
  "content": "文章内容...",
  "author": "作者名称",
  "category": "技术分类"
}
```

**验证规则**
- `title`: 必填，最多 200 字符
- `content`: 必填，最多 5000 字符
- `author`: 必填，最多 50 字符
- `category`: 可选

**响应成功 (201)**
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "id": 6,
    "title": "新文章标题",
    "content": "文章内容...",
    "author": "作者名称",
    "category": "技术分类",
    "view_count": 0,
    "created_at": "2024-01-15T11:00:00Z",
    "updated_at": "2024-01-15T11:00:00Z"
  }
}
```

**响应失败 (400)**
```json
{
  "code": 400,
  "message": "缺少必填字段: title",
  "data": null
}
```

### 4. 更新文章

**请求**
```
PUT /api/articles/:id
Content-Type: application/json

{
  "title": "更新标题",
  "content": "更新内容",
  "author": "更新作者",
  "category": "更新分类"
}
```

**说明**
- 所有字段都是可选的，只更新提供的字段
- 保留未提供的字段值不变

**响应成功 (200)**
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "id": 1,
    "title": "更新标题",
    "content": "更新内容",
    "author": "更新作者",
    "category": "更新分类",
    "view_count": 43,
    "created_at": "2024-01-15T10:30:45Z",
    "updated_at": "2024-01-15T11:05:00Z"
  }
}
```

### 5. 删除文章

**请求**
```
DELETE /api/articles/:id
```

**响应成功 (200)**
```json
{
  "code": 0,
  "message": "success",
  "data": null
}
```

**响应失败 (404)**
```json
{
  "code": 404,
  "message": "文章不存在",
  "data": null
}
```

### 6. 搜索文章

**请求**
```
GET /api/search?q=Go+并发
```

**参数**
| 参数 | 类型 | 说明 |
|------|------|------|
| q | string | 搜索关键词（在 title/content/author 中模糊搜索） |

**响应成功 (200)**
```json
{
  "code": 0,
  "message": "success",
  "data": [
    {
      "id": 1,
      "title": "Go 并发编程",
      "content": "...",
      "author": "张三",
      "category": "Go",
      "view_count": 43,
      "created_at": "2024-01-15T10:30:45Z",
      "updated_at": "2024-01-15T10:30:45Z"
    }
  ]
}
```

### 7. 按分类过滤文章

**请求**
```
GET /api/category/Go
```

**参数**
| 参数 | 类型 | 说明 |
|------|------|------|
| category | string | 分类名称 |

**响应成功 (200)**
```json
{
  "code": 0,
  "message": "success",
  "data": [
    {
      "id": 1,
      "title": "Go 并发编程",
      "content": "...",
      "author": "张三",
      "category": "Go",
      "view_count": 43,
      "created_at": "2024-01-15T10:30:45Z",
      "updated_at": "2024-01-15T10:30:45Z"
    },
    {
      "id": 3,
      "title": "Go Web 开发",
      "content": "...",
      "author": "李四",
      "category": "Go",
      "view_count": 28,
      "created_at": "2024-01-15T10:35:00Z",
      "updated_at": "2024-01-15T10:35:00Z"
    }
  ]
}
```

### 8. 获取统计信息

**请求**
```
GET /api/stats
```

**响应成功 (200)**
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "total_articles": 15,
    "total_views": 428,
    "categories": 4,
    "authors": 5
  }
}
```

## 💻 后端代码分析

### 项目初始化流程

```go
func main() {
    // 1. 创建 Gin 引擎
    router := gin.Default()

    // 2. 配置 CORS 中间件 (允许所有来源)
    router.Use(CORSMiddleware())

    // 3. 初始化数据库
    initDB()

    // 4. 自动迁移数据模型
    db.AutoMigrate(&Article{})

    // 5. 注册路由
    setupRoutes(router)

    // 6. 启动服务器
    router.Run(":8080")
}
```

### 核心数据模型

```go
type Article struct {
    ID        uint      `gorm:"primaryKey"`           // 主键
    Title     string    `binding:"required"`          // 标题（必填）
    Content   string    `binding:"required"`          // 内容（必填）
    Author    string    `binding:"required"`          // 作者（必填）
    Category  string    `gorm:"index"`                // 分类（可选，建立索引加快查询）
    ViewCount int       `gorm:"default:0"`            // 浏览次数
    CreatedAt time.Time `gorm:"autoCreateTime:milli"` // 创建时间
    UpdatedAt time.Time `gorm:"autoUpdateTime:milli"` // 更新时间
}
```

### 重要函数解析

#### 1. CORS 中间件

```go
func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
        
        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }
        
        c.Next()
    }
}
```

**解释**:
- 允许所有来源的跨域请求
- 支持 GET/POST/PUT/DELETE 方法
- 预检请求直接返回 204

#### 2. 获取文章列表（分页）

```go
func GetArticles(c *gin.Context) {
    var page, limit int
    c.BindQuery(&page)      // 绑定 page 参数
    c.BindQuery(&limit)     // 绑定 limit 参数
    
    if page < 1 { page = 1 }
    if limit < 1 { limit = 10 }
    if limit > 100 { limit = 100 }  // 防止过大请求
    
    var articles []Article
    var total int64
    
    // GORM 分页查询
    db.Offset((page - 1) * limit).
       Limit(limit).
       Find(&articles)
    
    db.Model(&Article{}).Count(&total)
    
    c.JSON(200, ResponseData{
        Code:    0,
        Message: "success",
        Data: map[string]interface{}{
            "articles": articles,
            "total":    total,
        },
    })
}
```

**关键点**:
- 使用 `Offset` 和 `Limit` 实现分页
- 分别统计总数（优化：可用 `Session` 合并查询）
- 防止恶意请求（limit 限制最大 100）

#### 3. 获取文章并增加浏览数

```go
func GetArticleByID(c *gin.Context) {
    id := c.Param("id")
    
    var article Article
    
    // 使用事务确保一致性
    result := db.Model(&article).
              Where("id = ?", id).
              Update("view_count", gorm.Expr("view_count + 1"))
    
    if result.Error != nil {
        c.JSON(500, ResponseData{
            Code:    -1,
            Message: "数据库错误",
            Data:    nil,
        })
        return
    }
    
    db.First(&article, id)
    
    c.JSON(200, ResponseData{
        Code:    0,
        Message: "success",
        Data:    article,
    })
}
```

**关键点**:
- 使用 `gorm.Expr("view_count + 1")` 原子性增加（避免竞态条件）
- 分离更新和查询操作
- 使用事务保证数据一致性

#### 4. 创建文章（验证）

```go
func CreateArticle(c *gin.Context) {
    var article Article
    
    // 绑定 JSON 并自动验证
    if err := c.ShouldBindJSON(&article); err != nil {
        c.JSON(400, ResponseData{
            Code:    400,
            Message: "缺少必填字段: " + err.Error(),
            Data:    nil,
        })
        return
    }
    
    // 手动验证
    if len(article.Title) > 200 {
        c.JSON(400, ResponseData{
            Code:    400,
            Message: "标题长度不能超过 200 字符",
            Data:    nil,
        })
        return
    }
    
    // 创建记录
    if result := db.Create(&article); result.Error != nil {
        c.JSON(500, ResponseData{
            Code:    -1,
            Message: "创建失败",
            Data:    nil,
        })
        return
    }
    
    c.JSON(201, ResponseData{
        Code:    0,
        Message: "success",
        Data:    article,
    })
}
```

**关键点**:
- 使用 `ShouldBindJSON` 自动解析和验证
- 组合结构体 tag 验证和手动验证
- 返回具体错误信息

#### 5. 搜索文章（模糊搜索）

```go
func SearchArticles(c *gin.Context) {
    keyword := c.Query("q")
    
    if keyword == "" {
        c.JSON(400, ResponseData{
            Code:    400,
            Message: "搜索关键词不能为空",
            Data:    nil,
        })
        return
    }
    
    var articles []Article
    
    // LIKE 模糊查询在 title、content、author 中
    pattern := "%" + keyword + "%"
    db.Where("title LIKE ? OR content LIKE ? OR author LIKE ?",
             pattern, pattern, pattern).
       Find(&articles)
    
    c.JSON(200, ResponseData{
        Code:    0,
        Message: "success",
        Data:    articles,
    })
}
```

**关键点**:
- 使用 LIKE 操作符模糊查询
- 参数化查询防止 SQL 注入
- 在多个字段中搜索

## 🎨 前端代码分析

### Vue.js 3 组合式 API 结构

```javascript
const { createApp, ref, computed, onMounted } = Vue;

createApp({
    setup() {
        // 1. 数据定义
        const articles = ref([]);
        const view = ref('list');
        const loading = ref(false);
        
        // 2. 计算属性
        const totalPages = computed(() => {
            return Math.ceil(totalCount.value / pageSize.value);
        });
        
        // 3. 方法定义
        const loadArticles = async () => { ... };
        const handleSearch = async () => { ... };
        
        // 4. 生命周期
        onMounted(() => {
            loadArticles();
            loadStats();
        });
        
        // 5. 返回模板使用的数据
        return { articles, view, loadArticles, ... };
    }
}).mount('#app');
```

### 关键功能实现

#### 1. 获取文章列表

```javascript
const loadArticles = async () => {
    loading.value = true;
    try {
        const response = await axios.get(`${API_BASE}/articles`, {
            params: {
                page: currentPage.value,
                limit: pageSize.value
            }
        });

        if (response.data.code === 0) {
            articles.value = response.data.data.articles || [];
            totalCount.value = response.data.data.total;
        }
    } catch (error) {
        showMessage('加载文章失败', 'error');
    } finally {
        loading.value = false;
    }
};
```

**重点**:
- 异步 API 调用
- 统一错误处理
- 数据绑定更新 UI

#### 2. 搜索功能

```javascript
const handleSearch = async () => {
    if (!searchKeyword.value.trim()) {
        currentPage.value = 1;
        loadArticles();  // 重置返回全部列表
        return;
    }

    loading.value = true;
    try {
        const response = await axios.get(`${API_BASE}/search`, {
            params: { q: searchKeyword.value }
        });

        if (response.data.code === 0) {
            articles.value = response.data.data || [];
            totalCount.value = articles.value.length;
            currentPage.value = 1;
        }
    } catch (error) {
        showMessage('搜索失败', 'error');
    } finally {
        loading.value = false;
    }
};
```

#### 3. 文章创建/编辑

```javascript
const handleSubmit = async () => {
    // 1. 验证
    if (!articleForm.value.title || !articleForm.value.content) {
        showMessage('请填写所有必填字段', 'error');
        return;
    }

    loading.value = true;
    try {
        if (editingArticle.value) {
            // 编辑 - PUT 请求
            const response = await axios.put(
                `${API_BASE}/articles/${editingArticle.value.id}`,
                articleForm.value
            );
            if (response.data.code === 0) {
                showMessage('文章更新成功');
                resetForm();
                loadArticles();
            }
        } else {
            // 创建 - POST 请求
            const response = await axios.post(
                `${API_BASE}/articles`,
                articleForm.value
            );
            if (response.data.code === 0) {
                showMessage('文章发布成功');
                resetForm();
                currentPage.value = 1;
                loadArticles();
            }
        }
    } catch (error) {
        showMessage(error.response?.data?.message || '操作失败', 'error');
    } finally {
        loading.value = false;
    }
};
```

#### 4. 删除操作

```javascript
const deleteArticle = async (id) => {
    if (!confirm('确定要删除这篇文章吗？')) return;

    loading.value = true;
    try {
        const response = await axios.delete(`${API_BASE}/articles/${id}`);
        if (response.data.code === 0) {
            showMessage('文章删除成功');
            view.value = 'list';
            loadArticles();
        }
    } catch (error) {
        showMessage('删除失败', 'error');
    } finally {
        loading.value = false;
    }
};
```

### UI 交互设计

#### 1. 响应式导航栏

```html
<div class="navbar">
    <ul class="navbar-menu">
        <li>
            <button @click="view = 'list'" :class="{ active: view === 'list' }">
                文章列表
            </button>
        </li>
        <li>
            <button @click="view = 'create'" :class="{ active: view === 'create' }">
                发布文章
            </button>
        </li>
    </ul>
</div>
```

**实现**:
- `@click` 事件监听切换视图
- `:class` 动态绑定活跃状态样式

#### 2. 分页组件

```html
<div class="pagination">
    <button @click="previousPage" :disabled="currentPage === 1">
        上一页
    </button>
    <button v-for="page in totalPages" :key="page" 
            @click="goToPage(page)"
            :class="{ active: page === currentPage }">
        {{ page }}
    </button>
    <button @click="nextPage" :disabled="currentPage >= totalPages">
        下一页
    </button>
</div>
```

**特性**:
- `v-for` 循环生成页码按钮
- 禁用不可用的前后按钮
- 高亮当前页

#### 3. 条件渲染

```html
<!-- 加载状态 -->
<div v-if="loading" class="loading">加载中...</div>

<!-- 文章列表 -->
<div v-else-if="articles.length > 0" class="articles">
    <!-- 文章卡片 -->
</div>

<!-- 空状态 -->
<div v-else class="empty-state">暂无文章</div>
```

## 🔍 Gin vs 原生 Go 对比

### 原生 Go (Todo App)

```go
func main() {
    http.HandleFunc("/api/todos", handleTodos)
    http.HandleFunc("/api/todos/toggle", handleToggleTodo)
    http.ListenAndServe(":8080", nil)
}

func handleTodos(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodGet {
        // 手动处理 GET 逻辑
    } else if r.Method == http.MethodPost {
        // 手动处理 POST 逻辑
    }
}
```

**缺点**:
- 需要手动判断 HTTP 方法
- 路由管理复杂
- 无内置中间件支持

### Gin 框架 (Blog API)

```go
func main() {
    router := gin.Default()
    
    router.GET("/api/articles", GetArticles)
    router.POST("/api/articles", CreateArticle)
    router.PUT("/api/articles/:id", UpdateArticle)
    router.DELETE("/api/articles/:id", DeleteArticle)
    
    router.Run(":8080")
}
```

**优势**:
- 清晰的 HTTP 方法和路由映射
- 内置中间件系统
- 自动 CORS、日志、异常恢复
- 路由参数自动提取和类型转换

## 🚦 GORM vs database/sql 对比

### database/sql (Todo App)

```go
rows, err := db.Query("SELECT id, title, completed FROM todos")
if err != nil {
    return nil, err
}
defer rows.Close()

for rows.Next() {
    var todo Todo
    err := rows.Scan(&todo.ID, &todo.Title, &todo.Completed)
    if err != nil {
        return nil, err
    }
    todos = append(todos, todo)
}
```

**缺点**:
- 手写 SQL，容易出错
- Scan 参数顺序必须与 SELECT 一致
- 没有类型安全检查

### GORM (Blog API)

```go
var articles []Article

db.Offset((page-1)*limit).
   Limit(limit).
   Find(&articles)
```

**优势**:
- 无需手写 SQL
- 自动类型映射
- 链式调用，代码简洁
- 支持复杂查询（Join、Preload、Transaction）
- 内置分页、搜索等常用功能

## 📊 性能优化建议

### 1. 数据库索引

```go
type Article struct {
    Category  string `gorm:"index"`           // 为常用查询字段添加索引
    CreatedAt time.Time `gorm:"index"`
}
```

### 2. 查询优化

```go
// ❌ N+1 查询问题
articles := []Article{}
db.Find(&articles)
for _, article := range articles {
    db.First(&author, article.AuthorID)  // 多次查询
}

// ✅ 使用 Preload 一次加载
db.Preload("Author").Find(&articles)
```

### 3. 缓存策略

```go
// 使用 Redis 缓存热点数据
if cachedArticle := cache.Get(articleID); cachedArticle != nil {
    return cachedArticle
}

article := db.First(...) // 从数据库查询
cache.Set(articleID, article, 24*time.Hour) // 缓存 24 小时
```

### 4. 连接池配置

```go
sqlDB, err := db.DB()
sqlDB.SetMaxIdleConns(10)        // 最大空闲连接
sqlDB.SetMaxOpenConns(100)       // 最大打开连接
sqlDB.SetConnMaxLifetime(time.Hour) // 连接最大生命周期
```

### 5. 前端优化

```javascript
// 使用虚拟滚动处理大列表
// 只渲染可见区域的元素，大幅提高性能

// 防抖搜索
const searchDebounce = debounce(handleSearch, 300);

// 请求取消
const controller = new AbortController();
axios.get(url, { signal: controller.signal });
```

## 🐛 常见问题排查

### Q: 前端请求返回 CORS 错误

**症状**: 
```
Access to XMLHttpRequest at 'http://localhost:8080/api/articles' 
from origin 'file://...' has been blocked by CORS policy
```

**解决**:
1. 确保后端启用了 CORS 中间件
2. 使用 `router.Use(CORSMiddleware())`
3. 前端通过 HTTP 服务器打开（不要用 file:// 协议）

### Q: 数据库找不到记录

**检查**:
1. 数据库文件是否存在：`ls -la blog.db`
2. 是否运行了 `AutoMigrate`：`db.AutoMigrate(&Article{})`
3. 查询语句是否正确：`db.Where("id = ?", id).First(...)`

### Q: 分页查询返回空数组

**常见原因**:
1. `Offset` 超出了总数据量
2. `Limit` 太小
3. 数据库为空

**调试**:
```go
var total int64
db.Model(&Article{}).Count(&total)
log.Printf("Total articles: %d, page: %d, limit: %d", total, page, limit)
```

## 📚 学习资源

### Go Web 开发
- [Gin 官方文档](https://github.com/gin-gonic/gin)
- [GORM 官方文档](https://gorm.io/)
- [Go 并发编程](https://go.dev/doc/effective_go#concurrency)

### 前端框架
- [Vue.js 3 官方文档](https://vuejs.org/)
- [Axios 文档](https://axios-http.com/)

### API 设计
- [RESTful API 最佳实践](https://restfulapi.net/)
- [HTTP 状态码参考](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status)

### 数据库
- [SQL 基础教程](https://www.sqlstyle.guide/)
- [SQLite 文档](https://www.sqlite.org/docs.html)

## ✅ 进阶任务

1. **添加用户认证**
   - 实现用户注册/登录
   - 使用 JWT token 认证
   - 限制普通用户只能编辑自己的文章

2. **添加评论功能**
   - 创建 Comment 模型
   - 实现评论的 CRUD 接口
   - 前端渲染评论列表

3. **实现全文搜索**
   - 集成 Elasticsearch
   - 提供高效的搜索体验
   - 支持中文分词

4. **性能监控**
   - 集成 Prometheus 监控
   - 实现请求耗时统计
   - 前端添加性能指标展示

5. **容器化部署**
   - 编写 Dockerfile
   - 创建 docker-compose.yml
   - 实现一键启动

## 📝 总结

本项目展示了 Go Web 开发的完整流程：

- **后端**: 使用 Gin 框架快速构建 RESTful API，使用 GORM 简化数据库操作
- **前端**: 使用 Vue.js 3 构建现代化交互界面，使用 Axios 管理 HTTP 请求
- **数据库**: SQLite 快速原型开发，生产环境建议迁移到 PostgreSQL/MySQL

通过对比项目 1 (原生 Go + database/sql)，我们可以看到框架的优势在于提高开发效率，同时保持 Go 语言的性能优势。

祝学习愉快！💪
