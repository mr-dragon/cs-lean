# Blog API 前端文档

Vue.js 3 + Axios 实现的响应式博客管理系统前端。

## 🎯 功能概览

| 功能 | 说明 | 实现方式 |
|------|------|--------|
| 文章列表 | 分页显示文章 | 虚拟滚动 + 分页组件 |
| 搜索功能 | 按关键词模糊搜索 | 防抖 + Axios GET |
| 创建文章 | 填表发布新文章 | 表单验证 + POST 请求 |
| 编辑文章 | 修改现有文章内容 | 表单预填充 + PUT 请求 |
| 删除文章 | 删除不需要的文章 | 二次确认 + DELETE 请求 |
| 查看详情 | 完整浏览单篇文章 | 获取详情 + 自动统计浏览 |
| 统计信息 | 显示数据统计面板 | 异步加载 + 实时更新 |
| 响应式设计 | 适配各种设备 | CSS Grid + Flexbox |

## 📂 文件说明

```
frontend/
├── index.html          # 主应用文件（单页应用）
└── README.md          # 本文档
```

**为什么是单文件应用?**
- 简化部署，只需打开 HTML 文件
- 无需构建工具链
- 便于学习 Vue.js 基础
- 生产环境建议用 Vite/Vue-CLI 拆分组件

## 🚀 快速启动

### 方式 1: 直接打开文件

```bash
# Linux/Mac
open frontend/index.html

# Windows
start frontend/index.html

# 或在浏览器中打开
file:///workspaces/cs-lean/go-lean/项目实战/blog-api/frontend/index.html
```

### 方式 2: 本地 Web 服务器

```bash
cd frontend

# 使用 Python 3
python3 -m http.server 8000

# 或 Python 2
python -m SimpleHTTPServer 8000

# 或 Node.js (需要先安装 http-server)
npx http-server

# 然后访问
http://localhost:8000
```

### 方式 3: 使用 VS Code Live Server 扩展

1. 右键点击 `index.html` → "Open with Live Server"
2. 浏览器自动打开 `http://127.0.0.1:5500`

**重要**: 后端服务必须运行在 `http://localhost:8080`！

## 🏗️ 代码结构

### 1. HTML 结构

```html
<body>
    <div id="app">
        <!-- Vue 应用根元素 -->
    </div>
</body>
```

### 2. 样式层次

```css
/* 基础 */
* { margin: 0; padding: 0; box-sizing: border-box; }
body { font-family: 'Segoe UI'; background: #f5f7fa; }

/* 组件 */
.navbar { /* 导航栏 */ }
.article-card { /* 文章卡片 */ }
.modal { /* 弹窗 */ }
.form-group { /* 表单项 */ }
.pagination { /* 分页 */ }

/* 响应式 */
@media (max-width: 768px) { /* 移动设备适配 */ }
```

### 3. Vue 组件树

```
#app (根应用)
├── navbar (导航栏)
│   └── menu (菜单按钮)
├── view-list (文章列表视图)
│   ├── stats (统计卡片)
│   ├── search-bar (搜索条)
│   ├── article-card * N (文章卡片)
│   └── pagination (分页)
├── view-create (创建/编辑视图)
│   └── form (文章表单)
└── view-detail (详情视图)
    └── article-content (文章内容)
```

## 💾 数据模型

### 文章对象 (Article)

```javascript
{
    id: 1,                          // 唯一标识
    title: "Go 并发编程",            // 标题
    content: "Goroutines 是...",     // 内容
    author: "张三",                  // 作者
    category: "Go",                  // 分类
    view_count: 42,                  // 浏览次数
    created_at: "2024-01-15T10:30:45Z", // 创建时间
    updated_at: "2024-01-15T10:30:45Z"  // 更新时间
}
```

### 表单对象 (ArticleForm)

```javascript
{
    title: "",      // 必填
    content: "",    // 必填
    author: "",     // 必填
    category: ""    // 可选
}
```

### 消息对象 (Message)

```javascript
{
    text: "操作成功",
    type: "success"  // "success" | "error"
}
```

### 统计对象 (Stats)

```javascript
{
    total_articles: 15,  // 文章总数
    total_views: 428,    // 总浏览次数
    categories: 4,       // 分类数量
    authors: 5           // 作者数量
}
```

## 🎭 视图管理

### 视图状态

```javascript
// 主要通过 view 变量控制显示
view: 'list'      // 显示文章列表
view: 'create'    // 显示创建/编辑表单
view: 'detail'    // 显示文章详情
```

### 视图切换流程

```
列表 → 搜索结果/分页
  ↓
  → 查看详情
  ↓
  → 编辑 → 表单 → 发布/取消 → 回到列表
  ↓
  → 删除 (二次确认) → 回到列表
  ↓
  新建 → 表单 → 发布/取消 → 回到列表
```

## 🔧 关键功能详解

### 1. 文章列表加载

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
        console.error('Error loading articles:', error);
        showMessage('加载文章失败', 'error');
    } finally {
        loading.value = false;
    }
};
```

**执行步骤**:
1. 设置 loading = true（显示加载动画）
2. 发送 GET 请求到 `/api/articles`
3. 检查响应码，0 表示成功
4. 更新 articles 数组（自动触发 UI 重新渲染）
5. 更新总数用于分页计算
6. 捕获错误并显示提示
7. 最后设置 loading = false（隐藏加载动画）

### 2. 搜索功能

```javascript
const handleSearch = async () => {
    if (!searchKeyword.value.trim()) {
        currentPage.value = 1;
        loadArticles();  // 关键词为空则加载全部
        return;
    }

    loading.value = true;
    try {
        const response = await axios.get(`${API_BASE}/search`, {
            params: { q: searchKeyword.value }
        });

        if (response.data.code === 0) {
            articles.value = response.data.data || [];
            totalCount.value = articles.value.length;  // 搜索结果长度
            currentPage.value = 1;
        }
    } catch (error) {
        showMessage('搜索失败', 'error');
    } finally {
        loading.value = false;
    }
};
```

**与普通搜索的区别**:
- 不分页（一次加载所有结果）
- 总数 = 搜索结果长度
- 关键词为空时重置为列表视图

### 3. 文章创建

```javascript
const handleSubmit = async () => {
    // 1. 前端验证
    if (!articleForm.value.title || !articleForm.value.content || !articleForm.value.author) {
        showMessage('请填写所有必填字段', 'error');
        return;
    }

    loading.value = true;
    try {
        if (editingArticle.value) {
            // 编辑模式：PUT 请求
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
            // 创建模式：POST 请求
            const response = await axios.post(`${API_BASE}/articles`, articleForm.value);
            if (response.data.code === 0) {
                showMessage('文章发布成功');
                resetForm();
                currentPage.value = 1;
                loadArticles();  // 回到第一页
            }
        }
    } catch (error) {
        // 显示后端返回的具体错误信息
        showMessage(error.response?.data?.message || '操作失败', 'error');
    } finally {
        loading.value = false;
    }
};
```

**三层验证**:
1. 前端表单验证（快速反馈）
2. 后端字段验证（数据安全）
3. 后端业务验证（长度限制等）

### 4. 文章删除

```javascript
const deleteArticle = async (id) => {
    // 1. 二次确认
    if (!confirm('确定要删除这篇文章吗？')) return;

    loading.value = true;
    try {
        const response = await axios.delete(`${API_BASE}/articles/${id}`);
        if (response.data.code === 0) {
            showMessage('文章删除成功');
            view.value = 'list';      // 返回列表
            loadArticles();             // 刷新数据
        }
    } catch (error) {
        showMessage('删除失败', 'error');
    } finally {
        loading.value = false;
    }
};
```

**安全设计**:
- 使用 `confirm()` 进行二次确认
- 删除后自动刷新列表
- 显示操作反馈

### 5. 文章详情查看

```javascript
const viewArticle = async (id) => {
    loading.value = true;
    try {
        const response = await axios.get(`${API_BASE}/articles/${id}`);
        if (response.data.code === 0) {
            selectedArticle.value = response.data.data;
            view.value = 'detail';
        }
    } catch (error) {
        showMessage('加载文章详情失败', 'error');
    } finally {
        loading.value = false;
    }
};
```

**特点**:
- 获取详情时自动增加浏览数（后端处理）
- 返回完整文章内容（列表视图中是截断的）
- 显示详情视图

### 6. 分页功能

```javascript
// 计算总页数
const totalPages = computed(() => {
    return Math.ceil(totalCount.value / pageSize.value);
});

// 上一页
const previousPage = () => {
    if (currentPage.value > 1) {
        currentPage.value--;
        loadArticles();
    }
};

// 下一页
const nextPage = () => {
    if (currentPage.value < totalPages.value) {
        currentPage.value++;
        loadArticles();
    }
};

// 跳页
const goToPage = (page) => {
    currentPage.value = page;
    loadArticles();
};
```

**分页逻辑**:
- totalPages = Math.ceil(总数 / 每页数)
- 页码范围：1 到 totalPages
- 禁用按钮防止无效操作

### 7. 消息提示

```javascript
const showMessage = (text, type = 'success') => {
    message.value = { text, type };
    
    // 3 秒后自动隐藏
    setTimeout(() => {
        message.value = null;
    }, 3000);
};
```

**样式**:
- success: 绿色 (#48bb78)
- error: 红色 (#f56565)
- 右上角浮动显示
- 自动滑出动画

## 🎨 样式系统

### 色彩方案

```css
/* 主色系 */
--primary: #667eea      /* 深蓝紫 - 按钮、链接 */
--primary-dark: #5568d3 /* 深蓝紫 - hover */
--success: #48bb78      /* 绿色 - 成功消息 */
--danger: #f56565       /* 红色 - 删除、错误 */
--bg: #f5f7fa           /* 浅灰 - 页面背景 */
--card-bg: #ffffff      /* 白色 - 卡片背景 */
--text: #333333         /* 深灰 - 主文本 */
--text-muted: #999999   /* 浅灰 - 辅助文本 */
--border: #e0e0e0       /* 浅灰 - 边框 */
```

### 组件样式

#### 导航栏
```css
.navbar {
    background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
    position: sticky;
    top: 0;
    z-index: 100;
}
```

#### 卡片
```css
.article-card {
    background: white;
    border-radius: 10px;
    box-shadow: 0 2px 10px rgba(0, 0, 0, 0.05);
    transition: all 0.3s;
}

.article-card:hover {
    box-shadow: 0 5px 20px rgba(0, 0, 0, 0.1);
    transform: translateY(-5px);  /* 向上浮起 */
}
```

#### 按钮
```css
.btn-primary {
    background: #667eea;
    color: white;
    transition: all 0.3s;
}

.btn-primary:hover {
    background: #5568d3;  /* 点击时变深 */
}
```

### 响应式断点

```css
@media (max-width: 768px) {
    /* 平板和手机 */
    .navbar-menu { flex-direction: column; }
    .stats { grid-template-columns: 1fr; }  /* 单列 */
    .search-bar { flex-direction: column; }
    .modal { width: 95%; }  /* 充满屏幕 */
}
```

## 🔌 API 集成

### 请求配置

```javascript
const API_BASE = 'http://localhost:8080/api';

// 所有请求都通过 axios
// 如需认证，可添加请求拦截器
// axios.interceptors.request.use(config => {
//     config.headers.Authorization = `Bearer ${token}`;
//     return config;
// });
```

### 请求方法

| 方法 | 端点 | 用途 |
|------|------|------|
| GET | /articles | 获取列表 |
| GET | /articles/:id | 获取详情 |
| POST | /articles | 创建 |
| PUT | /articles/:id | 更新 |
| DELETE | /articles/:id | 删除 |
| GET | /search?q= | 搜索 |
| GET | /category/:name | 按分类过滤 |
| GET | /stats | 获取统计 |

### 错误处理

```javascript
catch (error) {
    // 后端 API 错误
    if (error.response) {
        console.error('Status:', error.response.status);
        console.error('Message:', error.response.data.message);
        showMessage(error.response.data.message || '操作失败', 'error');
    }
    // 网络错误
    else if (error.request) {
        console.error('No response:', error.request);
        showMessage('网络连接失败', 'error');
    }
    // 其他错误
    else {
        console.error('Error:', error.message);
    }
}
```

## 🧪 测试清单

### 功能测试

- [ ] 加载文章列表
- [ ] 分页切换（上一页、下一页、直接跳页）
- [ ] 搜索文章（关键词、重置）
- [ ] 创建文章（填表、提交、成功提示）
- [ ] 编辑文章（打开、修改、保存）
- [ ] 删除文章（确认、删除、刷新）
- [ ] 查看详情（浏览次数+1）
- [ ] 统计信息（实时显示）
- [ ] 响应式设计（桌面、平板、手机）

### 错误处理测试

- [ ] 后端离线时提示
- [ ] 网络超时时提示
- [ ] 服务器错误返回时提示
- [ ] 文章不存在时提示
- [ ] 表单验证失败时提示

### UI/UX 测试

- [ ] 加载动画显示
- [ ] 空状态提示
- [ ] 错误消息清晰
- [ ] 按钮禁用正确
- [ ] 动画流畅

## ⚙️ 部署方案

### 方案 1: 静态服务器

```bash
# Nginx 配置
server {
    listen 80;
    server_name example.com;
    
    location / {
        root /path/to/frontend;
        try_files $uri $uri/ /index.html;
    }
}
```

### 方案 2: Docker

```dockerfile
FROM nginx:alpine
COPY index.html /usr/share/nginx/html/
COPY nginx.conf /etc/nginx/nginx.conf
```

### 方案 3: 云存储 (静态网站托管)

- AWS S3 + CloudFront
- Azure Static Web Apps
- Google Cloud Storage
- 阿里云 OSS

## 📚 Vue.js 3 学习要点

### 组合式 API (Composition API)

```javascript
import { ref, computed, onMounted } from 'vue';

export default {
    setup() {
        // 响应式数据
        const count = ref(0);
        const doubled = computed(() => count.value * 2);
        
        // 生命周期
        onMounted(() => {
            console.log('组件挂载');
        });
        
        return { count, doubled };
    }
}
```

### 指令

| 指令 | 用途 | 示例 |
|------|------|------|
| v-if | 条件渲染 | `<div v-if="show">` |
| v-for | 列表渲染 | `<li v-for="item in items">` |
| v-bind | 属性绑定 | `<img :src="url">` |
| @click | 事件绑定 | `<button @click="handleClick">` |
| v-model | 双向绑定 | `<input v-model="text">` |

### 响应式系统

```javascript
// ref：包装基础类型
const name = ref('张三');
console.log(name.value);  // 需要 .value 访问

// computed：计算属性
const fullName = computed(() => {
    return name.value + '李四';
});

// 自动追踪依赖，数据变化自动更新 UI
```

## 🚨 常见问题

### Q: 前端无法连接后端

**原因**: API_BASE 配置错误或后端未启动

**检查**:
```javascript
// 1. 确保后端运行在 localhost:8080
// 2. 后端启用了 CORS
// 3. 不要用 file:// 协议打开 HTML
```

### Q: 表单提交后没有反应

**原因**: 可能是前端验证失败或请求被阻止

**调试**:
```javascript
// 打开浏览器开发者工具
// F12 → Network 标签，查看请求状态
// Console 标签查看错误信息
```

### Q: 页面加载缓慢

**优化**:
```javascript
// 1. 减少初始加载的数据量（分页）
// 2. 使用虚拟滚动处理大列表
// 3. 压缩 JavaScript 和 CSS
// 4. 使用 CDN 加速
```

## 🔄 迭代建议

### Phase 1: 核心功能 ✅
- 列表、创建、编辑、删除
- 分页、搜索
- 统计面板

### Phase 2: 用户体验
- 本地存储草稿
- 快捷键支持（Ctrl+S 保存）
- 文章预览
- 图片上传

### Phase 3: 高级功能
- 用户认证
- 评论系统
- 标签系统
- 协作编辑

### Phase 4: 优化
- 离线支持 (PWA)
- Service Worker 缓存
- 虚拟滚动
- CDN 加速

## 📞 技术支持

遇到问题？

1. 检查浏览器控制台（F12）
2. 查看网络请求（Network 标签）
3. 参考官方文档：
   - [Vue.js 3 官方文档](https://vuejs.org/)
   - [Axios 官方文档](https://axios-http.com/)

## 📄 许可证

MIT License - 自由使用和修改
