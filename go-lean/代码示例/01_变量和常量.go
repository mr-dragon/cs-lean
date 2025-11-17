package main

import "fmt"

// ===== 常量声明 =====
// Java: final int MAX_USERS = 100;
const MaxUsers = 100
const (
	SmallBufSize  = 64
	LargeBufSize  = 256
	DefaultBufSize = SmallBufSize // 也可以在 const 块中引用其他常量
)

// ===== 变量声明 =====
func main() {
	// Java: int x = 10;
	var x int = 10
	fmt.Println("x:", x)

	// Java: int y = 20;  (类型推导)
	var y = 20
	fmt.Println("y:", y)

	// Java: 无对应 (短变量声明，只能在函数内使用)
	z := 30
	fmt.Println("z:", z)

	// ===== 多变量声明 =====
	// Java: int a = 1, b = 2, c = 3;
	var a, b, c int = 1, 2, 3
	fmt.Println("a, b, c:", a, b, c)

	// Java: 无对应 (短声明多个变量)
	x1, x2, x3 := 100, 200, 300
	fmt.Println("x1, x2, x3:", x1, x2, x3)

	// ===== 声明而不初始化 =====
	// Java: int uninitialized;  (Java 中使用前必须初始化，Go 中会赋予零值)
	var uninitializedInt int // 默认值: 0
	var uninitializedStr string // 默认值: ""
	var uninitializedBool bool // 默认值: false
	fmt.Println("Uninitialized int:", uninitializedInt, "string:", uninitializedStr, "bool:", uninitializedBool)

	// ===== 类型转换 =====
	// Java: int intValue = (int) 3.14;
	floatValue := 3.14
	intValue := int(floatValue)
	fmt.Println("Converted to int:", intValue)

	// Java: String str = String.valueOf(42);
	stringValue := fmt.Sprintf("%d", 42)
	fmt.Println("Converted to string:", stringValue)

	// ===== 匿名变量 =====
	// Java: 无对应
	// 使用 _ 忽略不需要的返回值
	result, _ := divideWithRemainder(10, 3)
	fmt.Println("Division result:", result)
}

// ===== 多返回值函数 =====
// Java: 需要返回一个对象或使用输出参数
func divideWithRemainder(a, b int) (int, int) {
	return a / b, a % b
}
