package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
	"time"
)

// ===== 数据模型 =====

type Article struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Title     string    `json:"title" binding:"required"`
	Content   string    `json:"content" binding:"required"`
	Author    string    `json:"author" binding:"required"`
	Category  string    `json:"category"`
	ViewCount int       `json:"view_count" gorm:"default:0"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ResponseData struct {
	Code    int         `json:"code"`      // 0 成功, 非 0 失败
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// ===== 全局变量 =====
var db *gorm.DB

// ===== 初始化数据库 =====
func initDB() error {
	var err error

	// 使用 SQLite 数据库
	db, err = gorm.Open(sqlite.Open("blog.db"), &gorm.Config{})
	if err != nil {
		return err
	}

	// 自动迁移（创建表）
	err = db.AutoMigrate(&Article{})
	if err != nil {
		return err
	}

	log.Println("Database initialized successfully")
	return nil
}

// ===== 工具函数 =====

func success(c *gin.Context, data interface{}, message string) {
	c.JSON(http.StatusOK, ResponseData{
		Code:    0,
		Message: message,
		Data:    data,
	})
}

func error(c *gin.Context, code int, message string) {
	c.JSON(http.StatusOK, ResponseData{
		Code:    code,
		Message: message,
		Data:    nil,
	})
}

// ===== API 处理器 =====

// GetArticles 获取文章列表（支持分页）
// GET /api/articles?page=1&limit=10
func GetArticles(c *gin.Context) {
	var articles []Article
	var total int64

	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "10")

	pageNum, _ := strconv.Atoi(page)
	pageSize, _ := strconv.Atoi(limit)

	if pageNum < 1 {
		pageNum = 1
	}
	if pageSize < 1 || pageSize > 50 {
		pageSize = 10
	}

	offset := (pageNum - 1) * pageSize

	// 获取总数
	db.Model(&Article{}).Count(&total)

	// 分页查询
	result := db.Offset(offset).Limit(pageSize).Order("created_at DESC").Find(&articles)

	if result.Error != nil {
		error(c, 500, "Failed to retrieve articles")
		return
	}

	success(c, gin.H{
		"articles": articles,
		"total":    total,
		"page":     pageNum,
		"limit":    pageSize,
	}, "Success")
}

// GetArticleByID 获取单篇文章
// GET /api/articles/:id
func GetArticleByID(c *gin.Context) {
	id := c.Param("id")

	var article Article
	result := db.First(&article, id)

	if result.Error == gorm.ErrRecordNotFound {
		error(c, 404, "Article not found")
		return
	}

	if result.Error != nil {
		error(c, 500, "Failed to retrieve article")
		return
	}

	// 增加浏览次数
	db.Model(&article).Update("view_count", article.ViewCount+1)

	success(c, article, "Success")
}

// CreateArticle 创建文章
// POST /api/articles
func CreateArticle(c *gin.Context) {
	var article Article

	// 绑定和验证 JSON 数据
	if err := c.ShouldBindJSON(&article); err != nil {
		error(c, 400, err.Error())
		return
	}

	// 创建记录
	result := db.Create(&article)

	if result.Error != nil {
		log.Println("Error creating article:", result.Error)
		error(c, 500, "Failed to create article")
		return
	}

	success(c, article, "Article created successfully")
}

// UpdateArticle 更新文章
// PUT /api/articles/:id
func UpdateArticle(c *gin.Context) {
	id := c.Param("id")

	var article Article
	var updateData Article

	// 检查文章是否存在
	if err := db.First(&article, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			error(c, 404, "Article not found")
		} else {
			error(c, 500, "Failed to retrieve article")
		}
		return
	}

	// 绑定更新数据
	if err := c.ShouldBindJSON(&updateData); err != nil {
		error(c, 400, err.Error())
		return
	}

	// 更新记录（只更新非零值字段）
	result := db.Model(&article).Updates(updateData)

	if result.Error != nil {
		log.Println("Error updating article:", result.Error)
		error(c, 500, "Failed to update article")
		return
	}

	success(c, article, "Article updated successfully")
}

// DeleteArticle 删除文章
// DELETE /api/articles/:id
func DeleteArticle(c *gin.Context) {
	id := c.Param("id")

	// 先检查文章是否存在
	var article Article
	if err := db.First(&article, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			error(c, 404, "Article not found")
		} else {
			error(c, 500, "Failed to retrieve article")
		}
		return
	}

	// 删除记录
	result := db.Delete(&article)

	if result.Error != nil {
		log.Println("Error deleting article:", result.Error)
		error(c, 500, "Failed to delete article")
		return
	}

	success(c, gin.H{"id": id}, "Article deleted successfully")
}

// SearchArticles 搜索文章
// GET /api/articles/search?q=keyword
func SearchArticles(c *gin.Context) {
	keyword := c.Query("q")

	if keyword == "" {
		error(c, 400, "Search keyword is required")
		return
	}

	var articles []Article
	result := db.Where("title LIKE ? OR content LIKE ? OR author LIKE ?", 
		"%"+keyword+"%", 
		"%"+keyword+"%", 
		"%"+keyword+"%").
		Order("created_at DESC").
		Find(&articles)

	if result.Error != nil {
		error(c, 500, "Failed to search articles")
		return
	}

	success(c, articles, "Search completed")
}

// GetArticlesByCategory 按分类获取文章
// GET /api/articles/category/:name
func GetArticlesByCategory(c *gin.Context) {
	category := c.Param("name")

	var articles []Article
	result := db.Where("category = ?", category).Order("created_at DESC").Find(&articles)

	if result.Error != nil {
		error(c, 500, "Failed to retrieve articles")
		return
	}

	success(c, articles, "Success")
}

// GetStats 获取统计信息
// GET /api/stats
func GetStats(c *gin.Context) {
	var total int64
	var totalViews int64

	db.Model(&Article{}).Count(&total)
	db.Model(&Article{}).Select("COALESCE(SUM(view_count), 0)").Row().Scan(&totalViews)

	success(c, gin.H{
		"total_articles": total,
		"total_views":    totalViews,
	}, "Success")
}

// ===== 路由配置 =====

func main() {
	// 初始化数据库
	if err := initDB(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// 创建 Gin 路由器
	router := gin.Default()

	// 添加 CORS 中间件
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	// 添加日志中间件
	// 默认已启用，如需禁用可使用 gin.New()

	// API 路由
	api := router.Group("/api")
	{
		// 文章相关路由
		articles := api.Group("/articles")
		{
			articles.GET("", GetArticles)           // 获取文章列表
			articles.GET("/:id", GetArticleByID)    // 获取单篇文章
			articles.POST("", CreateArticle)        // 创建文章
			articles.PUT("/:id", UpdateArticle)     // 更新文章
			articles.DELETE("/:id", DeleteArticle)  // 删除文章
		}

		// 搜索和分类路由
		api.GET("/search", SearchArticles)                    // 搜索文章
		api.GET("/category/:name", GetArticlesByCategory)     // 按分类获取

		// 统计信息
		api.GET("/stats", GetStats)
	}

	// 健康检查端点
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	log.Println("Server starting on http://localhost:8080")
	log.Println("API Documentation:")
	log.Println("  GET    /api/articles              - Get articles (with pagination)")
	log.Println("  GET    /api/articles/:id          - Get article by ID")
	log.Println("  POST   /api/articles              - Create article")
	log.Println("  PUT    /api/articles/:id          - Update article")
	log.Println("  DELETE /api/articles/:id          - Delete article")
	log.Println("  GET    /api/search?q=keyword     - Search articles")
	log.Println("  GET    /api/category/:name       - Get articles by category")
	log.Println("  GET    /api/stats                - Get statistics")

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
