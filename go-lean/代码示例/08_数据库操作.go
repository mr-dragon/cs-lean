package main

import (
	"database/sql"
	"fmt"
	"time"
)

// ===== 数据库操作（使用 database/sql 包）=====
// Java: Connection conn = DriverManager.getConnection(...);

func sqlBasics() {
	// 注意：这是示例代码，需要导入相应的驱动
	// SQLite: import _ "github.com/mattn/go-sqlite3"
	// MySQL: import _ "github.com/go-sql-driver/mysql"
	// PostgreSQL: import _ "github.com/lib/pq"

	fmt.Println("数据库连接示例：")
	fmt.Println("SQLite:    sql.Open(\"sqlite3\", \"test.db\")")
	fmt.Println("MySQL:     sql.Open(\"mysql\", \"user:password@tcp(localhost:3306)/dbname\")")
	fmt.Println("PostgreSQL: sql.Open(\"postgres\", \"user=user password=password dbname=dbname sslmode=disable\")")
}

type User struct {
	ID    int
	Name  string
	Email string
	Age   int
}

func createTable(db *sql.DB) error {
	// ===== 创建表 =====
	// Java: stmt.execute("CREATE TABLE ...");

	query := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		email TEXT NOT NULL,
		age INTEGER
	);
	`

	_, err := db.Exec(query)
	return err
}

func insertData(db *sql.DB) error {
	// ===== 插入数据 =====
	// Java: stmt.executeUpdate("INSERT INTO ...");

	query := "INSERT INTO users (name, email, age) VALUES (?, ?, ?)"
	_, err := db.Exec(query, "Alice", "alice@example.com", 30)
	if err != nil {
		return err
	}

	// 获取插入的行数和最后插入的 ID
	result, _ := db.Exec(query, "Bob", "bob@example.com", 25)
	lastID, _ := result.LastInsertId()
	fmt.Println("Last inserted ID:", lastID)

	return nil
}

func queryData(db *sql.DB) error {
	// ===== 查询单行 =====
	// Java: ResultSet rs = stmt.executeQuery("SELECT ...");

	var user User
	query := "SELECT id, name, email, age FROM users WHERE id = ?"
	
	err := db.QueryRow(query, 1).Scan(&user.ID, &user.Name, &user.Email, &user.Age)
	if err != nil {
		return err
	}

	fmt.Printf("User: ID=%d, Name=%s, Email=%s, Age=%d\n", user.ID, user.Name, user.Email, user.Age)
	return nil
}

func queryMultipleRows(db *sql.DB) error {
	// ===== 查询多行 =====
	// Java: while(rs.next()) { ... }

	query := "SELECT id, name, email, age FROM users"
	rows, err := db.Query(query)
	if err != nil {
		return err
	}
	defer rows.Close()

	fmt.Println("All users:")
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Age)
		if err != nil {
			return err
		}
		fmt.Printf("ID=%d, Name=%s, Email=%s, Age=%d\n", user.ID, user.Name, user.Email, user.Age)
	}

	return rows.Err()
}

func updateData(db *sql.DB) error {
	// ===== 更新数据 =====
	// Java: stmt.executeUpdate("UPDATE ...");

	query := "UPDATE users SET age = ? WHERE name = ?"
	result, err := db.Exec(query, 31, "Alice")
	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	fmt.Println("Rows affected:", rowsAffected)
	return nil
}

func deleteData(db *sql.DB) error {
	// ===== 删除数据 =====
	// Java: stmt.executeUpdate("DELETE ...");

	query := "DELETE FROM users WHERE name = ?"
	result, err := db.Exec(query, "Bob")
	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	fmt.Println("Rows deleted:", rowsAffected)
	return nil
}

func preparedStatement(db *sql.DB) error {
	// ===== 预处理语句（防止 SQL 注入）=====
	// Java: PreparedStatement pstmt = conn.prepareStatement("SELECT ... WHERE id = ?");

	stmt, err := db.Prepare("SELECT id, name, email, age FROM users WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// 多次使用同一个预处理语句
	for i := 1; i <= 2; i++ {
		var user User
		err := stmt.QueryRow(i).Scan(&user.ID, &user.Name, &user.Email, &user.Age)
		if err != nil {
			if err == sql.ErrNoRows {
				fmt.Printf("No user found with ID %d\n", i)
			}
			continue
		}
		fmt.Printf("User: ID=%d, Name=%s\n", user.ID, user.Name)
	}

	return nil
}

func transaction(db *sql.DB) error {
	// ===== 事务 =====
	// Java: conn.setAutoCommit(false); ... conn.commit(); conn.rollback();

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	// 执行多个操作
	_, err1 := tx.Exec("INSERT INTO users (name, email, age) VALUES (?, ?, ?)", "Charlie", "charlie@example.com", 35)
	_, err2 := tx.Exec("INSERT INTO users (name, email, age) VALUES (?, ?, ?)", "Diana", "diana@example.com", 28)

	if err1 != nil || err2 != nil {
		// 回滚事务
		// Java: conn.rollback();
		tx.Rollback()
		fmt.Println("Transaction rolled back")
		return err1
	}

	// 提交事务
	// Java: conn.commit();
	err = tx.Commit()
	if err != nil {
		return err
	}

	fmt.Println("Transaction committed")
	return nil
}

func connectionPooling() {
	// ===== 连接池 =====
	// Java: HikariCP, C3P0 等连接池库

	fmt.Println("Go 的 database/sql 包自动管理连接池")
	fmt.Println("配置方法：")
	fmt.Println("  db.SetMaxOpenConns(25)      // 最大打开连接数")
	fmt.Println("  db.SetMaxIdleConns(5)       // 最大闲置连接数")
	fmt.Println("  db.SetConnMaxLifetime(time.Hour) // 连接最大生存时间")
}

// ===== ORM 框架示例结构（GORM）=====
// 实际使用需要导入: import "gorm.io/gorm"

type Product struct {
	ID    uint      `gorm:"primaryKey"`
	Name  string    `gorm:"column:name"`
	Price float64   `gorm:"column:price"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func gormExample() {
	fmt.Println("\n=== GORM 使用示例 ===")
	fmt.Println(`
import "gorm.io/gorm"
import "gorm.io/driver/sqlite"

// 连接数据库
db, _ := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

// 自动迁移（创建表）
db.AutoMigrate(&Product{})

// 创建
db.Create(&Product{Name: "Laptop", Price: 999.99})

// 查询
var product Product
db.First(&product, 1)  // 根据主键查询

// 更新
db.Model(&product).Update("price", 899.99)

// 删除
db.Delete(&product)

// 高级查询
var products []Product
db.Where("price > ?", 100).Find(&products)
`)
}

func main() {
	fmt.Println("=== 数据库操作基础 ===")
	sqlBasics()

	fmt.Println("\n=== 标准库 database/sql 操作 ===")
	fmt.Println(`
典型的数据库操作流程：

1. 导入驱动
   import _ "github.com/mattn/go-sqlite3"

2. 打开连接
   db, err := sql.Open("sqlite3", "test.db")
   defer db.Close()

3. 检查连接
   err = db.Ping()

4. 执行查询
   rows, err := db.Query("SELECT ...")
   defer rows.Close()
   for rows.Next() {
       rows.Scan(&variables...)
   }

5. 执行更新/删除
   result, err := db.Exec("INSERT/UPDATE/DELETE ...")
   rowsAffected, _ := result.RowsAffected()

6. 事务处理
   tx, err := db.Begin()
   tx.Exec(...)
   tx.Commit() 或 tx.Rollback()

7. 预处理语句
   stmt, err := db.Prepare("SELECT ... WHERE id = ?")
   stmt.QueryRow(id).Scan(...)
`)

	fmt.Println("\n=== 错误处理 ===")
	fmt.Println(`
if err == sql.ErrNoRows {
    fmt.Println("No rows found")
} else if err != nil {
    fmt.Println("Database error:", err)
}
`)

	fmt.Println("\n=== 连接池配置 ===")
	connectionPooling()

	fmt.Println("\n=== GORM 框架使用 ===")
	gormExample()

	fmt.Println("\n=== 对比 Java 和 Go ===")
	fmt.Println(`
┌─────────────────────┬────────────────────────────────────┬──────────────────────┐
│      特性           │            Java                     │        Go             │
├─────────────────────┼────────────────────────────────────┼──────────────────────┤
│ 驱动加载            │ Class.forName("...")                │ import _ "..."       │
│ 连接字符串          │ DriverManager.getConnection()       │ sql.Open()           │
│ 执行查询            │ stmt.executeQuery()                 │ db.Query()           │
│ 执行更新/删除       │ stmt.executeUpdate()                │ db.Exec()            │
│ 遍历结果集          │ while(rs.next())                    │ for rows.Next()      │
│ 提取字段            │ rs.getInt(), rs.getString()         │ rows.Scan()          │
│ 预处理语句          │ PreparedStatement                    │ stmt, _ := db.Prepare() │
│ 事务                │ conn.setAutoCommit(false)           │ db.Begin()           │
│ ORM 框架            │ Hibernate, JPA, MyBatis             │ GORM, sqlc           │
└─────────────────────┴────────────────────────────────────┴──────────────────────┘
`)
}
