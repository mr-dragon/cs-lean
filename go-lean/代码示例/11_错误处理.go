package main

import (
	"fmt"
	"log"
	"os"
)

// ===== 错误处理 =====
// Java: try-catch-finally  或  throws Exception

func basicErrorHandling() {
	// ===== 使用 error 接口 =====
	// Java: try { ... } catch (Exception e) { ... }

	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Result:", result)

	// 返回多值，最后一个通常是 error
	result, err = divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)  // 输出: Error: division by zero
	}
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

// ===== 自定义错误类型 =====
// Java: extends Exception

type ValidationError struct {
	Field   string
	Message string
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("validation error on field %s: %s", e.Field, e.Message)
}

func validateEmail(email string) error {
	// Java: throw new ValidationException(...)
	if email == "" {
		return ValidationError{
			Field:   "email",
			Message: "email cannot be empty",
		}
	}
	return nil
}

// ===== 错误包装和展开 =====
// Java: throw new IllegalArgumentException(..., cause);  getCause()

import "errors"

func wrappedError() {
	// Java: throw new RuntimeException("operation failed", originalException);
	_, err := divide(10, 0)
	if err != nil {
		// 包装错误
		wrappedErr := fmt.Errorf("operation failed: %w", err)
		fmt.Println("Wrapped error:", wrappedErr)

		// 检查原始错误
		// Java: cause.getMessage()
		originalErr := errors.Unwrap(wrappedErr)
		fmt.Println("Original error:", originalErr)

		// 检查错误链中是否包含特定错误
		// Java: if (cause instanceof ...)
		if errors.Is(wrappedErr, err) {
			fmt.Println("Error chain contains the original error")
		}
	}
}

// ===== 错误的 As 方法 =====
// Java: if (exception instanceof SpecificException)

func errorAs() {
	err := validateEmail("")
	
	// 类型断言（type assertion）
	// Java: if (e instanceof ValidationError)
	var valErr ValidationError
	if errors.As(err, &valErr) {
		fmt.Printf("Validation error: %s on %s\n", valErr.Message, valErr.Field)
	}
}

// ===== Panic 和 Recover =====
// Java: unchecked exceptions, stack trace

func riskyOperation(shouldPanic bool) {
	if shouldPanic {
		// Java: throw new RuntimeException("something went wrong");
		panic("something went wrong")
	}
	fmt.Println("Operation completed successfully")
}

func panicAndRecover() {
	// Java: try-finally (用于 finally 块的清理)
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	riskyOperation(true)
	fmt.Println("This line will not be executed")
}

// ===== 使用 defer 进行清理 =====
// Java: try-finally 或 try-with-resources

func fileOperation() error {
	// Java: try (FileReader fr = new FileReader("file.txt")) { ... }
	file, err := os.Create("temp.txt")
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()  // 确保文件被关闭

	_, err = file.WriteString("Hello, Go!")
	if err != nil {
		return fmt.Errorf("failed to write: %w", err)
	}

	return nil
}

// ===== 多个 defer 的执行顺序 =====
// Java: 多个 finally 块按声明顺序执行（后进先出）

func multipleDeferExample() {
	fmt.Println("Step 1")
	defer fmt.Println("Deferred 1 (last)")
	
	fmt.Println("Step 2")
	defer fmt.Println("Deferred 2")
	
	fmt.Println("Step 3")
	defer fmt.Println("Deferred 3 (first)")
	
	fmt.Println("Step 4")
}

// ===== 错误处理最佳实践 =====

func bestPractices() error {
	// 1. 检查错误
	// Java: if (e != null)
	err := someOperation()
	if err != nil {
		// 2. 添加上下文（wrapping）
		// Java: throw new OperationException("operation failed", e);
		return fmt.Errorf("operation failed: %w", err)
	}

	// 3. 不要忽略错误
	// Java: 不要忽略异常
	err = anotherOperation()
	if err != nil {
		log.Fatal(err)  // 严重错误时终止程序
	}

	// 4. 区分可恢复和不可恢复错误
	// Java: 不同的 Exception 类型
	err = recoverableOperation()
	if err != nil {
		log.Println("Warning:", err)  // 可恢复，仅记录
	}

	return nil
}

func someOperation() error {
	return fmt.Errorf("operation error")
}

func anotherOperation() error {
	return nil
}

func recoverableOperation() error {
	return fmt.Errorf("recoverable error")
}

// ===== 定义一个简单的异常体系 =====
// Java: 继承 Exception 创建异常体系

type AppError struct {
	Code    int
	Message string
	Cause   error
}

func (e AppError) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf("Error %d: %s (caused by: %v)", e.Code, e.Message, e.Cause)
	}
	return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}

// 常见错误代码
const (
	ErrCodeValidation = 400
	ErrCodeNotFound   = 404
	ErrCodeInternal   = 500
)

func processUser(userID string) error {
	// Java: throw new NotFoundException("user not found");
	if userID == "" {
		return AppError{
			Code:    ErrCodeValidation,
			Message: "user ID cannot be empty",
		}
	}

	// 模拟数据库错误
	return AppError{
		Code:    ErrCodeInternal,
		Message: "database error",
		Cause:   fmt.Errorf("connection timeout"),
	}
}

// ===== 有时使用 log 包记录错误 =====
// Java: logger.error(), logger.warn(), logger.info()

func loggingExample() {
	// Go 标准库 log 包
	// log.Fatal() - 记录并退出
	// log.Panic() - 记录并 panic
	// log.Print() - 仅记录

	err := divide(10, 0)
	if err != nil {
		log.Println("Error occurred:", err)  // 记录错误但继续执行
		// log.Fatal(err)  // 记录并退出
	}
}

// ===== 错误处理总结表 =====

func main() {
	fmt.Println("=== Go 错误处理 ===\n")

	fmt.Println("1. 基本错误处理")
	basicErrorHandling()

	fmt.Println("\n2. 自定义错误")
	err := validateEmail("")
	fmt.Println(err)

	fmt.Println("\n3. 错误包装")
	wrappedError()

	fmt.Println("\n4. Panic 和 Recover")
	panicAndRecover()

	fmt.Println("\n5. 多个 defer")
	multipleDeferExample()

	fmt.Println("\n6. 文件操作（自动清理）")
	err = fileOperation()
	fmt.Println("File operation result:", err)
	os.Remove("temp.txt")

	fmt.Println("\n=== Go vs Java 错误处理对比 ===")
	fmt.Println(`
┌────────────────────┬──────────────────────────────┬──────────────────────┐
│     特性           │        Java                  │       Go              │
├────────────────────┼──────────────────────────────┼──────────────────────┤
│ 异常机制           │ Checked/Unchecked Exception  │ error 接口            │
│ 错误返回           │ 使用 throws 或 try-catch     │ 多返回值（最后一个） │
│ 自定义错误         │ extends Exception            │ 实现 Error 接口      │
│ 错误链             │ getCause()                   │ errors.Unwrap()      │
│ 类型检查           │ instanceof                   │ errors.As()          │
│ 不可恢复错误       │ Error, RuntimeException      │ panic() + recover()  │
│ 资源清理           │ try-finally, try-with-res   │ defer                │
│ 记录错误           │ logger.error()               │ log.Fatal()          │
│ 错误处理策略       │ 异常抛出和传播               │ 显式处理每个错误     │
└────────────────────┴──────────────────────────────┴──────────────────────┘

Go 错误处理哲学：
- 显式优于隐式（Explicit is better than implicit）
- 每个错误都要处理
- 不要隐藏错误
- 错误是值，不是异常

常见错误处理模式：

1. 检查并提前返回
   if err != nil {
       return err  // 或 return fmt.Errorf("context: %w", err)
   }

2. 包装错误（添加上下文）
   return fmt.Errorf("operation failed: %w", err)

3. 自定义错误类型
   type MyError struct { ... }
   func (e MyError) Error() string { ... }

4. 错误类型断言
   var typedErr MyError
   if errors.As(err, &typedErr) { ... }

5. 错误值检查
   if errors.Is(err, io.EOF) { ... }

6. 资源清理
   defer cleanup()  // 确保执行

7. 日志和监控
   log.Printf("Error: %v", err)

Java 相比 Go 的优势：
✓ 强制处理 checked exceptions
✓ 完整的堆栈跟踪
✓ 异常继承体系

Go 相比 Java 的优势：
✓ 更灵活，可以忽略不重要的错误
✓ 更少的异常，更多的值返回
✓ 性能更好（不需要异常处理的开销）
✓ 代码更清晰（显式错误流）
`)
}
