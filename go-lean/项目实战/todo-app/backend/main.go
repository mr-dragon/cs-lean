package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

// ===== 数据模型 =====
type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Desc  string `json:"desc"`
	Done  bool   `json:"done"`
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// ===== 全局数据库连接 =====
var db *sql.DB

// ===== 初始化数据库 =====
func initDB() error {
	var err error
	db, err = sql.Open("sqlite3", "todos.db")
	if err != nil {
		return err
	}

	// 测试连接
	err = db.Ping()
	if err != nil {
		return err
	}

	// 创建表
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

// ===== 中间件：CORS 跨域处理 =====
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// ===== 中间件：日志记录 =====
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[%s] %s %s", r.Method, r.URL.Path, r.RemoteAddr)
		next.ServeHTTP(w, r)
	})
}

// ===== 工具函数：返回 JSON 响应 =====
func sendJSON(w http.ResponseWriter, code int, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	resp := Response{
		Code:    code,
		Message: message,
		Data:    data,
	}

	json.NewEncoder(w).Encode(resp)
}

// ===== 工具函数：返回错误响应 =====
func sendError(w http.ResponseWriter, code int, message string) {
	sendJSON(w, code, message, nil)
}

// ===== API 处理器 =====

// GET /api/todos - 获取所有待办项
func getTodos(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, title, desc, done FROM todos ORDER BY id DESC")
	if err != nil {
		log.Println("Error querying todos:", err)
		sendError(w, 500, "Failed to retrieve todos")
		return
	}
	defer rows.Close()

	todos := []Todo{}
	for rows.Next() {
		var todo Todo
		err := rows.Scan(&todo.ID, &todo.Title, &todo.Desc, &todo.Done)
		if err != nil {
			log.Println("Error scanning todo:", err)
			continue
		}
		todos = append(todos, todo)
	}

	if err = rows.Err(); err != nil {
		log.Println("Error iterating todos:", err)
		sendError(w, 500, "Error reading todos")
		return
	}

	// 如果没有结果，返回空数组而不是 null
	if todos == nil {
		todos = []Todo{}
	}

	sendJSON(w, 0, "Success", todos)
}

// GET /api/todos/{id} - 获取单个待办项
func getTodoByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		sendError(w, 400, "Invalid ID format")
		return
	}

	var todo Todo
	err = db.QueryRow("SELECT id, title, desc, done FROM todos WHERE id = ?", id).
		Scan(&todo.ID, &todo.Title, &todo.Desc, &todo.Done)

	if err == sql.ErrNoRows {
		sendError(w, 404, "Todo not found")
		return
	} else if err != nil {
		log.Println("Error querying todo:", err)
		sendError(w, 500, "Failed to retrieve todo")
		return
	}

	sendJSON(w, 0, "Success", todo)
}

// POST /api/todos - 创建待办项
func createTodo(w http.ResponseWriter, r *http.Request) {
	var todo Todo

	// 解析请求体
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		sendError(w, 400, "Invalid request body")
		return
	}

	// 验证输入
	if todo.Title == "" {
		sendError(w, 400, "Title is required")
		return
	}

	// 插入数据库
	result, err := db.Exec(
		"INSERT INTO todos (title, desc, done) VALUES (?, ?, ?)",
		todo.Title,
		todo.Desc,
		todo.Done,
	)

	if err != nil {
		log.Println("Error inserting todo:", err)
		sendError(w, 500, "Failed to create todo")
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		sendError(w, 500, "Failed to get inserted ID")
		return
	}

	todo.ID = int(id)
	sendJSON(w, 0, "Todo created successfully", todo)
}

// PUT /api/todos/{id} - 更新待办项
func updateTodo(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		sendError(w, 400, "Invalid ID format")
		return
	}

	var todo Todo
	err = json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		sendError(w, 400, "Invalid request body")
		return
	}

	// 验证输入
	if todo.Title == "" {
		sendError(w, 400, "Title is required")
		return
	}

	// 更新数据库
	result, err := db.Exec(
		"UPDATE todos SET title = ?, desc = ?, done = ? WHERE id = ?",
		todo.Title,
		todo.Desc,
		todo.Done,
		id,
	)

	if err != nil {
		log.Println("Error updating todo:", err)
		sendError(w, 500, "Failed to update todo")
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		sendError(w, 500, "Failed to get rows affected")
		return
	}

	if rowsAffected == 0 {
		sendError(w, 404, "Todo not found")
		return
	}

	todo.ID = id
	sendJSON(w, 0, "Todo updated successfully", todo)
}

// DELETE /api/todos/{id} - 删除待办项
func deleteTodo(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		sendError(w, 400, "Invalid ID format")
		return
	}

	result, err := db.Exec("DELETE FROM todos WHERE id = ?", id)
	if err != nil {
		log.Println("Error deleting todo:", err)
		sendError(w, 500, "Failed to delete todo")
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		sendError(w, 500, "Failed to get rows affected")
		return
	}

	if rowsAffected == 0 {
		sendError(w, 404, "Todo not found")
		return
	}

	sendJSON(w, 0, "Todo deleted successfully", map[string]interface{}{"id": id})
}

// DELETE /api/todos - 清空所有完成的任务
func deleteDoneTodos(w http.ResponseWriter, r *http.Request) {
	result, err := db.Exec("DELETE FROM todos WHERE done = 1")
	if err != nil {
		log.Println("Error deleting done todos:", err)
		sendError(w, 500, "Failed to delete done todos")
		return
	}

	rowsAffected, _ := result.RowsAffected()
	sendJSON(w, 0, "Done todos deleted", map[string]interface{}{"deleted": rowsAffected})
}

// POST /api/todos/{id}/toggle - 切换完成状态
func toggleTodo(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		sendError(w, 400, "Invalid ID format")
		return
	}

	// 查询当前状态
	var done bool
	err = db.QueryRow("SELECT done FROM todos WHERE id = ?", id).Scan(&done)
	if err == sql.ErrNoRows {
		sendError(w, 404, "Todo not found")
		return
	} else if err != nil {
		sendError(w, 500, "Failed to query todo")
		return
	}

	// 更新状态
	_, err = db.Exec("UPDATE todos SET done = ? WHERE id = ?", !done, id)
	if err != nil {
		sendError(w, 500, "Failed to toggle todo")
		return
	}

	sendJSON(w, 0, "Todo toggled", map[string]interface{}{"id": id, "done": !done})
}

// ===== 路由配置 =====
func main() {
	// 初始化数据库
	err := initDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	log.Println("Database initialized successfully")

	// 创建 HTTP 服务器多路复用器
	mux := http.NewServeMux()

	// 注册 API 路由
	mux.HandleFunc("/api/todos", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			getTodos(w, r)
		} else if r.Method == http.MethodPost {
			createTodo(w, r)
		} else if r.Method == http.MethodDelete {
			deleteDoneTodos(w, r)
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/api/todos/detail", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			getTodoByID(w, r)
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/api/todos/update", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPut {
			updateTodo(w, r)
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/api/todos/delete", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodDelete {
			deleteTodo(w, r)
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/api/todos/toggle", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			toggleTodo(w, r)
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	// 应用中间件
	handler := corsMiddleware(loggingMiddleware(mux))

	// 启动服务器
	port := ":8080"
	log.Printf("Server starting on http://localhost%s", port)
	log.Printf("API Documentation:\n")
	log.Printf("  GET    /api/todos              - Get all todos\n")
	log.Printf("  GET    /api/todos/detail?id=N - Get todo by ID\n")
	log.Printf("  POST   /api/todos              - Create todo\n")
	log.Printf("  PUT    /api/todos/update?id=N - Update todo\n")
	log.Printf("  DELETE /api/todos/delete?id=N - Delete todo\n")
	log.Printf("  POST   /api/todos/toggle?id=N - Toggle todo status\n")
	log.Printf("  DELETE /api/todos              - Delete all done todos\n")

	err = http.ListenAndServe(port, handler)
	if err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
